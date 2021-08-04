package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	cfg ConfigFile
)

func init() {
	err := cfg.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	lTotal := len(cfg.Layouts)

	if lTotal <= 1 {
		panic(errors.New("this program needs more than one layout to work correctly"))
	}

	cfg.CurrentLayout++
	if cfg.CurrentLayout >= lTotal {
		cfg.CurrentLayout = 0 // use the First one
	}

	cmd := exec.Command("setxkbmap", cfg.Layouts[cfg.CurrentLayout])
	err := cmd.Run()
	if err != nil {
		panic(fmt.Errorf("cannot executing command setxkbmap:\n\t%s", err.Error()))
	}

	if cfg.SendNotification {
		cfg.NotificationMessage = strings.Replace(cfg.NotificationMessage,
			"%layout%", cfg.Layouts[cfg.CurrentLayout], 0)
		cmd := exec.Command(cfg.NotificationCommand, cfg.NotificationMessage)
		err := cmd.Run()
		if err != nil {
			panic(fmt.Errorf("cannot executing NotificationCommand:\n\t%s", err.Error()))
		}
	}

	err = cfg.Save()
	if err != nil {
		panic(err)
	}
}
