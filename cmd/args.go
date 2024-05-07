package cmd

import (
	"flag"
	"fmt"
)

var (
	version = "0.2v alert"
)

func Args() {
	showVersion := flag.Bool("version", false, "print the current version")
	showCount := flag.Bool("n", false, "print the amount of air raids alarms")
	showRegion := flag.Bool("c", false, "print the status of the air raid alarms according to your region in ~/.config/alert.conf")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s\n", version)
	} else if *showCount {
		handleAlerts()
	} else if *showRegion {
		currentRegion()
	} else {
		handleAlerts()
	}

}
