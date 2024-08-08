package main

import (
	"fmt"

	"github.com/jameskimmel/go_dyndns/config"
)

/*
type Config struct {
	Domain        string `json:"Domain"`
	Token         string `json:"Token"`
	IPv4Check     bool   `json:"IPv4Check"`
	IPv6Check     bool   `json:"IPv6Check"`
	IPv4Hardcoded string `json:"IPv4Hardcoded"`
	IPv6Hardcoded string `json:"IPv6Hardcoded"`
}

type Header map[string][]string

*/

func main() {
	fmt.Println(config.Hallo("hallo"))
	/*

		// read config file
		configJSON, err := os.ReadFile("config.json")
		if err != nil {
			fmt.Printf("could not marshal json file: %s\n", err)
			log.Fatal(err)
		}

		config := Config{}

		_ = json.Unmarshal([]byte(configJSON), &config)

		// create update urls for hardcoded

		updateURL := "https://update.dedyn.io/?hostname=" + config.Domain

		if config.IPv4Hardcoded != `no` {
			updateURL = updateURL + "&myipv4=" + config.IPv4Hardcoded
		}

		if config.IPv6Hardcoded != `no` {
			updateURL = updateURL + "&myipv6=" + config.IPv6Hardcoded

		}

		// do an IP check

		if config.IPv4Check {
			resp, err := http.Get("https://checkipv4.dedyn.io/")
			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()
			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)

			}
			ipv4 := string(resBody)
			//ipv4 := fmt.Sprintf("%s", resBody)
			updateURL = updateURL + "&myipv4=" + ipv4

		}

		if config.IPv6Check {
			resp, err := http.Get("https://checkipv6.dedyn.io/")
			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()
			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			ipv6 := string(resBody)
			//ipv6 := fmt.Sprintf("%s", resBody)
			updateURL = updateURL + "&myipv6=" + ipv6

		}

		// IP update

		req, err := http.NewRequest("GET", updateURL, nil)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)

		}
		req.Header.Set("Authorization", "Token "+config.Token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)

		}
		defer resp.Body.Close()


	*/

}
