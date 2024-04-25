package main

import (
	"alert/rendick/cmd"
	"fmt"
	"runtime"
	"strings"
)

var (
	Bold = "\033[1m"
	Red  = "\033[31m"

	Reset = "\033[0m"
)

func main() {
	os_slice := []string{"linux", "freebsd", "netbsd", "openbsd", "dragonfly", "windows", "darwin", "android"}
	os_types := false
	for _, str := range os_slice {
		if str == runtime.GOOS {
			os_types = true
			break
		}
	}
	if os_types == true {
		cmd.Args()
	} else {
		fmt.Printf("You are running "+Bold+Red+"%s"+Reset+" instead of: "+Bold+"%s\n"+Reset, runtime.GOOS, strings.Join(os_slice, ", "))
		return
	}
}
