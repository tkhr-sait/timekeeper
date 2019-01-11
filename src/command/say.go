package command

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var winsaycommand string

func InitSay() {
	if runtime.GOOS == "windows" {
		tmpfile, _ := ioutil.TempFile("", "say")
		script := `WScript.CreateObject("SAPI.SpVoice").Speak(WScript.Arguments(0));`
		tmpfile.WriteString(script)
		tmpfile.Close()
		winsaycommand = tmpfile.Name() + ".js"
		os.Rename(tmpfile.Name(), winsaycommand)
	}

}
func DestroySay(){
	if runtime.GOOS == "windows" {
		os.Remove(winsaycommand)
	}
}
func _say(message string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cscript", "//Nologo", winsaycommand, message)
	} else {
		return exec.Command("say", message)
	}
}
func SayNoWait(message string) {
	_say(message).Start()
}
func Say(message string) {
	_say(message).Run()
}

