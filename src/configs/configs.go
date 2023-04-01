package config

var Config *config = &config{}

type config struct {
	ApiConfigs struct {
		AppMode string `env:"APP_ENVIRONMENT_MODE" validate:"alpha"`
		AppPort uint16 `env:"APP_SERVER_PORT" validate:"numeric"`
	}
}

func (conf config) GetAppPort() uint16 {
	return conf.ApiConfigs.AppPort
}

func (conf config) GetAppMode() string {
	return conf.ApiConfigs.AppMode
}
