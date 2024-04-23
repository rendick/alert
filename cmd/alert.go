package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	// API
	api = "https://siren.pp.ua/api/v3/alerts"

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

func handleAlerts() {
	res, err := http.Get(api)
	if err != nil {
		fmt.Printf("[0/1] error getting an API: %s\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[0/1] error reading JSON file: %s\n", err)
		return
	}

	err = json.Unmarshal(body, &alerts)
	if err != nil {
		fmt.Printf("[0/1] error unmarshaling JSON file: %s\n", err)
		return
	}
	printAlerts(alerts)
}

func printAlerts(alerts []Alert) {
	replacement := map[string]string{
		"State":     "Область",
		"Community": "Громада",
	}

	for num, alert := range alerts {
		modifiedName := alert.RegionType

		for oldStr, newStr := range replacement {
			modifiedName = strings.ReplaceAll(modifiedName, oldStr, newStr)
		}

		fmt.Printf("%d."+Red+Bold+" Повітряна тривога:"+Reset+" %s [%s] %s \n",
			num+1,
			alert.RegionName,
			modifiedName,
			strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
	}
	fmt.Printf(Bold+"\nСтаном на: %s\n"+Reset, time.Now().Format("2006-01-02 15:04:05"))
	return
}
