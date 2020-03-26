package main

import (
	"fmt"
	"os/exec"

	"github.com/PJSoftware/AutoTweeter/logdata"
	"github.com/PJSoftware/AutoTweeter/tweeter"
)

func main() {
	ld := logdata.NewLogData("log.md")

	gitCommit(ld)
	err := tweeter.TweetHDC(ld)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully tweeted!")
	}
}

func gitCommit(ld *logdata.LogData) {
	msg := fmt.Sprintf("Day %d: %s, %s", ld.Day, ld.Topic, ld.Desc)
	fmt.Printf("Committing log.md with message '%s'", msg)

	var cmds []*exec.Cmd
	cmds = append(cmds, exec.Command("git", "add", "log.md"))
	cmds = append(cmds, exec.Command("git", "commit", "log.md", "-m", msg))
	cmds = append(cmds, exec.Command("git", "push"))
	for _, cmd := range cmds {
		cmd.Run()
	}
}
