package configReader

type Config struct {
	APIURL   string
	APIKEY   string
	BOTTOKEN string
}

func Readconfig() Config {
	// api_url, _ := os.LookupEnv("APIURL")
	// api_key, _ := os.LookupEnv("APIKEY")
	// bot_token, _ := os.LookupEnv("TELEGRAMBOTTOKEN")
	api_url := "https://api.ru/openai/v1/response/"
	api_key := "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"
	bot_token := "6992821026:AAE-D-3AQ9EDJClUyXyL7XBEpZKu2FZU8U4"
	return Config{
		APIURL:   api_url,
		APIKEY:   api_key,
		BOTTOKEN: bot_token,
	}
}
