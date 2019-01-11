
package main

import (
	"fmt"
	"strconv"
	"command"
)

var remainSecond int = 30
var lastSecond int = 30

func timer(seconds int, input chan string) {
	second := 0
	for second < seconds {
		second++
		min := (seconds - second) / 60
		sec := (seconds - second) % 60
		fmt.Printf("\r%02d:%02d ", min, sec)

		if second == lastSecond {
			command.SayNoWait("あと" + strconv.Itoa(second) + "秒です。終わりそうですか？")
		} else if (second != 0) && (second%remainSecond == 0) {
			message := "あと"
			if min != 0 {
				message = message + strconv.Itoa(min) + "分"
			}
			if sec != 0 {
				message = message + strconv.Itoa(sec) + "秒"
			}
			message = message + "です。"
			command.SayNoWait(message)
		}

		result := ReadWithTimeout(input)
		if result != "" {
			str := trim(result)
			if str == "s" || str == "S" || str == "skip" || str == "Skip" {
				break
			}
			i, err := strconv.Atoi(str)
			if err == nil {
				seconds = i * 60
				second = 0
			}
		}
	}
	command.Say("終了")
}

