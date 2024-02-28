package config

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBNAME   string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	SSL      string `env:"DB_SSL"`
	TimeZone string `env:"DB_TIMEZONE"`
}
