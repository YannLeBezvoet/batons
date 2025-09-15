package config

import (
	"encoding/json"
	"errors"
	"os"
)

type ConfigStruct struct {
	MoveCursorLeft  rune `json:"MoveCursorLeft"`
	MoveCursorRight rune `json:"MoveCursorRight"`
	MoveCursorUp    rune `json:"MoveCursorUp"`
	MoveCursorDown  rune `json:"MoveCursorDown"`
}

func InitConfig(configPath string) (ConfigStruct, error) {
	exists, err := CheckStartingConfig(configPath)
	if err != nil {
		return ConfigStruct{}, err
	}
	if !exists {
		err := CreateDefaultConfig(configPath)
		if err != nil {
			return ConfigStruct{}, err
		}
	}
	configStruct, err := LoadConfig(configPath)
	if err != nil {
		return ConfigStruct{}, err
	}
	return configStruct, nil
}

func CheckStartingConfig(configPath string) (bool, error) {
	if _, err := os.Stat(configPath); err == nil {
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil

	} else {
		return false, err
	}
}

func CreateDefaultConfig(configPath string) error {
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	myJson := `{"MoveCursorLeft": 113, "MoveCursorRight": 100, "MoveCursorUp": 122, "MoveCursorDown": 115}`
	_, err = file.WriteString(myJson)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(configPath string) (ConfigStruct, error) {
	var config ConfigStruct

	file, err := os.Open(configPath)
	if err != nil {
		return ConfigStruct{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return ConfigStruct{}, err
	}
	return config, nil
}

func SaveConfig(configPath string, config ConfigStruct) error {
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		return err
	}
	return nil
}
