package main

import (
	"strings"
	"encoding/json"
	"io/ioutil"
	"time"
	"./providers"
	"fmt"
)

type Provider interface {
	Update(record string, ip string)
}

type ProviderConfig struct {
	Domain       string
	Record       string
	ProviderInfo map[string]interface{}
}

type Config struct {
	Interval int
	Providers map[string]ProviderConfig
}

func main() {
	rawconf, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err.Error())
	}

	var config Config
	err = json.Unmarshal(rawconf, &config)
	if err != nil {
		panic(err.Error())
	}

	providers := make([]Provider, 0)
	for key, _ := range config.Providers {
		switch strings.ToLower(key) {
		case "gandi":
			apikey, ok := config.Providers[key].ProviderInfo["Apikey"]
			if ok {
			providers.Gandi(apikey)
			}
			break
		}
	}
	switch stirngs.ToLower()
}
