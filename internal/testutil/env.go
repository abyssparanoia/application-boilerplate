package testutil

type environment struct {
	DefaultDBHost     string `env:"DB_HOST,required"`
	DefaultDBUser     string `env:"DB_USER,required"`
	DefaultDBDatabase string `env:"DB_DATABASE,required"`
	DefaultDBPassword string `env:"DB_PASSWORD,required"`
	DBLogEnabled      bool   `env:"DB_LOG_ENABLED" envDefault:"false"`
}
