package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("gomo").
		SetVersion("1.0.0").
		SetDescription("Pomodoro timer for your terminal")

	commando.
		Register(nil).
		AddArgument("time", "Time in minutes", "20").
		SetDescription("Starts the timer").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("Starting timer for", args["time"])
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
