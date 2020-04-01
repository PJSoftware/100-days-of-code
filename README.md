# TweetCommit

This was created as part of my (first) [#100DaysOfCode](https://github.com/PJSoftware/100-days-of-code) challenge.

Every day, after finishing whatever coding I was doing, I would update the `log.md` file there to document the day's work, then commit it, push it to `github`, then jump across to `twitter` and tweet something about what I'd done (including, as always, a link back to the log file.)

Which is great 'n all, but I'm lazy. Why perform all those extra steps when I could get the computer to do them for me? I already had `increment-log.py` adding the new empty entry stub for me; I just needed some code to:

* Read the `log.md` file and extract relevant details from the latest entry.
* Construct a git commit message, run `git commit` and `git push`.
* Construct a tweet and post it to twitter.

Easy, *n'est-ce pas*?

Well actually ... yeah. I found this info at [TutorialEdge](https://tutorialedge.net/golang/writing-a-twitter-bot-golang/) which helped me get the twitter part working, albeit using a remote library which, if the `go.sum` file is to be believed, pulled in numerous other dependencies! I would like, at some point, to tackle working with the Twitter API directly, but for the purposes of this project I just wanted to get something working straight away!

And the rest of the code was pretty straight-forward. Read and parse a text file, and run a few external `git` commands. Job done.

## Update

While the code works, I have been experiencing an odd delay when I run it. I think it would be a good idea to add output to give a better idea of where the delay is occurring.

At the same time, I'll look at adding an error type to better handle internal errors. And maybe even some unit tests...
