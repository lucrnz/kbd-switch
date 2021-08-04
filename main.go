package main

import (
	"os"
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

func operationStartup() {
	err := changeKeyboardLayout(cfg.Layouts[cfg.CurrentLayout], cfg.SendNotification, cfg.NotificationCommand, cfg.NotificationMessage)
	if err != nil {
		panic(err)
	}
}

func operationCycle() {
	lTotal := len(cfg.Layouts)
	cfg.CurrentLayout++
	if cfg.CurrentLayout >= lTotal {
		cfg.CurrentLayout = 0
	}

	err := changeKeyboardLayout(cfg.Layouts[cfg.CurrentLayout], cfg.SendNotification, cfg.NotificationCommand, cfg.NotificationMessage)
	if err != nil {
		panic(err)
	}

	err = cfg.Save()
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "--startup" {
		operationStartup()
	} else {
		operationCycle()
	}
}
