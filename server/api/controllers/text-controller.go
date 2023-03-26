package controllers

import (
	"os"
	"os/exec"
)

func RunPython() {
	cmd := exec.Command("python", "script.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	// log.Println(cmd.Run())
}
