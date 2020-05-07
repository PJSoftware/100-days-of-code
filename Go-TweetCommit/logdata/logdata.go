package logdata

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// LogData holds the information extracted from our log file
type LogData struct {
	Round int
	Day   int
	Topic string
	Desc  string
	re    *parseRE
	logFN string
}

type parseRE struct {
	Round *regexp.Regexp
	Day   *regexp.Regexp
	Title *regexp.Regexp
	Desc  *regexp.Regexp
	Split *regexp.Regexp
}

// NewLogData reads specified log file
func NewLogData(logfn string) *LogData {
	fmt.Printf("Parsing contents of '%s'\n", logfn)
	ld := new(LogData)
	ld.re = rePatterns()
	ld.logFN = logfn
	err := ld.parse()
	if err != nil {
		log.Fatalf("Error reading '%s': %v", logfn, err)
	}
	return ld
}

func rePatterns() *parseRE {
	re := new(parseRE)

	re.Round = regexp.MustCompile(`^[#] .+[(]Round (\d+)[)]`)
	re.Day = regexp.MustCompile(`^[#]{2} Day (\d+):`)
	re.Title = regexp.MustCompile(`^[#]{3} (.+)$`)
	re.Desc = regexp.MustCompile(`^[*]{2}Today's Progress:[*]{2} (.+)`)
	re.Split = regexp.MustCompile(`[[]([^]]+)[]][(][^)]+[)]`)
	return re
}

func (ld *LogData) parse() error {
	file, err := os.Open(ld.logFN)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

readFile:
	for scanner.Scan() {
		line := scanner.Text()

		if ld.re.Round.MatchString(line) {
			err = ld.extractRound(line)
		} else if ld.re.Day.MatchString(line) {
			err = ld.extractDay(line)
		} else if ld.re.Title.MatchString(line) {
			err = ld.extractTitle(line)
		} else if ld.re.Desc.MatchString(line) {
			err = ld.extractDesc(line)
			break readFile
		}
	}

	if err != nil {
		return err
	}

	// 			switch pat.task {
	// 			case "RoundInfo":
	// 			case "DayInfo":
	// 			case "SubTitle":
	// 			case "DescInfo":
	// 				break readFile
	// 			}
	// 		}
	// 	}
	// }

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (ld *LogData) extractRound(line string) error {
	if ld.Round != 0 {
		return fmt.Errorf("Round encountered twice; error in %s format", ld.logFN)
	}
	ld.Round, _ = strconv.Atoi(ld.re.Round.FindStringSubmatch(line)[1])
	return nil
}

func (ld *LogData) extractDay(line string) error {
	if ld.Day != 0 {
		return fmt.Errorf("Day encountered twice; error in %s format", ld.logFN)
	}
	ld.Day, _ = strconv.Atoi(ld.re.Day.FindStringSubmatch(line)[1])
	return nil
}

func (ld *LogData) extractTitle(line string) error {
	if ld.Topic != "" {
		return fmt.Errorf("Topic encountered twice; error in %s format", ld.logFN)
	}
	ld.Topic = ld.re.Title.FindStringSubmatch(line)[1]
	if ld.re.Split.MatchString(ld.Topic) {
		ld.Topic = ld.re.Split.ReplaceAllString(ld.Topic, "$1")
	}
	return nil
}

func (ld *LogData) extractDesc(line string) error {
	if ld.Round == 0 || ld.Day == 0 || ld.Topic == "" {
		return fmt.Errorf("Description found before other info; error in %s format", ld.logFN)
	}
	ld.Desc = ld.re.Desc.FindStringSubmatch(line)[1]
	return nil
}
