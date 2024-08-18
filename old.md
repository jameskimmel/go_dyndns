	// set the URL to have the domain
	// updateurl = "https://update.dedyn.io/?hostname=" + config.Config.Domain


	// force update
	// fmt.Println("do you want to force an update? y/N")
	// fmt.Scanln(&input)
	// if (input == "y") || (input == "Y") {
	// 	forceUpate = true
	// } else {
	// 	forceUpate = false
	// }




	// write  config to file and exit probram

	// 	// for testing
	// 	lastSetIPv4 = "1.1.1.1"

	// 	configWrite := ConfigStruct{
	// 		LastSetIPv4: lastSetIPv4,
	// 	}

	// 	// write config struct to file
	// 	file, err := json.MarshalIndent(configWrite, "", " ")
	// 	if err != nil {
	// 		fmt.Println("something went wrong marshaling the config struct")
	// 		log.Fatal(err)
	// 	}
	// 	err = os.WriteFile("config.json", file, 0640)
	// 	if err != nil {
	// 		fmt.Println("something went wrong writing the config file")
	// 		log.Fatal(err)
	// 	}


	// if time.Duration(difference.Minutes()) > 5 {
// 	difference.Minutes()
// }

// fmt.Println(now.Sub(lastUpdate))

// if now.Sub(lastUpdate) > time.Duration(300) {
// 	fmt.Println("weniger als 5 minuten dazwischen")
// }

// if now.Sub(lastUpdate) < time.Duration(300) {
// 	fmt.Println("mehr als 5 minuten dazwischen")

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




// 	config := ConfigStruct{
	// 		Domain:            domain,
	// 		Token:             token,
	// 		EnableIPv4:        ipv4,
	// 		EnableIPv6:        ipv6,
	// 		HardcodedIPv4:     hardcodedIPv4,
	// 		HardcodedIPv6:     hardcodedIPv6,
	// 		MinMinutesBetween: minMinutesBetween,
	// 		LastUpdate:        LastUpdate,
	// 		LastSetIPv4:       LastSetIPv4,
	// 		LastSetIPv6:       LastSetIPv6,
	// 	}

	// 	file, err := os.ReadFile("config.json")
	// 	if err != nil {
	// 		fmt.Printf("could not read config.json. %s\n", err)
	// 		log.Fatal(err)
	// 	}

	// err = json.Unmarshal([]byte(file), &config)
	//
	//	if err != nil {
	//		fmt.Printf("could not marshal config.json: %s\n", err)
	//		log.Fatal(err)
	//	}
	//
	// ConfigSettings = config