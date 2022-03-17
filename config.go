package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ConfigFile contains current status about the program
type ConfigFile struct {
	CurrentLayout       int
	Layouts             []string
	SendNotification    bool
	NotificationCommand string
	NotificationMessage string
	filePath            string
}

// Load is...
func (cfg *ConfigFile) Load() error {
	cfg.filePath = "./config.json"
	if _, err := os.Stat(cfg.filePath); err != nil {
		cfg.filePath = os.Getenv("HOME") + "/.config/kbd-switch.json"
		if _, err := os.Stat(cfg.filePath); err != nil {
			return fmt.Errorf("cannot find config file:\n\t%s", err.Error())
		}
	}
	cfgFileBytes, err := os.ReadFile(cfg.filePath)
	if err != nil {
		return fmt.Errorf("error reading config file:\n\t%s", err.Error())
	}
	if err = json.Unmarshal(cfgFileBytes, cfg); err != nil {
		return fmt.Errorf("error loading config file:\n\t%s", err.Error())
	}

	if cfg.Layouts == nil {
		return errors.New("config file is missing Layouts array")
	}

	if len(cfg.Layouts) < 2 {
		return errors.New("this program needs more than one layout to work correctly")
	}
	return nil
}

// Save is...
func (cfg *ConfigFile) Save() error {
	jsonBytes, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return fmt.Errorf("error saving config file:\n\t%s", err.Error())
	}
	if err = os.WriteFile(cfg.filePath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("error saving config file:\n\t%s", err.Error())
	}
	return nil
}
