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
	ld := new(LogData)
	err := ld.parse(logfn)
	if err != nil {
		log.Fatalf("Error reading '%s': %v", logfn, err)
	}
	return ld
}

func (ld *LogData) parse(logfn string) error {
	file, err := os.Open(logfn)
	if err != nil {
		return err
	}
	defer file.Close()

	patterns := []string{
		`^[#] .+[(]Round (\d+)[)]`,
		`^[#]{2} Day (\d+):`,
		`^[#]{3} [[](.+)[]][(].+[)]\s*(\S.+)?$`,
		`^[*]{2}Today's Progress:[*]{2} (.+)`,
	}
	var rec []*regexp.Regexp
	for _, pat := range patterns {
		rec = append(rec, regexp.MustCompile(pat))
	}

	scanner := bufio.NewScanner(file)

readFile:
	for scanner.Scan() {
		line := scanner.Text()
		for i, re := range rec {
			if re.MatchString(line) {
				switch i {
				case 0:
					if ld.Round != 0 {
						return fmt.Errorf("Round encountered twice; error in %s format", logfn)
					}
					ld.Round, _ = strconv.Atoi(re.FindStringSubmatch(line)[1])
				case 1:
					if ld.Day != 0 {
						return fmt.Errorf("Day encountered twice; error in %s format", logfn)
					}
					ld.Day, _ = strconv.Atoi(re.FindStringSubmatch(line)[1])
				case 2:
					if ld.Topic != "" {
						return fmt.Errorf("Topic encountered twice; error in %s format", logfn)
					}
					tp := re.FindStringSubmatch(line)[1]
					ts := re.FindStringSubmatch(line)[2]
					if tp == "text" {
						return fmt.Errorf("Latest entry for day %d not filled in", ld.Day)
					}
					if ts != "" {
						tp = tp + " " + ts
					}
					ld.Topic = tp
				case 3:
					if ld.Round == 0 || ld.Day == 0 || ld.Topic == "" {
						return fmt.Errorf("Description found before other info; error in %s format", logfn)
					}
					ld.Desc = re.FindStringSubmatch(line)[1]
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
