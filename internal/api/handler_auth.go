package api

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authKey = "097a5daf05c90ed9105d983f74346c01e1ab5e2199064266b31b742b71303db4"

// APIAuthenticateHandler is a middleware for authenticating the request based on ApiKey.
func (s *HTTPServer) APIAuthenticateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.ClientIP(), "127.0.0.1") && !strings.Contains(c.ClientIP(), "::1") {
			var token struct {
				Authorization string `json:"Authorization" binding:"required"`
			}

			err := c.ShouldBindHeader(&token)
			if err != nil {
				s.log.Errorf("Invalid header: %v", err)
				unauthorized(c, http.StatusUnauthorized, err.Error())
				return
			}
			s.log.Infof("New request from %v, path %v", c.ClientIP(), c.Request.URL.Path)

			if strings.ToLower(
				fmt.Sprintf("%x", sha256.Sum256([]byte(token.Authorization))),
			) != authKey { // only basic authentication =)) HackMeIfYouCan
				s.log.Errorf("Invalid token: %v", token.Authorization)
				unauthorized(c, http.StatusUnauthorized, err.Error())
				return
			}
		}

		c.Next()
	}
}

func unauthorized(c *gin.Context, code int, message string) {
	c.Abort()
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
