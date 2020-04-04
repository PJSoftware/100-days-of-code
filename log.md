# Pete's Log: 100 Days Of Code Challenge (Round 1)

## Day 16: April 4, 2020: Saturday

### [Win-Spotlight](https://github.com/PJSoftware/Win-Spotlight) Converted to Module

**Today's Progress:** Still trying to wrap my head around the benefits of `Go` modules.

**Thoughts:** It does seem to make my pathing work a little better--by forcing me to use the github paths to everything. Since I'm also trying to wrap my head around how best to use github to host `Go` packages, this seems to be a step in the right direction.

At the same time, I was wrestling with the fact that `go build` was, by default, creating a `Win-Spotlight.exe` file rather than the preferred `UpdateSpotlight.exe`. I could probably have renamed the repository to fix this, but I opted, instead, to use a `cmd/UpdateSpotlight` folder. This means I now have to `cd` into that folder to build my code, but maybe I'll add a `build.sh` to the parent folder to get around that.

## Day 15: April 3, 2020: Friday

### [Log Incrementer increment-log.py](https://github.com/PJSoftware/100-days-of-code/blob/master/increment-log.py) Update 2

**Today's Progress:** Fixed the formatting of the date

**Thoughts:** While making my last tweak of `increment-log.py`, I spent a little time looking for a formatting code to use with `strftime` which would *not* zero-pad the day of the month. I found something online which told me I could use `%-d` instead of `%d`. Turns out this was not correct--and because that particular line of code is only interpreted at runtime, I would not find out until I tried to run the script today.

To get the format I actually wanted, I had to specifically check for the leading zero and remove it.

To ensure this does not happen again--well, I guess it's time to look into unit testing in Python!

## Day 14: April 2, 2020: Thursday

### [TweetCommit](https://github.com/PJSoftware/TweetCommit) Output Revised

**Today's Progress:** Revised output to better explain what is happening.

**Thoughts:** `TweetCommit.exe` is working okay, but it needs a little refinement. It actually *seems* to hang for a while, so I've added output to give a better idea of what is going on. (I *think* I know where the delay is happening, but until I run it with the new output enabled I really don't know for sure!)

**Update:** `git push` is where the biggest delay lies...

### `increment-log.py` Revised

**Thoughts:** I thought there was something in my code which detected an attempt to add a date in the future to the `log.md` file, but apparently there was not.

Now there is:

```python
    if date > date.today():
        print("Increment cancelled; new date would be in the future!")
        return
```

## Day 13: April 1, 2020: Wednesday

### [Win-Spotlight](https://github.com/PJSoftware/Win-Spotlight) Revisited

**Lost Time:** Okay, so I missed a few days between my last coding attempt and now. In that time I watched plenty of NetFlix and did zero coding. (You guys: `#TheOA` is amazing! Check it out!)

**Today's Progress:** Made a quick fix to file naming; added tests and error handling.

#### Wallpaper Naming

`Win-Spotlight` is one of the first things I wrote in Go. It looks for new `Spotlight` images (delivered on the regular by Windows 10) and copies them out of their delivery folder into a specified wallpaper folder. This gave me the opportunity to play with `INI` and `JSON` files in `Go` -- not to mention exploring working with `UTF16`. Plus it gives me a folder full of awesome wallpapers to use on my machine.

It also confirmed, once and for all, that I had to move away from using the old text editor I've been using for the last 18-odd years -- and, honestly, much as I loved it, that has been the best thing I could have done!

I noticed that, despite my efforts to intercept it in the renaming process, some of the files in my wallpaper folder had a doubled-up copyright symbol in their name. I renamed them there, but obviously the only way to truly fix the problem is by curing the illness rather than treating the symptoms. *(Signs that you are living through the **Coronapocalypse**: medical analogies in your log files!)* So my first work for today was to fix up that particular piece of code.

#### Unit Tests

**The Plan:**

So, did my fix actually do what I think it does? How would I know?

As somebody who is mostly self-taught, and who has done most of their coding in isolation\*, I only picked up the concept of `TDD` (Test Driven Development) a couple of years ago. Great idea, loved it, and promptly started applying it (no doubt naively) to my `Perl` programming. It's probably time to work on doing the same in my ongoing `Go` work!

I probably won't write a completely comprehensive test suite today, but I'd at least like to get enough done that I can confirm my new code works as intended.

\* This has shaped much of my approach to coding (obviously) and it will probably come up again!

**Achieved:** I added one unit test for `newFilename()` to ensure that our output is as expected. This enabled me to rewrite the function to be more efficient, and to put all our intended regexp transformations in a table. (I originally coded it as a `map`, but the order of transforms was important to the end result, so I redefined it as a `slice` of `struct`s.)

#### Error Handling

**The Plan:**

As somebody who is mostly self-taught, and who has done most of their coding in isolation\*, my approach to error handling in other languages has been ... minimal. A few years back I ran across the concept of throwing exceptions and letting the calling code catch them if necessary. Plenty of this started creeping into my `Perl` code from that point on.

