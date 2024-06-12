package server

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// Constructor
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":6987",
		LogLevel:    "debug",
		DatabaseURL: "admin:Erunda228@tcp(db:3306)/journal",
	}
}
