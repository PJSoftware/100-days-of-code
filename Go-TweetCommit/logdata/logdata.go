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
	ld := new(LogData)
	ld.re = rePatterns()
	ld.logFN = logfn
	if logfn != "" {
		fmt.Printf("Parsing contents of '%s'\n", logfn)
		err := ld.parse()
		if err != nil {
			log.Fatalf("Error reading '%s': %v", logfn, err)
		}
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

	for scanner.Scan() {
		line := scanner.Text()
		err = ld.extractData(line)
		if err != nil {
			return err
		}
		if ld.Desc != "" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (ld *LogData) extractData(line string) error {
	if ld.re.Round.MatchString(line) {
		round, err := ld.extractRound(line)
		if err != nil {
			return err
		}
		ld.Round = round
	} else if ld.re.Day.MatchString(line) {
		day, err := ld.extractDay(line)
		if err != nil {
			return err
		}
		ld.Day = day
	} else if ld.re.Title.MatchString(line) {
		title, err := ld.extractTitle(line)
		if err != nil {
			return err
		}
		ld.Topic = title
	} else if ld.re.Desc.MatchString(line) {
		desc, err := ld.extractDesc(line)
		if err != nil {
			return err
		}
		ld.Desc = desc
	}
	return nil
}

func (ld *LogData) extractRound(line string) (int, error) {
	if ld.Round != 0 {
		return 0, fmt.Errorf("Round encountered twice; error in file format")
	}
	return strconv.Atoi(ld.re.Round.FindStringSubmatch(line)[1])
}

func (ld *LogData) extractDay(line string) (int, error) {
	if ld.Day != 0 {
		return 0, fmt.Errorf("Day encountered twice; error in file format")
	}
	return strconv.Atoi(ld.re.Day.FindStringSubmatch(line)[1])
}

func (ld *LogData) extractTitle(line string) (string, error) {
	if ld.Topic != "" {
		return "", fmt.Errorf("Topic encountered twice; error in file format")
	}
	topic := ld.re.Title.FindStringSubmatch(line)[1]
	if ld.re.Split.MatchString(topic) {
		topic = ld.re.Split.ReplaceAllString(topic, "$1")
	}
	return topic, nil
}

func (ld *LogData) extractDesc(line string) (string, error) {
	if ld.Round == 0 || ld.Day == 0 || ld.Topic == "" {
		return "", fmt.Errorf("Description found before other info; error in file format")
	}
	return ld.re.Desc.FindStringSubmatch(line)[1], nil
}
