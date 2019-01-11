package main
import (
	"os"
	"bufio"
	"time"
)

var input chan string

func InitStdin() {
	input = make(chan string, 1)
	go GetInput()
}
func GetInput() {
	in := bufio.NewReader(os.Stdin)
	for {
		result, _ := in.ReadString('\n')
		input <- result
	}
}
func ReadWithTimeout(timeout time.Duration) string {
	select {
	case result := <-input:
		return result
	case <-time.After(timeout * time.Millisecond):
		return ""
	}
}


