package confighelper

import (
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/models"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap/zapcore"
)

func InitConfig(path string, v interface{}) error {
	_, decodeErr := toml.DecodeFile(models.GetConfigPath(), &models.AppConfig)
	if decodeErr != nil {
		return decodeErr
	}
	loggerhelper.Init(models.AppConfig.LogPath, zapcore.InfoLevel)
	return nil
}
