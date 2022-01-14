package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"golang.design/x/hotkey"
	"gopkg.in/yaml.v2"
)

func main() {
	confPath := os.Args[1]
	if confPath == "" {
		confPath = "./config.yaml"
	}

	log.Info().Str("configPath", confPath).Msg("starting echoClip")

	conf, err := NewYamlConfig().FromFile(confPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config file")
	}
	log.Info().Str("key", conf.Key).Strs("modifiers", conf.Modifiers).Msg("config unmarshaled successfully")

	var mods []hotkey.Modifier

	for i := range conf.Modifiers {
		m, err := strToMod(conf.Modifiers[i])
		if err != nil {
			log.Fatal().Err(err).Msg("failed to parse modifiers from config")
		}
		mods = append(mods, m)
	}

	var key hotkey.Key

	hk := hotkey.New(mods, key)

}

func strToMod(in string) (hotkey.Modifier, error) {
	switch in {
	case "win":
		return hotkey.ModWin, nil
	case "ctrl":
		return hotkey.ModCtrl, nil
	case "alt":
		return hotkey.ModAlt, nil
	case "shift":
		return hotkey.ModShift, nil
	default:
		return 0, fmt.Errorf("%v is not a valid modifier key", in)
	}
}

func NewYamlConfig() *YamlConfig {
	return &YamlConfig{}
}

type YamlConfig struct {
	ApiVersion string   `yaml:apiVersion`
	Key        string   `yaml:key`
	Modifiers  []string `yaml:modifiers`
}

func (yc *YamlConfig) FromFile(path string) (*YamlConfig, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, yc)
	if err != nil {
		return nil, err
	}

	err = yc.validate()
	if err != nil {
		return nil, err
	}

	return yc, nil
}

func (yc *YamlConfig) validate() error {
	if yc.ApiVersion != "echoClip:v1" {
		return fmt.Errorf("error, wrong config apiVersion")
	}
	return nil
}
