package cmd

import (
	"flag"
	"fmt"
)

var (
	version = "0.2v alert"
)

func Args() {
	showVersion := flag.Bool("version", false, "Print the current version")
	showCount := flag.Bool("n", false, "Print the ")
	showRegion := flag.Bool("c", false, "print")
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
