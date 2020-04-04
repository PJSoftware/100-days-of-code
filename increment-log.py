#!/usr/bin/env python

import datetime
import re
import os
import subprocess

from enum import Enum

# GI => GroupIndex
class GI(Enum):
    Day = 1
    Month = 2
    MDay = 3
    Year = 4
    
def incLog():
    gs = subprocess.run(['git', 'diff-index', 'master', 'log.md'], stdout=subprocess.PIPE)
    if gs.stdout != b'':
        print("log.md has uncommitted changes; ignoring")
        return

    rf = open("log.md", "r")
    wf = open("log.md.out", "w")
    matched = False

    for line in rf:
        if not matched:
            m = re.match(r"^#+ Day (\d+): (\S+) (\d+), (\d+)", line)
            if m:
                insertDay(wf, m)
                matched = True
        wf.write(line)

    rf.close()
    wf.close()
    os.remove("log.md")
    os.rename("log.md.out", "log.md")

def insertDay(wf, m):
    day = int(m.group(GI.Day.value))
    ds = m.group(GI.MDay.value)+" "+m.group(GI.Month.value)+" "+m.group(GI.Year.value)
    date = datetime.datetime.strptime(ds, "%d %B %Y")
    date = date + datetime.timedelta(days=1)
    if date > date.today():
        print("Increment cancelled; new date would be in the future!")
        return

    h2 = "## Day " + str(day+1) + ": "
    datestr1 = date.strftime("%B")
    datestr2 = date.strftime("%d, %Y: %A")
    if datestr2[:1] == "0":
        datestr2 = datestr2[1:]
    
    datestr = datestr1 + " " + datestr2
    h2 = h2 + datestr

    print("Adding Day", day+1, "stub [" + datestr + "]")
    wf.write(h2 + "\n\n")
    wf.write("### [text](url) Day n\n\n")
    wf.write("**Today's Progress:** \n\n")
    wf.write("**Thoughts:** \n\n")

incLog()