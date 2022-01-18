package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"git.tcp.direct/kayos/sendkeys"
	"github.com/rs/zerolog/log"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
	"gopkg.in/yaml.v2"
)

func main() {
	dirPath := strings.Replace(filepath.Dir(os.Args[0]), `\`, "/", -1)

	confPath := fmt.Sprintf("%v/config.yaml", dirPath)
	if len(os.Args) > 1 {
		confPath = os.Args[1]
	}

	log.Info().Str("programPath", os.Args[0]).Str("configPath", confPath).Msg("starting echoClip")

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
	key, err = strToKey(conf.Key)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse hotkey from config")
	}

	k, err := sendkeys.NewKBWrapWithOptions(sendkeys.Noisy)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create kb object")
		return
	}

	var clipString string
	for {
		mainthread.Init(makeHKThread(mods, key))
		clipString = string(clipboard.Read(clipboard.FmtText))
		if clipString != "" {
			break
		}
		log.Warn().Msg("clipboard was empty")
	}

	if err = k.Type(clipString); err != nil {
		log.Warn().Err(err).Msg("failed to send key presses")
	}
}

func makeHKThread(mods []hotkey.Modifier, key hotkey.Key) func() {
	return func() {
		hk := hotkey.New(mods, key)
		err := hk.Register()
		if err != nil {
			log.Warn().Err(err).Msg("error while registering hotkey")
			return
		}
		log.Info().Str("hotkey", hk.String()).Msg("hotkey registered")

		<-hk.Keydown()
		log.Info().Str("hotkey", hk.String()).Msg("hotkey pressed")
		<-hk.Keyup()
		log.Info().Str("hotkey", hk.String()).Msg("hotkey release")
		hk.Unregister()
		log.Info().Str("hotkey", hk.String()).Msg("hotkey unregistered")
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
	// if yc.ApiVersion != "echoClip/v1" {
	// 	return fmt.Errorf("error, wrong config apiVersion")
	// }
	return nil
}
