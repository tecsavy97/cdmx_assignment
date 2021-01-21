package main

import (
	"codemax_assignment/helpers/confighelper"
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/middleware"
	"codemax_assignment/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func main() {
	initErr := confighelper.InitConfig(models.GetConfigPath(), &models.AppConfig)
	loggerhelper.Init(models.AppConfig.LogPath, zapcore.InfoLevel)
	if initErr != nil {
		loggerhelper.LogError("FAILED_TO_INIT_APP_CONFIG:: ", initErr)
	}
	if startErr := startServer(); startErr != nil {
		loggerhelper.LogError("FAILED_TO_START_SERVER:: ", startErr, "at Port: ", models.AppConfig.AppPort)
	}
}

// startServer - start server with app port
func startServer() error {
	router := gin.Default()
	// enabling cors for remote hits
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	router.Use(cors.New(md))
	router.GET("/", checkServerStatus())
	middleware.InitMiddleWare(router)
	s := &http.Server{
		Addr:    models.AppConfig.AppPort,
		Handler: router,
	}
	loggerhelper.LogInfo("Server Starting at Port ::", models.AppConfig.AppPort)
	if startServerErr := s.ListenAndServe(); startServerErr != nil {
		return startServerErr
	}
	return nil
}

// checkServerStatus- to check status of server if running or not
func checkServerStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is Running!"})
	}
}
