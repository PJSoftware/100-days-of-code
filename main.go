package main

import (
	"fmt"

	"github.com/PJSoftware/AutoTweeter/tweeter"
)

func main() {
	err := tweeter.TweetHDC(1, 10, "This is a message")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully tweeted!")
	}
}
