
package main

import (
	"fmt"
	"strconv"
	"command"
)

var remainSecond int = 30
var lastSecond int = 30

func Timer(seconds int) {
	second := 0
	for second < seconds {
		second++
		min := (seconds - second) / 60
		sec := (seconds - second) % 60
		fmt.Printf("\r%02d:%02d ", min, sec)

		if (seconds - second) == lastSecond {
			command.SayNoWait("あと" + strconv.Itoa(lastSecond) + "秒です。終わりそうですか？")
		} else if ((seconds - second) != 0) && ((seconds - second)%remainSecond == 0) {
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

		result := ReadWithTimeout(500)
		if result != "" {
			str := Trim(result)
			if str == "s" || str == "S" || str == "skip" || str == "Skip" {
				break
			}
			i, err := strconv.Atoi(str)
			if err == nil {
				seconds = i * 60
				second = 0
			}
		}
		result = command.ReadWithTimeout(500)
		if result != "" {
			str := Trim(result)
			if str == "s" {
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

