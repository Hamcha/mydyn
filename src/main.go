package main

import (
	"./providers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Provider interface {
	Update(record string, ip string) (error, bool)
}

type ProviderConfig struct {
	Domain       string
	Record       string
	ProviderInfo map[string]interface{}
}

type Config struct {
	Interval  time.Duration
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

	providerList := make(map[string]Provider)
	for key, _ := range config.Providers {
		switch strings.ToLower(key) {
		case "gandi":
			apikey, ok := config.Providers[key].ProviderInfo["Apikey"].(string)
			if ok {
				gandi, err := providers.Gandi(apikey, config.Providers[key].Domain)
				if err != nil {
					fmt.Println(now() + " [Gandi] Failed: " + err.Error())
				} else {
					providerList["Gandi"] = gandi
				}
			} else {
				fmt.Println(now() + " [Gandi] API key missing")
			}
			break
		}
	}

	oldip := getIp()
	queue := make(map[string]Provider, 0)
	for {
		// If no providers are left exit
		if len(providerList) < 1 {
			fmt.Println(now() + " [-] No providers in list, exiting..")
			return
		}

		ip := getIp()

		if ip != oldip {
			for key, provider := range providerList {
				queue[key] = provider
			}
			oldip = ip
		}

		if len(queue) > 0 {
			for key, provider := range queue {
				err, fatal := provider.Update(config.Providers[key].Record, ip)
				if err != nil {
					if fatal {
						fmt.Println(now() + " [" + key + "] ERROR - " + err.Error())
					} else {
						fmt.Println(now() + " [" + key + "] FATAL - " + err.Error())
						delete(providerList, key)
						delete(queue, key)
					}
				} else {
					fmt.Println(now() + " [" + key + "] Updated")
					delete(queue, key)
				}
			}
		}

		// Wait next interval
		time.Sleep(config.Interval * time.Second)
	}
}

func now() string {
	return time.Now().Format(time.Stamp)
}

func getIp() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		fmt.Println(now() + " [-] Can't get IP address!")
		panic(err.Error())
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}
