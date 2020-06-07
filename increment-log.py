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
    print("Retrieving latest code from remote origin")
    subprocess.run(['git', 'checkout', 'master'])
    subprocess.run(['git', 'pull'])

    gs = subprocess.run(['git', 'diff-index', 'master',
                         'log.md'], stdout=subprocess.PIPE)
    if gs.stdout != b'':
        print("log.md has uncommitted changes; ignoring")
        return

    rf = open("log.md", "r")
    wf = open("log.md.out", "w")
    matched = False

    day = 0
    for line in rf:
        if not matched:
            m = re.match(r"^#+ Day (\d+): (\S+) (\d+), (\d+)", line)
            if m:
                day = int(m.group(GI.Day.value))
                insertDay(wf, m)
                matched = True
        wf.write(line)

    rf.close()
    wf.close()
    os.remove("log.md")
    os.rename("log.md.out", "log.md")

    if day > 0:
        subprocess.run(['git', 'checkout', '-b', 'day-'+str(day+1), 'master'])


def insertDay(wf, m):
    day = int(m.group(GI.Day.value))
    ds = m.group(GI.MDay.value)+" "+m.group(GI.Month.value) + \
        " "+m.group(GI.Year.value)
    logdate = datetime.datetime.strptime(ds, "%d %B %Y")
    logdate = logdate + datetime.timedelta(days=1)
    if logdate > date.today():
        print("Increment cancelled; new date would be in the future!")
        return

    if logdate < date.today():
        print("Next date should be "+logdate.strftime("%d, %Y: %A")+"; updating to today!")
        logdate = date.today()

    h2 = "## Day " + str(day+1) + ": "
    datestr1 = logdate.strftime("%B")
    datestr2 = logdate.strftime("%d, %Y: %A")
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
