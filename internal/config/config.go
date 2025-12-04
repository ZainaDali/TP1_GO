package config

type AppConfig struct {
	Server struct {
		Port int
	} `mapstructure:"server"`
	Database struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"database"`
}

// Config est l'instance globale de notre configuration
var Config AppConfig

func InitConfig() {
	Config = AppConfig{
		Server: struct {
			Port int
		}{Port: 8080},
		Database: struct {
			Name string `mapstructure:"name"`
		}{Name: "crm_users.db"},
	}
}
