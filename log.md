# Pete's Log: 100 Days Of Code Challenge (Round 1)

## Day 11: March 25, 2020: Wednesday

### [AutoTweeter](https://github.com/PJSoftware/AutoTweeter) Day 2

**Today's Progress:** AutoTweeter completed; now commits AND tweets

**Thoughts:** Of course, until I run my script (renamed TweetCommit.exe in this folder) I won't actually know if it works as intended. That's the joy of code which tweets and/or commits: you don't want to be test-running it a thousand times on potentially live systems.

**Update:** It works!

## Day 10: March 24, 2020: Tuesday

### [AutoTweeter](https://github.com/PJSoftware/AutoTweeter) Day 1

**Today's Progress:** Preliminary work on AutoTweeter

**Thoughts:** This code is adapted from [TutorialEdge.net](https://tutorialedge.net/golang/writing-a-twitter-bot-golang/) and uses a couple of external libraries. Ultimately I think I'd like to tackle writing something from scratch, but for now I just wanted something that worked.

The **TutorialEdge** code worked by setting your API user tokens as part of your environment--because as author **Elliot Forbes** rightly says, information like that should not be committed to your repository. Since I work on two different home machines at the moment--a desktop and a laptop--I opted instead to save my API tokens in a json file on my home network so that both computers can access them easily.

Tweeting via this code does actually work, but that is only part of what I want it to do. The second step would be for it to read and parse *this* log file, pull out the most recent Day number, project name, and "Today's Progress" blurb, and construct a tweet from those.  However, it should *also* commit (and push) this file to github using the same information. Possibly "AutoTweeter" is a bit of a misnomer, and the repo name may change at some point.

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
