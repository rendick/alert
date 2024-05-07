package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"
)

const (
	red   = "\033[31m"
	bold  = "\033[1m"
	reset = "\033[0m"
	api   = "https://siren.pp.ua/api/v3/alerts"
)

type Alert struct {
	RegionName    string `json:"regionName"`
	LastUpdate    string `json:"lastUpdate"`
	RegionEngName string `json:"regionEngName"`
	RegionId      string `json:"regionId"`
	RegionType    string `json:"regionType"`
}

func handleAlerts() {
	res, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error getting API: %s\n", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	var alerts []Alert
	err = json.Unmarshal(body, &alerts)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON file: %s\n", err)
		return
	}

	if shouldShowCount() {
		countAlerts(alerts)
	} else {
		printAlerts(alerts)
	}
}

func printAlerts(alerts []Alert) {
	replacements := map[string]string{
		"State":     "Область",
		"Community": "Громада",
	}

	for num, alert := range alerts {
		modifiedName := alert.RegionType
		for oldStr, newStr := range replacements {
			modifiedName = strings.ReplaceAll(modifiedName, oldStr, newStr)
		}

		fmt.Printf("%d. "+bold+red+"Повітряна тривога:"+reset+" %s [%s] %s\n",
			num+1,
			alert.RegionName,
			modifiedName,
			strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
	}

	fmt.Printf(bold+"\nСтаном на: %s\n"+reset, time.Now().Format("2006-01-02 15:04:05"))
}

func countAlerts(alerts []Alert) {
	count := len(alerts)
	fmt.Println(count)
}

func shouldShowCount() bool {
	for _, arg := range os.Args {
		if arg == "-n" {
			return true
		}
	}
	return false
}

func currentRegion() {
	userName, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configDir, err := ioutil.ReadFile(userName.HomeDir + "/.config/alert.conf")
	if err != nil {
		log.Fatal(err)
	}

	targetRegion := strings.TrimSpace(string(configDir))

	res, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error getting API: %s\n", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	var alerts []Alert
	err = json.Unmarshal(body, &alerts)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON file: %s\n", err)
		return
	}

	replacements := map[string]string{
		"State":     "Область",
		"Community": "Громада",
	}

	var filteredAlerts []Alert
	for _, alert := range alerts {
		if alert.RegionName == targetRegion {
			filteredAlerts = append(filteredAlerts, alert)
		}
	}

	for _, alert := range filteredAlerts {
		modifiedName := alert.RegionType
		for oldStr, newStr := range replacements {
			modifiedName = strings.ReplaceAll(modifiedName, oldStr, newStr)
		}

		fmt.Printf(bold+red+"Повітряна тривога:"+reset+" %s %s\n",
			alert.RegionName,
			strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
	}
}
