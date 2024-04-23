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
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s\n", version)
	} else {
		handleAlerts()
	}
}
