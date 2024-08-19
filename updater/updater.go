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
var skipUpdateIPv4 bool
var skipUpdateIPv6 bool
var updateURL string

func Updater() {

	// check if the time difference to the last update is large enough
	// if not, exit the program
	checkTimeDifference()

	// If IPv4 is enabled and not hardcoded, check for IPv4 changes
	if config.ConfigSettings.EnableIPv4 && config.ConfigSettings.HardcodedIPv4 == "" {
		getIPv4()
	} else {
		newIPv4 = config.ConfigSettings.HardcodedIPv4
	}

	// If IPv6 is enabled and not hardcoded IPv6, check for IPv6 changes
	if config.ConfigSettings.EnableIPv6 && config.ConfigSettings.HardcodedIPv6 == "" {
		getIPv6()
	} else {
		newIPv6 = config.ConfigSettings.HardcodedIPv6
	}

	// compare the new IP4 with the last set one. If nothing has changed, update can be skipped
	if config.ConfigSettings.EnableIPv4 && config.ConfigSettings.LastSetIPv4 == newIPv4 {
		fmt.Println("Your IPv4 has not changed and is still " + newIPv4)
		skipUpdateIPv4 = true
	}

	// compare the new IP6 with the last set one. If nothing has changed, update can be skipped
	if config.ConfigSettings.EnableIPv6 && config.ConfigSettings.LastSetIPv6 == newIPv6 {
		fmt.Println("Your IPv6 has not changed and is still " + newIPv6)
		skipUpdateIPv6 = true
	}

	// if an IPv protocol is not enabled, we can also skip it
	if !config.ConfigSettings.EnableIPv4 {
		skipUpdateIPv4 = true
	}
	if !config.ConfigSettings.EnableIPv6 {
		skipUpdateIPv4 = true
	}

	// if there is no need to update IPv4 or IPv6, we can close the program now
	if skipUpdateIPv4 && skipUpdateIPv6 {
		fmt.Println("No update needed, IP(s) have not changed. The program will now exit.")
		os.Exit(0)
	}

	// create the update URL for the update request
	createURL()

	// Do the actual update
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
	// Checks if there is a big enough time difference for another update. The wizard uses 5min as the default.
	now := time.Now()
	difference := now.Sub(config.ConfigSettings.LastUpdate)
	if difference.Minutes() < config.ConfigSettings.MinMinutesBetween {
		fmt.Println("last update was less than", config.ConfigSettings.MinMinutesBetween, "min ago. You can lower the value in the config.json file. Program will exit.")
		os.Exit(0)
	}
}

func createURL() {

	// set the domain
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
	if resp.StatusCode != 200 {
		fmt.Println("Something went wrong. DeSEC returned an error when asking for an update")
		log.Fatal(resp.Status)

	}

	if resp.StatusCode == 200 {
		fmt.Println("Successfully set the new IP(s) on DeSEC!")

	}
	// after a sucessful update, change LastSet IPs
	config.ConfigSettings.LastSetIPv4 = newIPv4
	config.ConfigSettings.LastSetIPv6 = newIPv6

}
