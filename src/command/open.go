package command

import (
	"os/exec"
	"runtime"
)

func Open(url string) {
	command := "open"
	if runtime.GOOS == "windows" {
		command = "explorer"
	}
	exec.Command(command, url).Start()
}
