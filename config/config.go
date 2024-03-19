package config

import (
	"encoding/json"
	"log"
	"os"

)

type Config struct {
	Port        				string `json:"port"`
	DataBaseString				string `json:"dataBaseString"`
	SecretAccessKey             string `json:"secretAccessKey"`
	SecretAccessKeyLifeInHoures int    `json:"secretAccessKeyLifeInHoures"`
	SecretRefreshKey            string `json:"secretRefreshKey"`
	SecretRefresKeyLifeInHoures int    `json:"secretRefresKeyLifeInHoures"`
	K                           int    `json:"k"`
	J                           int    `json:"j"`
	ReportingPeriodInHoures     int    `json:"reportingPeriodInHoures"`
	TelegramToken				string `json:"telegramToken"`
	MailLogin 					string `json:"mailLogin"`
	MailPassword 				string `json:"mailPassword"`
}

func GetConfig() *Config {

	content, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &config
}