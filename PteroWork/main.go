package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

fenv := "/workspaces/codespaces-blank/PteroWork/.env"

type ENV struct {
	Key string `json:"Key"`
	Url string `json:"URL"`
}

func setenv(fenv string) {
	var env ENV
	raw, err := ioutil.ReadFile(fenv)
	if err != nil {
		log.Println("Error: Could not read config.")
	}
	json.Unmarshal(raw, &env)
	return env
}

func main() {

	setenv()
	api, err := cloudflare.New(ENV.Key)

}
