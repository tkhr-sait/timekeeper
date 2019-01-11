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

func Header(title string, index int, lines []string) {
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
func Trim(str string) string {
	return strings.TrimRight(strings.TrimRight(str, "\n"), "\r")
}
func ErrorExit(message string) {
	fmt.Println(message)
	os.Exit(-1)
}
func main() {
	// check args
	if len(os.Args) != 2 {
		ErrorExit("usage: " + os.Args[0] + " [scenario.txt]")
	}

	infilepath := os.Args[1]
	file, err := os.Open(infilepath)
	if err != nil {
		ErrorExit(err.Error())
	}
	var lines []string
	scenario := bufio.NewReader(file)
	for {
		line, err := scenario.ReadString('\n')
		if err == io.EOF {
			break
		}
		// check format
		str := strings.Split(Trim(line), ",")
		switch str[0] {
		case "timer":
			if len(str) != 3 {
				ErrorExit("timerコマンド形式不正 timer,[message],[minutes] 入力=" + Trim(line))
			}
			_, err := strconv.Atoi(str[2])
			if err != nil {
				ErrorExit("[minutes]は数値 timer,[message],[minutes] 入力=" + Trim(line))
			}
		case "open":
			if len(str) != 2 {
				ErrorExit("openコマンド形式不正 open,[url] 入力=" + Trim(line))
			}
		case "say":
			if len(str) != 2 {
				ErrorExit("sayコマンド形式不正 say,[message] 入力=" + Trim(line))
			}
		case "wait":
			if len(str) != 2 {
				ErrorExit("waitコマンド形式不正 wait,[message] 入力=" + Trim(line))
			}
		default:
			ErrorExit("指定可能なコマンド timer,open,say,wait 入力=" + Trim(line))
		}
		lines = append(lines, Trim(line))
	}
	file.Close()

	// stdin
	InitStdin()
	command.InitRecog()
	// say
	command.InitSay()
	defer command.DestroySay()

	for index, line := range lines {
		command.ClearScreen()
		Header(infilepath, index, lines)
		str := strings.Split(line, ",")
		switch str[0] {
		case "timer":
			message := str[1] + "　時間は" + str[2] + "分です。"
			fmt.Println(message)
			command.Say(message)
			i, _ := strconv.Atoi(str[2])
			Timer(i*60)
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
				ret := ReadWithTimeout(500)
				if ret != "" {
					break
				}
				ret = Trim(command.ReadWithTimeout(500))
				if ret == "s" {
					break
				}
			}
		}
	}
}
