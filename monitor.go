package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

type MonitorValues struct {
	MonitorValues []MonitorValue `json:"ModifiedItems"`
}

type MonitorValue struct {
	Name  string `json:"Name"`
	Index string `json:"Index"`
	Value string `json:"Value"`
}

type Config struct {
	URL string `json:"url"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {
	fmt.Println("Starting Renson Endura Delta monitor...")
	fmt.Println("")

	c := cron.New()
	c.AddFunc("@daily", func() {
		var configuration = LoadConfiguration("config.json")
		resp, err := http.Get(configuration.URL)
		if err != nil {
			fmt.Printf(err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var monitorValues MonitorValues
		json.Unmarshal(body, &monitorValues)

		for i := 0; i < len(monitorValues.MonitorValues); i++ {
			if monitorValues.MonitorValues[i].Name == "Filter remaining time" {
				fmt.Println("Remaining time for filter: " + monitorValues.MonitorValues[i].Value + " days")
			}
		}
	})
	c.Start()

	for {
	}
}
