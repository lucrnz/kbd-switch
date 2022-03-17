package main

import (
	"os"
)

var (
	cfg ConfigFile
)

func init() {
	if err := cfg.Load(); err != nil {
		panic(err)
	}
}

func operationStartup() {
	if err := changeKeyboardLayout(cfg.Layouts[cfg.CurrentLayout],
		cfg.SendNotification,
		cfg.NotificationCommand,
		cfg.NotificationMessage); err != nil {
		panic(err)
	}
}

func operationCycle() {
	lTotal := len(cfg.Layouts)
	cfg.CurrentLayout++

	if cfg.CurrentLayout >= lTotal {
		cfg.CurrentLayout = 0
	}

	if err := changeKeyboardLayout(cfg.Layouts[cfg.CurrentLayout],
		cfg.SendNotification,
		cfg.NotificationCommand,
		cfg.NotificationMessage); err != nil {
		panic(err)
	}

	if err := cfg.Save(); err != nil {
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
