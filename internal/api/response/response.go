package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// NewAPIJSONResponse creates a new APIResponse with given data.
// It also sets the header "Content-Type" to "application/json".
func NewAPIJSONResponse(c *gin.Context, data interface{}, err error) APIResponse {
	c.Header("Content-Type", "application/json")
	res := APIResponse{
		Data: data,
	}
	if err != nil {
		res.Error = err.Error()
	}
	return res
}

// NewAPIResponse creates a new APIResponse with given data.
func NewAPIResponse(c *gin.Context, data interface{}, err error) APIResponse {
	c.Header("Content-Type", "application/json")
	res := APIResponse{
		Data: data,
	}
	if err != nil {
		res.Error = err.Error()
	}
	return res
}
