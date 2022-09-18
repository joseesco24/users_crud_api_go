package config

var Config *config = &config{}

type config struct {
	ApiConfigs struct {
		AppPort uint16 `env:"APP_PORT" validate:"numeric"`
		AppMode string `env:"APP_MODE" validate:"alpha"`
	}
}

func (conf config) GetAppPort() uint16 {

	return conf.ApiConfigs.AppPort

}

func (conf config) GetAppMode() string {

	return conf.ApiConfigs.AppMode

}
