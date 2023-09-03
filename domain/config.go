package domain

// Config app config
type Config struct {
	DB Databases `yaml:"db"`
}

// Databases contains list of databases used in app
type Databases struct {
	Playground DBConfig `yaml:"playground"`
}

// DBConfig list all database connection requirement
type DBConfig struct {
	DSN string `yaml:"dsn"`
}