Of course, `Go` takes error handling in the complete opposite direction, treating errors as values which can be passed back to the calling code, there to be checked or ignored as required -- with the caveat that ignoring them *must* be a deliberate choice rather than a simple oversight. In many ways this is the same logic as throwing an exception: the programmer can choose to either deal with it or ignore it, but they *must* choose. I like it. Even if it does tend to clutter the code...

During my *lost-time* hiatus (I'm just gonna go ahead and blame alien abduction!) I actually did a little reading on error handling in Go. It's something I'll be exploring further, but certainly adding an `Error` type in this code (if I have the time) would be a good start.

\* Told you!

**Achieved:** I fiddled around a bit with an `errors` sub-package, and applied it (very simply) in a few places. At the moment I'm simply lifting it from various online locations; I need to do more thinking about how best to implement it for this application (and/or any other.)

## Day 12: March 26, 2020: Thursday

### [TweetCommit (was AutoTweeter)](https://github.com/PJSoftware/TweetCommit) Final Tweak

**Today's Progress:** Renamed. Added code to prevent tweeting the same entry more than once.

**Thoughts:** So ... `increment-log.py` has safeguards to prevent it from incrementing the log when there are uncommitted changes, or when it only contains the stub from the previous run. And `TweetCommit.exe` has safeguards to prevent it from doing anything when the log has not yet been filled in. But as far as I can tell, if I happen to run it twice in a row, there is nothing to stop it from tweeting out the same tweet it has already sent. Which is, I guess, not actually *wrong*, but it *is* a little spammy.

I guess the best way to achieve that would be to check whether git considers the `log.md` file clean or not. We have already done this in `Python` with:

```python
    gs = subprocess.run(['git', 'diff-index', 'master', 'log.md'], stdout=subprocess.PIPE)
    if gs.stdout != b'':
        print("log.md has uncommitted changes; ignoring")
        return
```

The question is, how do we do this in `Go`? Presumably the same way: run `git diff-index` and interpret the result. Time to do a little more reading around the `exec` library, I guess.

**Solution:**

```go
    out, err := exec.Command("git", "diff-index", "HEAD", fn).Output()
    if err != nil {
        log.Fatal(err)
    }
    str := string(out)

    if str == "" { // no changes
        return true
    }

```

## Day 11: March 25, 2020: Wednesday

### [TweetCommit](https://github.com/PJSoftware/TweetCommit) Day 2

**Today's Progress:** TweetCommit completed; now commits AND tweets

**Thoughts:** Of course, until I run my script I won't actually know if it works as intended. That's the joy of code which tweets and/or commits: you don't want to be test-running it a thousand times on potentially live systems.

**Update:** It works!

## Day 10: March 24, 2020: Tuesday

### [TweetCommit](https://github.com/PJSoftware/TweetCommit) Day 1

**Today's Progress:** Preliminary work on TweetCommit

**Thoughts:** This code is adapted from [TutorialEdge.net](https://tutorialedge.net/golang/writing-a-twitter-bot-golang/) and uses a couple of external libraries. Ultimately I think I'd like to tackle writing something from scratch, but for now I just wanted something that worked.

The **TutorialEdge** code worked by setting your API user tokens as part of your environment--because as author **Elliot Forbes** rightly says, information like that should not be committed to your repository. Since I work on two different home machines at the moment--a desktop and a laptop--I opted instead to save my API tokens in a json file on my home network so that both computers can access them easily.

Tweeting via this code does actually work, but that is only part of what I want it to do. The second step would be for it to read and parse *this* log file, pull out the most recent Day number, project name, and "Today's Progress" blurb, and construct a tweet from those.  However, it should *also* commit (and push) this file to github using the same information. Possibly "AutoTweet" is a bit of a misnomer, and the repo name may change at some point. (Done.)

**Things to Investigate Further:** Go Modules. Also, how best to publish individual Go packages to github so I can use them in multiple projects if required.

## Day 9: March 23, 2020: Monday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 9

**Today's Progress:** Added '-all' option

**Thoughts:** The '-all' option directs the solver to attempt to solve all puzzles contained in the 'puzzles' folder. Summaries of the results will be displayed. This provides an easier way to test the latest solvers on all available puzzles.

Having looked at the final state of the two puzzles which cannot currently be solved, it seems that the planned fourth solver will not be able to make much progress on one of them ("expert.sp") and therefore, perhaps I need to revise my intended approach somewhat.

I may need to start actually researching advanced solving techniques.

## Day 8: March 22, 2020: Sunday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 8

**Today's Progress:** Added third solver

**Thoughts:** The code for this one is a little ... messy. I feel it could probably be refactored quite a bit, and that will probably happen when I add the next solver, because I feel they'll both be using similar code. Adding this solver allowed us to solve our "hard" puzzle.

## Day 7: March 21, 2020: Saturday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 7

**Today's Progress:** Added another solver.

**Thoughts:** The second solving method is to examine each block and, for each value that is not already used, examine all empty squares to determine if any of them only have one possible location. To simplify the code calling each solution in turn, I built a function table; all future solutions simply need to be added to the table in order to be processed.

The new code solves at least one of the puzzles in my meagre repertoire which was previously unsolvable, so I guess that's progress.

Next step: work on solving the remaining cells of the unsolvable puzzles, determine my next strategy, and codify it.

## Day 6: March 20, 2020: Friday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 6

**Today's Progress:** Added ShowWorking and OPV support

**Thoughts:** I actually did not do any of this coding until Saturday morning, but I did a lot of thinking about it on Friday, so it counts. (I'll be doing more coding on Saturday too!)

First, I refined the display of the grid to improve readability.

Second, I added the ability for my solvers to optionally show their working. At the same time I clarified the "OnePossibleValue" approach which my only actual solver func uses. This became necessary because I had added a couple more "easy"/"medium" puzzles (from a magazine Mum had, while I was at her place on Thursday night) which I *expected* my solver to handle, but which it could not touch. It did not make sense, because from a quick visual inspection I could easily fill in a couple of values. Something was going wrong, and I needed to figure out what.

Turns out what was going on was that my code was working perfectly as intended. It was my brain that was glitching; the method I was using to get started on the new puzzles was not simply "OnePossibleValue"; I was actually visually identifying narrow points in individual blocks and then eliminating possible values for surrounding cells. It was just so automatic that I thought I was doing OPV processing.

Did I just debug my brain?

Now I need to analyse what I was actually doing by eye, and convert that back into code. But that's a task for tomorrow\*.

*\* Later today!*

## Day 5: March 19, 2020: Thursday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 5

**Today's Progress:** Rewrote the code. It now works -- although it cannot solve more difficult puzzles.

**Thoughts:** While pulling the code apart so I could clean up some of the unnecessary complications, I noticed a loop around a block of code which did not actually make sense. That may not have been the cause of the bug, but it certainly wasn't helping. The new version is cleaner, and it works.

The next step will be to figure out how to solve more complex puzzles.

## Day 4: March 18, 2020: Wednesday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 4

**Today's Progress:** Zero!

**Thoughts:** I've spent some time adding print statements and attempting to decipher what I'm seeing. It seems as though, for each set of cellCollection objects, only the final cell is being updated. I have arrays of arrays of pointers to arrays of pointers -- and I think somewhere in there, my understanding of how things work does not match with reality.

More, smarter debugging and/or testing is required.

### Log Incrementer [increment-log.py](https://github.com/PJSoftware/100-days-of-code/blob/master/increment-log.py) Update 1

**Today's Progress:** Added detection of git status of log.md; will not add a new day's entry until previous changes have been committed.

**Thoughts:** Now that this is in place, we no longer need the somewhat clumsy method of making our new "##" an "###" to allow detection that way. Both approaches achieve the same end result: preventing us from inadvertently running the script multiple times between edits. Also, added an enum.

## Day 3: March 17, 2020: Tuesday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 3

**Today's Progress:** The First Pass solver technically should be working, but I appear to have a bug in the code which updates the cell possible values.

**Thoughts:** I'm not sure quite what the problem is, but some of my "possible cell value" data is not being set properly. I need to dig into it a little deeper. Looks like Day 4 will be primarily a debugging session.

## Day 2: March 16, 2020: Monday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 2

**Today's Progress:** Revised code structure. Import() and Display() now working as intended.

**Thoughts:** As well as the overhaul of the Sudoku Grid internals, I ran into a little confusion over whether my grid was numbered 1-9 x 1-9, or 0-8 x 0-8. Internally, of course, the array index is zero-based, but in my head I'd been thinking one-based, and converting where required. I think this is because the valid values are 1-9. At some point during today's coding I started using zero-based loops, and when the two systems clashed, I ultimately set up the gridCoord array, so I could use "for xi, xn := range gridCoord" and use either xi or xn as required. Ultimately, this revealed to me that I really only require zero-based numbering of rows and columns!

**Note**: I decided that keeping this log in reverse order (most recent day first) actually makes it easier to read. YMMV (and I may yet change my mind; it's known to happen!) I'm also thinking that very soon now I shall write a script or somthing to automate inserting the new day's date and template...

### Log Incrementer [increment-log.py](https://github.com/PJSoftware/100-days-of-code/blob/master/increment-log.py)

**Today's Progress:** Added **increment-log.py** to this repository.

**Thoughts:** I wanted a script to automate adding a new day to our **log.md** file. It would have been a matter of moments to write this in Perl--so I opted to write it in Python instead. I am still at the "google every line of code" stage with Python, so it took a little longer. Additionally, much as I wanted to use this script straight away, it was really done during Day 2, so rather than throw myself out of sync, I've relocated it here so Day 3 can proceed as intended...

## Day 1: March 15, 2020: Sunday

### [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver) Day 1

**Today's Progress:** I worked out the basic code structure and added the basic ability to import a puzzle from an external file. The code appears to compile and run at this point, although I really need to add some tests.

**Thoughts:** I feel like this was a good start.
