package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("gomo").
		SetVersion("0.0.0").
		SetDescription("Pomodoro timer for your terminal")

	commando.
		Register("time").
		AddArgument("time", "Time in minutes", "5").
		SetDescription("Starts the timer").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("Starting timer for", args["time"].Value, "minutes")
			inputTimeParsed, _ := strconv.Atoi(args["time"].Value)
			pomodoroDuration := time.Duration(inputTimeParsed) * time.Second
			pomodoroStartTime := time.Now()
			pomodoroEndTime := pomodoroStartTime.Add(pomodoroDuration)
			for !time.Now().After(pomodoroEndTime) {
				time.Sleep(time.Second)
				secondsSinceStart := time.Now().Sub(pomodoroStartTime).Round(time.Second)
				fmt.Printf("\r %v", secondsSinceStart)
			}
			fmt.Println("\nPomodoro finished")

		})
	commando.Parse(nil)
}
