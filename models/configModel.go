package models

//AppConfig- to be used globally
var AppConfig Config

// Config - to read app config toml
type Config struct {
	AppPort       string
	AccessLogPath string
	LogPath       string
	EmailConfPath string
	PublicKey     string
	SecretKey     string
	Host          string
}

//GetConfigPath -  Get toml file path
func GetConfigPath() string {
	return "./config/config.toml"
}
