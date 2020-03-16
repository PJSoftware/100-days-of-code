# Pete's Log: 100 Days Of Code Challenge (Round 1)

## Day 3: March 17, 2020: Tuesday

**Today's Progress:** Added **increment-log.py** to this repository.

**Thoughts:** I wanted a script to automate adding a new day to our **log.md** file. It would have been a matter of moments to write this in Perl--so I opted to write it in Python instead. I am still at the "google every line of code" stage with Python, so it took a little longer.

**Link to work:** [increment-log.py](https://github.com/PJSoftware/100-days-of-code/blob/master/increment-log.py)

## Day 2: March 16, 2020: Monday

**Today's Progress:** Revised code structure. Import() and Display() now working as intended.

**Thoughts:** As well as the overhaul of the Sudoku Grid internals, I ran into a little confusion over whether my grid was numbered 1-9 x 1-9, or 0-8 x 0-8. Internally, of course, the array index is zero-based, but in my head I'd been thinking one-based, and converting where required. I think this is because the valid values are 1-9. At some point during today's coding I started using zero-based loops, and when the two systems clashed, I ultimately set up the gridCoord array, so I could use "for xi, xn := range gridCoord" and use either xi or xn as required. Ultimately, this revealed to me that I really only require zero-based numbering of rows and columns!

**Link to work:** [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver)

**Note**: I decided that keeping this log in reverse order (most recent day first) actually makes it easier to read. YMMV (and I may yet change my mind; it's known to happen!) I'm also thinking that very soon now I shall write a script or somthing to automate inserting the new day's date and template...

## Day 1: March 15, 2020: Sunday

**Today's Progress:** I worked out the basic code structure and added the basic ability to import a puzzle from an external file. The code appears to compile and run at this point, although I really need to add some tests.

**Thoughts:** I feel like this was a good start.

**Link to work:** [Sudoku Solver](https://github.com/PJSoftware/sudoku-solver)
