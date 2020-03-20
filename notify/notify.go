package notify

import (
	"os"
	"os/exec"

	"gopkg.in/toast.v1"
)

// SendLinux sends notifications to systems running linux via notify-send
func SendLinux(title string, message string, icon string, urgency string, delay string) error {
	currentDir, direrr := os.Getwd()
	if direrr != nil {
		return direrr
	}
	dir := currentDir + "/" + icon

	cmd := exec.Command("notify-send", "-i", dir, "-u", urgency, "-t", delay, title, message)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// SendWin sends notifications to systems running windows via Toast
func SendWin(title string, message string, icon string) error {
	notif = toast.Notification{
		AppID: "GoWeather"
		Title: title
		Messsage: message
		Actions: []toast.Action{}
	}
	error := notif.Push()
	if err != nil {
		return err
	}
	return nil
}
