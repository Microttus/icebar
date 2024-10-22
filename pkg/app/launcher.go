package app

import (
	"fmt"
	"github.com/microttus/icebar/pkg/gui"
	"github.com/microttus/icebar/pkg/utils"
	"log"
	"os/exec"
)

func Launch(app *gui.App, appName string, execPath string) error {
	log.Printf("Launching application: %s (%s)", appName, execPath)
	cmd := exec.Command(execPath)
	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to launch %s: %v", appName, err)
		utils.ShowErrorDialog(app.Window, fmt.Sprintf("Failed to launch %s:\n%v", appName, err))

	}
	return nil
}
