package routes

import (
	"codemax_assignment/handlers"

	"github.com/gin-gonic/gin"
)

//InitEmailRoutes- initialise email routes
func InitEmailRoutes(o, r *gin.RouterGroup) {
	r.POST("/send-email", handlers.SendEmailHandler())
}
