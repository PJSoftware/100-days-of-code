package logdata

import "testing"

func TestREPatterns(t *testing.T) {
	re := rePatterns()
	for _, pat := range re {
		var str string
		var exp []string
		switch pat.task {
		case "RoundInfo":
			str = "# Pete's Log: 100 Days Of Code Challenge (Round 1)"
			exp = []string{"1"}
		case "DayInfo":
			str = "## Day 21: April 9, 2020: Thursday"
			exp = []string{"21"}
		case "URLInfo":
			str = "### [pdfill package](https://github.com/PJSoftware/pdfill) Split() Functionality Complete"
			exp = []string{"pdfill package", "Split() Functionality Complete"}
		case "DescInfo":
			str = "**Today's Progress:** Split() is now fully working"
			exp = []string{"Split() is now fully working"}
		default:
			t.Errorf("Unexpected pattern name '%s'", pat.task)
		}
		matches := pat.regexp.FindAllStringSubmatch(str, -1)
		if len(matches) == 0 {
			t.Errorf("%s: No match found", pat.task)
		} else if len(matches[0]) < 2 {
			t.Errorf("%s: Match seems malformed", pat.task)
		} else {
			got := matches[0][1]
			if got != exp[0] {
				t.Errorf("%s: Expected '%s', got '%s'", pat.task, exp[0], got)
			}
			if len(exp) > 1 {
				got := matches[0][2]
				if got != exp[1] {
					t.Errorf("%s: Also expected '%s', got '%s'", pat.task, exp[1], got)
				}
			}
		}

	}
}
