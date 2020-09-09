package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	ApiKey       string `json:"ApiKey"`
	ExpiredToken string `json:"ExpiredToken"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config) {
	b, err := ioutil.ReadFile("../config.json")

	if err != nil {
		processError(err)
	}
	err = json.Unmarshal(b, cfg)
	if err != nil {
		processError(err)
	}
}
