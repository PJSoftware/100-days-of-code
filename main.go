package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/PJSoftware/AutoTweeter/logdata"
)

func main() {
	fn := "log.md"
	if gitUnchanged(fn) {
		fmt.Println(fn + " has not been modified; cancelling!")
		return
	}

	// ld := logdata.NewLogData(fn)
	// wannaProceed(ld)

	// gitCommit(ld)
	// err := tweeter.TweetHDC(ld)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Successfully tweeted!")
	// }
}

func gitUnchanged(fn string) bool {
	fmt.Printf("Checking '%s'\n", fn)
	out, err := exec.Command("git", "diff-index", "HEAD", fn).Output()
	if err != nil {
		log.Fatal(err)
	}
	str := string(out)

	if str == "" {
		return true
	}
	return false
}

func wannaProceed(ld *logdata.LogData) {

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
