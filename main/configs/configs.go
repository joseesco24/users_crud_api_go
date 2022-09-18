package config

var Config *config = &config{}

type config struct {
	ApiConfigs struct {
		AppMode string `env:"APP_MODE" envDefault:"development" validate:"alpha"`
		AppPort uint16 `env:"APP_PORT" envDefault:"3064" validate:"numeric"`
	}
}

func (conf config) GetAppPort() uint16 {

	return conf.ApiConfigs.AppPort

}

func (conf config) GetAppMode() string {

	return conf.ApiConfigs.AppMode

}
