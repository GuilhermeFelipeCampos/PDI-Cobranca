package configs

type ConfigsDB struct {
	DBUser     string `env:"USER" env-default:"postgres"`
	DBPort     string `env:"DB_PORT" env-default:"5432"`
	DBHost     string `env:"DB_HOST" env-default:"localhost"`
	DBPassword string `env:"DB_PASSWORD" env-default:"admin"`
	DBName     string `env:"DB_Name" env-default:"pdi_cobranca_db"`
}
