package middleware

import (
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/helpers/tokenhelper"
	"codemax_assignment/routes"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitMiddleWare - to initialise open and restricted routes
func InitMiddleWare(g *gin.Engine) {
	g.Use(cors.Default())
	o := g.Group("/o")
	o.Use(OpenRequestMiddleware())
	r := g.Group("/r")
	r.Use(RestrictedRequestMiddleware())
	routes.InitUserRoutes(o, r)
	routes.InitEmailRoutes(o, r)
}

// OpenRequestMiddleware - if open route with context "/o" is requested
func OpenRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggerhelper.LogInfo("OpenRequestMiddleware Requested!!")
	}
}

// RestrictedRequestMiddleware - if restricted route with context "/r" is requested
func RestrictedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggerhelper.LogInfo("RestrictedRequestMiddleware Requested!!")
		token := c.GetHeader("Authorization")
		if strings.Trim(token, "") == "" {
			loggerhelper.LogInfo("ERR_TOKEN_INVALID")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "ERR_TOKEN_DECODE_INVALID", "message": "Use /login route to create API Token"})
		}
		decodeTokenErr := tokenhelper.AccessByToken(c)
		if decodeTokenErr != nil {
			loggerhelper.LogInfo("ERR_DECODE_TOKEN_FAILED", decodeTokenErr)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "ERR_DECODE_TOKEN_FAILED", "message": decodeTokenErr.Error()})
		}
		c.Next()
	}
}
