package main

import (
	"fmt"

	"github.com/PJSoftware/TweetCommit/git"
	"github.com/PJSoftware/TweetCommit/logdata"
	"github.com/PJSoftware/TweetCommit/tweeter"
)

func main() {
	fmt.Printf("TweetCommit v%s\n", version)
	fmt.Printf("Copyright © 2020 by PJSoftware\n\n")

	fn := "log.md"
	if git.Unchanged(fn) {
		fmt.Println(fn + " has not been modified; cancelling!")
		return
	}

	ld := logdata.NewLogData(fn)
	tw := tweeter.NewTweeter()
	if tw.Err != nil {
		fmt.Printf("Error initialising Twitter: %v\n", tw.Err)
		return
	}

	err := tw.TweetHDC(ld)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully tweeted!")
	git.Commit(ld)
}
