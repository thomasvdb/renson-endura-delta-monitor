package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/robfig/cron"
	"gopkg.in/mailgun/mailgun-go.v1"
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
	URL          string             `json:"url"`
	MailSettings MailSettingsConfig `json:"mailSettings"`
}

type MailSettingsConfig struct {
	Domain   string `json:"domain"`
	APIKey   string `json:"apiKey"`
	FromName string `json:"fromName"`
	MailTo   string `json:"mailTo"`
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
				fmt.Println(monitorValues.MonitorValues[i].Value)

				sendStatusUpdate(configuration.MailSettings)
				resp.Body.Close()
			}
		}
	})
	c.Start()

	for {
	}
}

func sendStatusUpdate(mailSettings MailSettingsConfig) (string, error) {
	mg := mailgun.NewMailgun(mailSettings.Domain, mailSettings.APIKey, "")
	m := mg.NewMessage(
		mailSettings.FromName+"<"+mailSettings.Domain+">",
		"90 days are passed since the last cleaning of the filters!",
		"Renson Endura Delta filter cleaning notice",
		mailSettings.MailTo)

	_, id, err := mg.Send(m)
	return id, err
}
