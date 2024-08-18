package updater

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jameskimmel/go_dyndns/config"
)

var newIPv4 string
var newIPv6 string
var SkipUpdateIPv4 bool
var SkipUpdateIPv6 bool
var updateURL string

func Updater() {

	// check if the time difference to the last update is large enough
	// if not, exit the program
	checkTimeDifference()

	// If IPv4 is enabled and not hardcoded, do update routine
	if config.ConfigSettings.EnableIPv4 && config.ConfigSettings.HardcodedIPv4 == "" {
		getIPv4()
	} else {
		newIPv4 = config.ConfigSettings.HardcodedIPv4
	}

	// If IPv6 is enabled and not hardcoded, do update routine
	if config.ConfigSettings.EnableIPv6 && config.ConfigSettings.HardcodedIPv6 == "" {
		getIPv6()
	} else {
		newIPv6 = config.ConfigSettings.HardcodedIPv6
	}

	// compare the new IP4 with the last set ones. If nothing has changed, update can be skipped
	if config.ConfigSettings.EnableIPv4 && config.ConfigSettings.LastSetIPv4 == newIPv4 {
		fmt.Println("Your IPv4 has not changed and is still " + newIPv4)
		SkipUpdateIPv4 = true
	}

	// compare the new IP6 with the last set ones. If nothing has changed, update can be skipped
	if config.ConfigSettings.EnableIPv6 && config.ConfigSettings.LastSetIPv6 == newIPv6 {
		fmt.Println("Your IPv6 has not changed and is still " + newIPv6)
		SkipUpdateIPv6 = true
	}

	// if an IPv protocol is not enabled, we can skip it
	if !config.ConfigSettings.EnableIPv4 {
		SkipUpdateIPv4 = true
	}
	if !config.ConfigSettings.EnableIPv6 {
		SkipUpdateIPv4 = true
	}

	// if there is no need to update IPv4 or IPv6, we can close the program now
	if SkipUpdateIPv4 && SkipUpdateIPv6 {
		fmt.Println("No update needed, IP(s) have not changed. The program will now exit.")
		os.Exit(0)
	}

	// create the Update URL
	createURL()

	// do a http get request to update the records
	requestUpdate()

}

func getIPv4() {

	resp, err := http.Get("https://checkipv4.dedyn.io/")
	if err != nil {
		fmt.Println("Could not connect to checkipv4.dedyn.io! Make sure you have a working IPv4 connection or disable IPv4.")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	newIPv4 = string(resBody)

}

func getIPv6() {

	resp, err := http.Get("https://checkipv6.dedyn.io/")
	if err != nil {
		fmt.Println("Could not connect to checkipv6.dedyn.io! Make sure you have a working IPv6 connection or disable IPv6.")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	newIPv6 = string(resBody)

}

func checkTimeDifference() {
	now := time.Now()
	difference := now.Sub(config.ConfigSettings.LastUpdate)
	if difference.Minutes() < config.ConfigSettings.MinMinutesBetween {
		fmt.Println("last update was less than 5min ago. You can lower the value in the config.json file. Program will exit.")
		os.Exit(0)
	}
}

func createURL() {

	// set the URL to have the domain
	updateURL = "https://update.dedyn.io/?hostname=" + config.ConfigSettings.Domain

	// set IPv4
	if config.ConfigSettings.EnableIPv4 {
		updateURL = updateURL + "&myipv4=" + newIPv4
	}

	// set IPv6
	if config.ConfigSettings.EnableIPv6 {
		updateURL = updateURL + "&myipv6=" + newIPv6
	}

	fmt.Println("This is your update URL we will use: " + updateURL)
}

func requestUpdate() {
	req, err := http.NewRequest("GET", updateURL, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Token "+config.ConfigSettings.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// to do catch error

	// after a sucessful update, change LastSet IPs
	config.ConfigSettings.LastSetIPv4 = newIPv4
	config.ConfigSettings.LastSetIPv6 = newIPv6

}
