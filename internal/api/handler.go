package api

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/internal/api/request"
	"github.com/LampardNguyen234/whale-alert/internal/api/response"
	"github.com/LampardNguyen234/whale-alert/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (s *HTTPServer) startHandler() {
	// Ping handler
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	// Swagger docs handler
	s.Engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminGroup := s.Engine.Group("/admin")

	accountGroup := adminGroup.Group("/account")
	accountGroup.POST("/update", s.APIAuthenticateHandler(), s.apiUpdateAccountDetail)
	accountGroup.GET("/all", s.APIAuthenticateHandler(), s.apiGetAllAccountDetail)

	err := s.Engine.Run("0.0.0.0:12321")
	if err != nil {
		s.log.Panicf("%v", err)
	}
	return
}

// apiUpdateAccountDetail godoc
// @Summary Update Account Detail
// @Description Store the detail of an account
// @Accept json
// @Param request body request.APIUpdateAccountDetail true "Request body"
// @Produce json
// @Success 200 {object} response.APIResponse.
// @Security ApiKeyAuth
// @Router  /admin/account/update [post]
func (s *HTTPServer) apiUpdateAccountDetail(c *gin.Context) {
	var err error
	statusCode := http.StatusOK
	defer func() {
		var resp string
		if err == nil {
			resp = "ok"
		}
		c.JSON(statusCode, response.NewAPIJSONResponse(c, resp, err))
	}()

	var req request.APIUpdateAccountDetail
	err = c.MustBindWith(&req, binding.JSON)
	if err != nil {
		s.log.Errorf("Bind request error: %v", err)
		statusCode = http.StatusBadRequest
		return
	}

	addr, err := common.AccountAddressToHex(req.Address)
	if err != nil {
		s.log.Errorf("Invalid address %v: err", req.Address, err)
		statusCode = http.StatusBadRequest
		err = fmt.Errorf("invalid address %v", req.Address)
	}

	go func() {
		d := store.AccountDetail{
			Address: addr,
			Name:    req.Name,
		}
		err = s.db.StoreAccountDetail(d)
		if err != nil {
			s.log.Errorf("failed to store account detail %v: %v", d, err)
		}
	}()
}

// apiGetAllAccountDetail godoc
// @Summary Message Status
// @Description Check the status of a message given its ID.
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse.
// @Security ApiKeyAuth
// @Router  /admin/account/all [get]
func (s *HTTPServer) apiGetAllAccountDetail(c *gin.Context) {
	var err error
	statusCode := http.StatusOK
	resp := response.APIAllAccountDetailResponse{}
	defer func() {
		c.JSON(statusCode, response.NewAPIJSONResponse(c, resp, err))
	}()

	allAccounts, err := s.db.GetAllAccountDetails()
	if err != nil {
		s.log.Errorf("failed to GetAllAccountDetails: %v", err)
		statusCode = http.StatusBadRequest
		err = fmt.Errorf("internal server error")
		return
	}

	resp.Result = allAccounts
}
