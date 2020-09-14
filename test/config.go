package test

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiKey        string `json:"ApiKey"`
	DeletedApiKey string `json:"DeletedApiKey"`
	ExpiredToken  string `json:"ExpiredToken"`
}

func ReadFile(cfg *Config) error {
	b, err := ioutil.ReadFile("../config.json")

	if err != nil {
		return err
	}
	err = json.Unmarshal(b, cfg)
	if err != nil {
		return err
	}
	return nil
}
