package shared

type Environment struct {
	Port            string `env:"PORT,required"`
	Envrionment     string `env:"ENV,required"`
	ProjectID       string `env:"PROJECT_ID,required"`
	CredentialsPath string `env:"GOOGLE_APPLICATION_CREDENTIALS,required"`
	MinLogSeverity  string `env:"MIN_LOG_SEVERITY,required"`
	DBHost          string `env:"DB_HOST,required"`
	DBUser          string `env:"DB_USER,required"`
	DBPassword      string `env:"DB_PASSWORD,required"`
	DBDatabase      string `env:"DB_DATABASE,required"`
}
