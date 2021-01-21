package routes

import (
	"codemax_assignment/handlers"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(o, r *gin.RouterGroup) {
	o.POST("/login", handlers.GenerateToken())
}
