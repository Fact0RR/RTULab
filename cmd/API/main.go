package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Fact0RR/RTULab/internal/data"
)

func getConfig() *data.Config {

	content, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config data.Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &config
}

func serData() (*[]byte, error) {
	data := []byte(`{"foo":"bar","DS": 1 }`)
	r := bytes.NewReader(data)
	resp, err := http.Post("https://recruit.rtuitlab.dev/serialize", "application/json", r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return &b, nil
}

func deserData(by *[]byte) error {
	log.Println(string(*by))
	r := bytes.NewReader(*by)
	resp, err := http.Post("https://recruit.rtuitlab.dev/deserialize", "application/json", r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(b))
	return nil
}

func migr(config *data.Config) {
	
}

func main() {
	config := getConfig()
	b, _ := serData()

	deserData(b)
	//log.Println(config.DataBaseURL)

	migr(config)
}
