package auth

import (
	"fmt"
	"npsparkbrowser/file"
)

var apiKey string

func ApiKey() string {
	if apiKey == "" {
		value, err := loadApiKey()
		if err != nil {
			fmt.Println(err)
		}
		apiKey = value
	}
	return apiKey
}

func loadApiKey() (string, error) {
	value, err := file.Read()
	return value, err
}
