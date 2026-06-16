package config

import "os"

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		// Go's MySQL driver safely parses password special characters like @ and $ natively
		dbUrl = "root:@sam12$$@tcp(localhost:3306)/userdb?parseTime=true"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	return &Config{DBUrl: dbUrl, Port: port}
}
