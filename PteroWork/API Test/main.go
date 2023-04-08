package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var cpath string = "/workspaces/codespaces-blank/config.json"

type API struct {
	Key string `json:"Key"`
	Url string `json:"URL"`
}

type Output struct {
	Object string `json:"object"`
	Data   struct {
		Service string `json:"data.object"`
	}
}

type Data struct {
}

func main() {

	Unmarshal()
	// end = Output{}

}
func LoadConfig(jsonconf string) API {
	var GoConfig API
	raw, err := ioutil.ReadFile(jsonconf)
	if err != nil {
		log.Println("Error: Could not read config.")
	}
	json.Unmarshal(raw, &GoConfig)
	return GoConfig
}

func Unmarshal() []byte {
	var APIOut Output

	Config := LoadConfig(cpath)

	request, _ := http.NewRequest("GET", Config.Url, nil)
	request.Header.Add("cookie", "pterodactyl_session=eyJpdiI6InhIVXp5ZE43WlMxUU1NQ1pyNWRFa1E9PSIsInZhbHVlIjoiQTNpcE9JV3FlcmZ6Ym9vS0dBTmxXMGtST2xyTFJvVEM5NWVWbVFJSnV6S1dwcTVGWHBhZzdjMHpkN0RNdDVkQiIsIm1hYyI6IjAxYTI5NDY1OWMzNDJlZWU2OTc3ZDYxYzIyMzlhZTFiYWY1ZjgwMjAwZjY3MDU4ZDYwMzhjOTRmYjMzNDliN2YifQ%253D%253D")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+Config.Key)
	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &APIOut.Object)
	return data
}

func DiskSpace() {
	fmt.Println(disk.Service)
}
