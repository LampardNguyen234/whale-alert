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
	accountGroup.GET("/all-monitored", s.APIAuthenticateHandler(), s.apiGetAllMonitoredAccounts)

	tokenGroup := adminGroup.Group("/token")
	tokenGroup.POST("/update", s.APIAuthenticateHandler(), s.apiUpdateTokenDetail)
	tokenGroup.GET("/all", s.APIAuthenticateHandler(), s.apiGetAllTokenDetail)

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
		s.log.Errorf("failed to bind request: %v", err)
		statusCode = http.StatusBadRequest
		return
	}

	addr, err := common.AccountAddressToHex(req.Address)
	if err != nil {
		s.log.Errorf("invalid address %v: err", req.Address, err)
		statusCode = http.StatusBadRequest
		err = fmt.Errorf("invalid address %v", req.Address)
	}

	go func() {
		d := store.AccountDetail{
			Address:   addr,
			Name:      req.Name,
			Monitored: req.Monitored,
		}
		err = s.db.StoreAccountDetail(d)
		if err != nil {
			s.log.Errorf("failed to store account detail %v: %v", d, err)
		}
	}()
}

// apiGetAllAccountDetail godoc
// @Summary Get all account details
// @Description Get all stored account details.
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

// apiGetAllMonitoredAccounts godoc
// @Summary Get all monitored accounts
// @Description Retrieve all the monitored accounts
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse.
// @Security ApiKeyAuth
// @Router  /admin/account/all-monitored [get]
func (s *HTTPServer) apiGetAllMonitoredAccounts(c *gin.Context) {
	var err error
	statusCode := http.StatusOK
	resp := response.APIAllMonitoredAccountResponse{}
	defer func() {
		c.JSON(statusCode, response.NewAPIJSONResponse(c, resp, err))
	}()

	allAccounts, err := s.db.GetAllMonitoredAccounts()
	if err != nil {
		s.log.Errorf("failed to GetAllAccountDetails: %v", err)
		statusCode = http.StatusBadRequest
		err = fmt.Errorf("internal server error")
		return
	}

	resp.Result = allAccounts
}

// apiUpdateTokenDetail godoc
// @Summary Update Token Detail
// @Description Store the detail of a token
// @Accept json
// @Param request body request.APIUpdateTokenDetail true "Request body"
// @Produce json
// @Success 200 {object} response.APIResponse.
// @Security ApiKeyAuth
// @Router  /admin/token/update [post]
func (s *HTTPServer) apiUpdateTokenDetail(c *gin.Context) {
	var err error
	statusCode := http.StatusBadRequest
	defer func() {
		var resp string
		if err == nil {
			resp = "ok"
		}
		c.JSON(statusCode, response.NewAPIJSONResponse(c, resp, err))
	}()

	var req request.APIUpdateTokenDetail
	err = c.MustBindWith(&req, binding.JSON)
	if err != nil {
		s.log.Errorf("failed to bind request: %v", err)
		return
	}
	if _, err = req.IsValid(); err != nil {
		return
	}

	addr, err := common.AccountAddressToEthAddr(req.TokenAddress)
	if err != nil {
		s.log.Errorf("invalid address %v: err", req.TokenAddress, err)
		err = fmt.Errorf("invalid address %v", req.TokenAddress)
		return
	}
	req.TokenAddress = addr.Hex()
	statusCode = http.StatusOK

	go func() {
		err = s.db.UpdateTokenDetail(store.TokenDetail(req))
		if err != nil {
			s.log.Errorf("failed to store token detail %v: %v", req, err)
		}
	}()
}

// apiGetAllAccountDetail godoc
// @Summary Get all token details
// @Description Get all stored token details.
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse.
// @Security ApiKeyAuth
// @Router  /admin/token/all [get]
func (s *HTTPServer) apiGetAllTokenDetail(c *gin.Context) {
	var err error
	statusCode := http.StatusBadRequest
	resp := response.APIAllTokenDetailResponse{}
	defer func() {
		c.JSON(statusCode, response.NewAPIJSONResponse(c, resp, err))
	}()

	resp.Result, err = s.db.GetAllTokenDetails()
	if err != nil {
		s.log.Errorf("failed to get all token details: %v", err)
		return
	}

	statusCode = http.StatusOK
}
