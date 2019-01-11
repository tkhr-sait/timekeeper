package main

import (
	"fmt"
	"command"
	"io"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func header(title string, index int, lines []string) {
	fmt.Printf("---[%s]---\n", title)
	if index-2 >= 0 {
		fmt.Printf("  %s\n", lines[index-2])
	}
	if index-1 >= 0 {
		fmt.Printf("  %s\n", lines[index-1])
	}
	fmt.Printf("> %s\n", lines[index])
	if index+1 < len(lines) {
		fmt.Printf("  %s\n", lines[index+1])
	}
	if index+2 < len(lines) {
		fmt.Printf("  %s\n", lines[index+2])
	}
	fmt.Printf("---------------------\n")
}
func trim(str string) string {
	return strings.TrimRight(strings.TrimRight(str, "\n"), "\r")
}

func main() {
	// check args
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " [scenario.txt]")
		os.Exit(-1)
	}

	infilepath := os.Args[1]
	file, err := os.Open(infilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	var lines []string
	scenario := bufio.NewReader(file)
	for {
		line, err := scenario.ReadString('\n')
		if err == io.EOF {
			break
		}
		// check format
		lines = append(lines, trim(line))
	}
	file.Close()

	// stdin
	input := InitStdin()
	// say
	command.InitSay()
	defer command.DestroySay()

	for index, line := range lines {
		command.ClearScreen()
		header(infilepath, index, lines)
		str := strings.Split(line, ",")
		switch str[0] {
		case "timer":
			message := str[1] + "　時間は" + str[2] + "分です。"
			fmt.Println(message)
			command.Say(message)
			i, _ := strconv.Atoi(str[2])
			timer(i*60, input)
		case "open":
			fmt.Println(str[1])
			command.Open(str[1])
		case "say":
			fmt.Println(str[1])
			command.Say(str[1])
		case "wait":
			fmt.Print(str[1] + " (Enter)")
			command.Say(str[1])
			for {
				ret := ReadWithTimeout(input)
				if ret != "" {
					break
				}
			}
		}
	}
}
