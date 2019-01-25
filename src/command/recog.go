package command

import (
    "bufio"
    "time"
	"os"
	"os/exec"
	"runtime"
)

var recog chan string

func InitRecog() {
	recog = make(chan string, 1)
	go GetRecog()
}
func GetRecog() {
    var cmd = _recog()
	var out,_ = cmd.StdoutPipe()
	cmd.Start()
	in := bufio.NewReader(out)
	for {
		result, _ := in.ReadString('\n')
		recog <- result
	}
}
func ReadWithTimeout(timeout time.Duration) string {
	select {
	case result := <-recog:
		return result
	case <-time.After(timeout * time.Millisecond):
		return ""
	}
}

func _recog() *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cscript", "//Nologo", os.TempDir() + "/recog.wsf")
	} else if runtime.GOOS == "darwin" {
		// FIXME やりかた
		return exec.Command("swift","/tmp/recog.swift")
	}
	return nil
}

