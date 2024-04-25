package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"
)

var (
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"

	api = "https://siren.pp.ua/api/v3/alerts"
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
		fmt.Printf("[0/1] error getting API: %s\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[0/1] error reading JSON file: %s\n", err)
		return
	}

	var alerts []Alert
	err = json.Unmarshal(body, &alerts)
	if err != nil {
		fmt.Printf("[0/1] error unmarshaling JSON file: %s\n", err)
		return
	}

	if shouldShowCount() {
		countAlerts(alerts)
		return
	}

	printAlerts(alerts)
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

		fmt.Printf("%d. "+Red+Bold+"Повітряна тривога:"+Reset+" %s [%s] %s \n",
			num+1,
			alert.RegionName,
			modifiedName,
			strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
	}

	fmt.Printf(Bold+"\nСтаном на: %s\n"+Reset, time.Now().Format("2006-01-02 15:04:05"))
}

func countAlerts(alerts []Alert) {
	count := len(alerts)
	fmt.Printf("%d\n", count)
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
	fmt.Println(strings.TrimSpace(string(configDir)))
}

func shouldShowCount() bool {
	for _, arg := range os.Args {
		if arg == "-n" {
			return true
		} else if arg == "-c" {
			return true
		}
	}

	return false
}
