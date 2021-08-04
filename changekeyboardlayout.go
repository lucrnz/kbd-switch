package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func changeKeyboardLayout(newLayout string, sendNotification bool,
	notificationCommand string, notificationMessage string) error {
	fmt.Println("changing keyboard layout to: " + newLayout)
	cmd := exec.Command("setxkbmap", newLayout)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("cannot executing command setxkbmap:\n\t%s", err.Error())
	}

	if sendNotification {
		notificationMessage = strings.ReplaceAll(notificationMessage, "%layout%", newLayout)
		cmd := exec.Command(notificationCommand, notificationMessage)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("cannot executing NotificationCommand:\n\t%s", err.Error())
		}
	}
	return nil
}
