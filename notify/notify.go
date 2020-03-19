package notify

import (
	"os"
	"os/exec"
)

// SendLinux sends notifications to systems running linux via notify-send
func SendLinux(title string, message string, icon string, urgency string, delay string) interface{} {
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
