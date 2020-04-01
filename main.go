package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/PJSoftware/TweetCommit/logdata"
	"github.com/PJSoftware/TweetCommit/tweeter"
)

func main() {
	fn := "log.md"
	if gitUnchanged(fn) {
		fmt.Println(fn + " has not been modified; cancelling!")
		return
	}

	ld := logdata.NewLogData(fn)
	gitCommit(ld)

	err := tweeter.TweetHDC(ld)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully tweeted!")
	}
}

func gitUnchanged(fn string) bool {
	fmt.Printf("Checking diff-index of '%s'\n", fn)
	out, err := exec.Command("git", "diff-index", "HEAD", fn).Output()
	if err != nil {
		log.Fatal(err)
	}
	str := string(out)

	if str == "" {
		fmt.Println("File has changes which need to be committed!")
		return true
	}
	return false
}

func gitCommit(ld *logdata.LogData) {
	msg := fmt.Sprintf("Day %d: %s, %s", ld.Day, ld.Topic, ld.Desc)
	fmt.Printf("Committing log.md with following message:\n  * '%s'\n", msg)

	type cmdObj struct {
		cmdS string
		cmdX *exec.Cmd
	}
	var cmds []cmdObj
	cmds = append(cmds, cmdObj{"add", exec.Command("git", "add", "log.md")})
	cmds = append(cmds, cmdObj{"commit", exec.Command("git", "commit", "log.md", "-m", msg)})
	cmds = append(cmds, cmdObj{"push", exec.Command("git", "push")})
	for _, cmd := range cmds {
		fmt.Printf("  * Running command: git %v ...\n", cmd.cmdS)
		cmd.cmdX.Run()
	}
}
