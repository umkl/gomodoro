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
		SetVersion("1.0.0").
		SetDescription("Pomodoro timer for your terminal")

	commando.
		Register("time").
		AddArgument("time", "Time in minutes", "20").
		SetDescription("Starts the timer").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("Starting timer for", args["time"].Value, " minutes")
			inputTimeParsed, _ := strconv.Atoi(args["time"].Value)
			pomodoroDuration := time.Duration(inputTimeParsed) * time.Minute
			pomodoroStartTime := time.Now()
			fmt.Print(pomodoroStartTime.Second(), pomodoroDuration.Minutes())
			// for time.Now() != pomodoroStartTime.Add(pomodoroDuration*time.Minute) {
			// 	time.Sleep(time.Second) // Sleep for 1 second
			// 	secondsSinceStart := time.Now().Sub(pomodoroStartTime).Round(time.Second)
			// 	fmt.Printf("\r %v", secondsSinceStart)
			// }
		})
	commando.Parse(nil)
}

// func main() {
// 	// Parse command-line arguments
// 	var opts options
// 	args, err := flags.ParseArgs(&opts, os.Args[1:])
// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	// Convert to internal config
// 	cfg := config.New()
// 	cfg.Verbose = opts.Verbose
//   	// more taking command line options and putting them into a config struct.

// 	if opts.Pass {
// 		// ask the user for a password
// 	}

// 	// Run the App
// 	err = app.Run(cfg)
// 	if err != nil {
// 		// do stuff
// 		os.Exit(1)
// 	}
// }
