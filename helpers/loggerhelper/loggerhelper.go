package loggerhelper

import (
	"codemax_assignment/models"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugar *zap.SugaredLogger

//logger initialisation
func Init(fileName string, loglevel zapcore.Level) {

	os.MkdirAll(path.Dir(models.AppConfig.AccessLogPath), os.ModePerm)
	f, _ := os.Create(models.AppConfig.AccessLogPath)
	gin.DefaultWriter = io.MultiWriter(f)

	os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: fileName,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		loglevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()
	sugar = logger.Sugar()
}

//info sugar logging
func LogInfo(args ...interface{}) {
	sugar.Info(args)
}

//error sugar logging
func LogError(args ...interface{}) {
	sugar.Error(args)
}
