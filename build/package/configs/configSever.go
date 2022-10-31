package configs

type ConfigsApp struct {
	AppName  string `env:"APP_NAME" env-default:"PDI-COBRANCA"`
	Port     string `env:"APP_PORT" env-default:"8081"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}
