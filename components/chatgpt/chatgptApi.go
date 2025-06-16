package chatgpt

import (
	configReader "bot/config"
	"io"
	"log"
	"net/http"
	"strings"
)

func RequestOpenAi(message string) string {
	api_key := configReader.Readconfig().APIKEY
	client := &http.Client{}
	var data = strings.NewReader(`{
        "model": "gpt-4o-mini",
        "input": "Привет!"
    }`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/responses", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyText)
}
