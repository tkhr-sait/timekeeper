package main
import (
	"os"
	"bufio"
	"time"
)

var timeout time.Duration = 1000

func InitStdin() chan string {
	input := make(chan string, 1)
	go GetInput(input)
	return input
}
func GetInput(input chan string) {
	in := bufio.NewReader(os.Stdin)
	for {
		result, _ := in.ReadString('\n')
		input <- result
	}
}
func ReadWithTimeout(input chan string) string {
	select {
	case result := <-input:
		return result
	case <-time.After(timeout * time.Millisecond):
		return ""
	}
}


