package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	URL struct {
		CoreApi string `yaml:"core_api"`
		Public  string `yaml:"public"`
	} `yaml:"url"`
}

var ZKPasteConfig Config

func init() {
	f, err := os.Open("config.yml")
	if err != nil {
		ZKPasteConfig.URL.CoreApi = "https://core.zkpaste.com"
		ZKPasteConfig.URL.Public = "https://zkpaste.com"
	} else {
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&ZKPasteConfig)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
}
