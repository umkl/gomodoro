package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("gomo").
		SetVersion("0.0.0").
		SetDescription("Pomodoro timer for your terminal")

	commando.
		Register("time").
		AddArgument("time", "Time in minutes", "25").
		SetDescription("Starts the timer").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("Starting timer for", args["time"].Value, "minutes")
			inputTimeParsed, _ := strconv.Atoi(args["time"].Value)
			pomodoroDuration := time.Duration(inputTimeParsed) * time.Minute
			pomodoroStartTime := time.Now()
			pomodoroEndTime := pomodoroStartTime.Add(pomodoroDuration)

			for !time.Now().After(pomodoroEndTime) {
				time.Sleep(time.Second)
				secondsSinceStart := time.Now().Sub(pomodoroStartTime).Round(time.Second)
				fmt.Printf("\r %v", secondsSinceStart)
			}
			fmt.Println("\nPomodoro finished")
			f, err := os.Open("assets/alarm.wav")
			if err != nil {
				log.Fatal(err)
			}
			streamer, format, err := wav.Decode(f)
			if err != nil {
				log.Fatal(err)
			}
			defer streamer.Close()
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
			select {} //hangs the program forever
		})
	commando.Parse(nil)
}
