package main

import (
	"fmt"
	"os"

	"github.com/jameskimmel/go_dyndns/config"
	"github.com/jameskimmel/go_dyndns/updater"
)

func main() {

	// check if config exists and can be read. Otherwise offer a wizard
	config.CheckConfig()

	// read the config
	config.ReadConfig()

	// start the updater
	updater.Updater()

	// write the IP(s) and the current time to the config file and exit
	config.UpdateConfig()
	fmt.Println("Everything done. Program will now exit.")
	os.Exit(0)

}
