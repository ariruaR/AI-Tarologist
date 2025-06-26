package configReader

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

type Config struct {
	APIURL   string
	APIKEY   string
	BOTTOKEN string
}

func init() {
	if err := env.Load("config/.env"); err != nil {
		log.Println(err)
	}
}
func Readconfig() Config {
	api_url, _ := os.LookupEnv("APIURL")
	api_key, _ := os.LookupEnv("APIKEY")
	bot_token, _ := os.LookupEnv("TELEGRAMBOTTOKEN")
	return Config{
		APIURL:   api_url,
		APIKEY:   api_key,
		BOTTOKEN: bot_token,
	}
}
