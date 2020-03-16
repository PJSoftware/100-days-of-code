#!/usr/bin/env python

import datetime
import re
import os

def incLog():
    rf = open("log.md", "r")
    wf = open("log.md.out", "w")
    added = False

    for line in rf:
        if not added:
            m = re.match(r"^(#+) Day (\d+): (\S+) (\d+), (\d+)", line)
            if m:
                insertDay(wf, m)
                added = True
        wf.write(line)

    rf.close()
    wf.close()
    os.remove("log.md")
    os.rename("log.md.out", "log.md")

def insertDay(wf, m):
    if m.group(1) == "###":
        return

    day = int(m.group(2))
    ds = m.group(4)+" "+m.group(3)+" "+m.group(5)
    date = datetime.datetime.strptime(ds, "%d %B %Y")
    date = date + datetime.timedelta(days=1)

    h2 = "### Day " + str(day+1) + ": "
    h2 = h2 + date.strftime("%B %d, %Y: %A")

    wf.write(h2 + "\n\n")
    wf.write("**Today's Progress:** \n\n")
    wf.write("**Thoughts:** \n\n")
    wf.write("**Link to work:** [text](url)\n\n")

incLog()