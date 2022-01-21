package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	startTime := time.Now()
	pomo := time.Duration(time.Minute * 25)
	for i := 0; i < 1500; i++ {
		time.Sleep(time.Second)
		currentTime := time.Now()
		subtr := currentTime.Sub(startTime).Round(time.Second)
		// fmt.Println(subtr)
		cTime := pomo - subtr
		fmt.Printf("\r %v", cTime)
	}

}
