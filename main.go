package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Website       string `json:"Website"`
	Token         string `json:"Token"`
	IPv4          bool   `json:"IPv4"`
	IPv6          bool   `json:"IPv6"`
	IPv4Hardcoded string `json:"IPv4Hardcoded"`
	IPv6Hardcoded string `json:"IPv6Hardcoded"`
	WizardDone    bool   `json:"WizardDone"`
}

func main() {

	configJSON, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("could not marshal json file: %s\n", err)
		log.Fatal(err)
	}

	config := Config{}

	_ = json.Unmarshal([]byte(configJSON), &config)

	fmt.Println(config.Website)

}
