package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	// API
	api = "https://siren.pp.ua/api/v3/alerts"

	// Information
	help_eng = ` 
Usage: alert [EMPRY or OPTION]

--english or -en: translation of output into English [won't work for some time]

--help or -h:     help 
--version or -v:  version

Air raid alert in Ukraine CLI: <https://github.com/rendick/alert/>
`
	version = "0.1v Alpha"

	// Color
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

type Alert struct {
	RegionName    string `json:"regionName"`
	LastUpdate    string `json:"lastUpdate"`
	RegionEngName string `json:"regionEngName"`
	RegionId      string `json:"regionId"`
	RegionType    string `json:"regionType"`
}

var alerts []Alert

func main() {
	if len(os.Args) != 2 {
		handleAlerts()
	} else {
		handleCommands(os.Args[1])
	}
}

func handleAlerts() {
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if len(os.Args) != 2 {
		body, err := io.ReadAll(res.Body)
		if len(os.Args) != 2 {
			if err != nil {
				log.Fatal("Помилка читання JSON файлу: ", err)
			}
		} else {
			if err != nil {
				log.Fatal("Error reading JSON file: ", err)
			}
		}

		err = json.Unmarshal(body, &alerts)
		if len(os.Args) != 2 {
			if err != nil {
				log.Fatal("Помилка перетворення даних у структуру у JSON файлі: ", err)
				return
			}
		} else {
			if err != nil {
				log.Fatal("Error unmarshaling JSON file: ", err)
				return
			}
		}
		printAlerts(alerts)
	}
}

func printAlerts(alerts []Alert) {

	for num, alert := range alerts {
		if alert.RegionType == "State" {
			alert.RegionType = "Область"
		} else if alert.RegionType == "Community" {
			alert.RegionType = "Громада"
		}
		fmt.Printf("%d."+Red+Bold+" Повітряна тривога:"+Reset+" %s [%s] %s \n",
			num+1,
			alert.RegionName,
			alert.RegionType,
			strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
	}
	fmt.Printf(Bold+"\nСтаном на: %s\n"+Reset, time.Now().Format("2006-01-02 15:04:05"))
	os.Exit(0)
}

func handleCommands(command string) {
	switch command {
	case "--help", "-help", "-h", "--h":
		fmt.Println(help_eng)
	case "--version", "-version", "-v", "--v":
		fmt.Println(version)
	}
}
