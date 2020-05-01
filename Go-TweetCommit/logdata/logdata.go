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
}

// NewLogData reads specified log file
func NewLogData(logfn string) *LogData {
	fmt.Printf("Parsing contents of '%s'\n", logfn)
	ld := new(LogData)
	err := ld.parse(logfn)
	if err != nil {
		log.Fatalf("Error reading '%s': %v", logfn, err)
	}
	return ld
}

//ParseRE used to use named regexp tests
type ParseRE struct {
	task    string
	pattern string
	regexp  *regexp.Regexp
}

func rePatterns() []*ParseRE {
	re := []*ParseRE{
		{"RoundInfo", `^[#] .+[(]Round (\d+)[)]`, nil},
		{"DayInfo", `^[#]{2} Day (\d+):`, nil},
		{"URLInfo", `^[#]{3} [[]([^]]+)[]][(][^)]+[)]\s*(\S.+)?$`, nil},
		{"DescInfo", `^[*]{2}Today's Progress:[*]{2} (.+)`, nil},
	}
	for _, pat := range re {
		pat.regexp = regexp.MustCompile(pat.pattern)
	}
	return re
}

func (ld *LogData) parse(logfn string) error {
	file, err := os.Open(logfn)
	if err != nil {
		return err
	}
	defer file.Close()

	patterns := rePatterns()
	scanner := bufio.NewScanner(file)

readFile:
	for scanner.Scan() {
		line := scanner.Text()
		for _, pat := range patterns {
			if pat.regexp.MatchString(line) {
				switch pat.task {
				case "RoundInfo":
					if ld.Round != 0 {
						return fmt.Errorf("Round encountered twice; error in %s format", logfn)
					}
					ld.Round, _ = strconv.Atoi(pat.regexp.FindStringSubmatch(line)[1])
				case "DayInfo":
					if ld.Day != 0 {
						return fmt.Errorf("Day encountered twice; error in %s format", logfn)
					}
					ld.Day, _ = strconv.Atoi(pat.regexp.FindStringSubmatch(line)[1])
				case "URLInfo":
					if ld.Topic != "" {
						return fmt.Errorf("Topic encountered twice; error in %s format", logfn)
					}
					tp := pat.regexp.FindStringSubmatch(line)[1]
					ts := pat.regexp.FindStringSubmatch(line)[2]
					if tp == "text" {
						return fmt.Errorf("Latest entry for day %d not filled in", ld.Day)
					}
					if ts != "" {
						tp = tp + " " + ts
					}
					ld.Topic = tp
				case "DescInfo":
					if ld.Round == 0 || ld.Day == 0 || ld.Topic == "" {
						return fmt.Errorf("Description found before other info; error in %s format", logfn)
					}
					ld.Desc = pat.regexp.FindStringSubmatch(line)[1]
					break readFile
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
