package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var input string
var domain string
var token string
var ipv4 bool
var ipv6 bool
var hardcodedIPv4 string
var hardcodedIPv6 string
var minMinutesBetween float64
var LastUpdate time.Time
var LastSetIPv4 string
var LastSetIPv6 string

var ConfigSettings ConfigStruct

type ConfigStruct struct {
	Domain            string    `json:"Domain"`
	Token             string    `json:"Token"`
	EnableIPv4        bool      `json:"EnableIPv4"`
	EnableIPv6        bool      `json:"EnableIPv6"`
	HardcodedIPv4     string    `json:"HardcodedIPv4"`
	HardcodedIPv6     string    `json:"HardcodedIPv6"`
	MinMinutesBetween float64   `json:"MinMinutesBetween"`
	LastUpdate        time.Time `json:"LastUpdate"`
	LastSetIPv4       string    `json:"LastSetIPv4"`
	LastSetIPv6       string    `json:"LastSetIPv6"`
}

func CheckConfig() {
	_, err := os.ReadFile("config.json")
	if err != nil {
		// No config file found. Ask if user wants to run the config wizard
		fmt.Println("Found no config file. Do you want to create one using the wizard? Y/n")
		fmt.Scanln(&input)

		// If user inserts Y, y or the default empty, it will start the wizard. Else it will exit.
		if !(input == "y") && !(input == "Y") && !(input == "") {
			fmt.Println("Ok. Create a config file first and restart go_dyndns. Program will now exit.")
			os.Exit(0)
		} else {
			wizard()
		}
	}
}

func ReadConfig() {

	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("could not read config.json. %s\n", err)
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &ConfigSettings)
	if err != nil {
		fmt.Printf("could not marshal config.json: %s\n", err)
		log.Fatal(err)
	}

}

func UpdateConfig() {

	// use current time as last update
	ConfigSettings.LastUpdate = time.Now()

	// write config struct to file
	file, err := json.MarshalIndent(ConfigSettings, "", " ")
	if err != nil {
		fmt.Println("something went wrong marshaling the config struct")
		log.Fatal(err)
	}
	err = os.WriteFile("config.json", file, 0640)
	if err != nil {
		fmt.Println("something went wrong writing the config file")
		log.Fatal(err)
	}
	fmt.Println("sucessfully written changes to config file")

}

func wizard() {

	// check if writing a config file is possible before starting
	// the wizard
	configFile, err := os.Create("config.json")
	if err != nil {
		fmt.Println("unable to write a config file :()")
		log.Fatal(err)
	}
	configFile.Close()

	// Start with the Wizard

	// ask for the domain
	fmt.Println("Please insert the domain you want to update and press enter. For example my.domain.com or mydomain.com")
	fmt.Scanln(&domain)

	// ask for the token
	fmt.Println("Please insert the token and press enter. This is NOT your deSEC Password!")
	fmt.Scanln(&token)

	// ask if IPv4 should be enabled
	fmt.Println("Do you want to enable IPv4? Y/n")
	fmt.Scanln(&input)
	if !(input == "n") || (input == "N") {
		ipv4 = true
	} else {
		ipv4 = false
	}

	// ask if IPv6 should be enabled
	fmt.Println("Do you want to enable IPv6? Y/n")
	fmt.Scanln(&input)
	if !(input == "n") || (input == "N") {
		ipv6 = true
	} else {
		ipv6 = false
	}

	// // hardcoded IPv4
	// fmt.Println("Instead of checking your IPv4, you can also set a hardcoded IPv4. This is normally not needed. Press enter to skip")
	// fmt.Scanln(&input)
	// if (input == "n") || (input == "N") || (input == "") {
	// 	hardcodedIPv4 = ""
	// } else {
	// 	hardcodedIPv4 = input
	// }

	// // hardcoded IPv6
	// fmt.Println("Instead of checking your IPv6, you can also set a hardcoded IPv6. This is normally not needed. Press enter to skip")
	// fmt.Scanln(&input)
	// if (input == "n") || (input == "N") || (input == "") {
	// 	hardcodedIPv6 = ""
	// } else {
	// 	hardcodedIPv6 = input
	// }

	// default value for min minutes in between
	minMinutesBetween = 5

	configWriteWizard := ConfigStruct{
		Domain:            domain,
		Token:             token,
		EnableIPv4:        ipv4,
		EnableIPv6:        ipv6,
		HardcodedIPv4:     hardcodedIPv4,
		HardcodedIPv6:     hardcodedIPv6,
		MinMinutesBetween: minMinutesBetween,
		LastUpdate:        LastUpdate,
		LastSetIPv4:       LastSetIPv4,
		LastSetIPv6:       LastSetIPv6,
	}

	// write config struct to file
	file, err := json.MarshalIndent(configWriteWizard, "", " ")
	if err != nil {
		fmt.Println("something went wrong marshaling the config struct")
		log.Fatal(err)
	}
	err = os.WriteFile("config.json", file, 0640)
	if err != nil {
		fmt.Println("something went wrong writing the config file")
		log.Fatal(err)
	}

}
