package config

import (
	"encoding/json"
	"errors"
	"os"
)

func InitConfig(configPath string) (map[string]string, error) {
	exists, err := CheckStartingConfig(configPath)
	if err != nil {
		return nil, err
	}
	if !exists {
		err := CreateDefaultConfig(configPath)
		if err != nil {
			return nil, err
		}
	}
	config, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func CheckStartingConfig(configPath string) (bool, error) {
	if _, err := os.Stat("/path/to/whatever"); err == nil {
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
	myJson := `{"MoveCursorLeft": "q", "MoveCursorRight": "d", "MoveCursorUp": "w", "MoveCursorDown": "s"}`
	_, err = file.WriteString(myJson)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(configPath string) (map[string]string, error) {
	config := make(map[string]string)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
