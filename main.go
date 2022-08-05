package main

import (
	"fmt"
"github.com/thatisuday/commando"
)

func main(){
commando.
SetExecutableName("gomodoro").
SetVersion("1.0.0").
SetDescription("Pomodoro timer for your terminal")

commando.	
	Register("pomo").
	AddArgument("tm", "set a custom time how long ", "25")
	
	
	commando.Parse(nil)


}


// func main() {
// 	fmt.Println(time.Now())
// 	startTime := time.Now()
// 	pomo := time.Duration(time.Minute * 25)
// 	for i := 0; i < 1500; i++ {
// 		time.Sleep(time.Second)
// 		currentTime := time.Now()
// 		subtr := currentTime.Sub(startTime).Round(time.Second)
// 		// fmt.Println(subtr)
// 		cTime := pomo - subtr
// 		fmt.Printf("\r %v", cTime)
// 	}

// }

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