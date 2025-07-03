package chatgpt

import (
	configReader "bot/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func RequestOpenAi(message string) string {
	api_key := configReader.Readconfig().APIKEY
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{
  "model": "deepseek/deepseek-r1-0528:free",
  "messages": [
    {
      "role": "user",
      "content": "%s"
    }
  ]
  
}`, message))
	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", data)
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
