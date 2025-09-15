package config

import (
	"errors"
	"os"
)

func CheckStartingConfig(configPath string) (bool, error) {
	if _, err := os.Stat("/path/to/whatever"); err == nil {
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil

	} else {
		return false, err
	}
}
