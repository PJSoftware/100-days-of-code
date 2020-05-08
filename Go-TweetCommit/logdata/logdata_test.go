package logdata

import (
	"testing"
)

func initTestValues() []string {
	inputList := []string{
		"# Pete's Log: 100 Days Of Code Challenge (Round 1)",
		"## Day 21: April 9, 2020: Thursday",
		"### New [pdfill package](https://github.com/PJSoftware/pdfill) Split() Complete",
		"**Today's Progress:** Split() is now fully working",
	}
	return inputList
}

func TestExtractData(t *testing.T) {
	ld := NewLogData("") // No file specified so parsing not triggered
	inputList := initTestValues()
	expected := struct {
		Round int
		Day   int
		Topic string
		Desc  string
	}{
		Round: 1,
		Day:   21,
		Topic: "New pdfill package Split() Complete",
		Desc:  "Split() is now fully working",
	}

	for _, line := range inputList {
		err := ld.extractData(line)
		if err != nil {
			t.Errorf("Got unexpected error: %s", err)
		}
	}

	checkResult(t, "round", ld.Round, expected.Round)
	checkResult(t, "day", ld.Day, expected.Day)
	checkResult(t, "topic", ld.Topic, expected.Topic)
	checkResult(t, "desc", ld.Desc, expected.Desc)
}

func checkResult(t *testing.T, desc string, got interface{}, exp interface{}) {
	switch got.(type) {
	case int:
		g := got.(int)
		e := exp.(int)
		if g != e {
			t.Errorf("For '%s', expected %d, got %d", desc, e, g)
		}
	case string:
		g := got.(string)
		e := exp.(string)
		if g != e {
			t.Errorf("For '%s':\n  want '%s'\n  got  '%s'", desc, e, g)
		}
	default:
		t.Errorf("Unexpected result type for '%s'", desc)
	}
}

func TestExtractDataErrors(t *testing.T) {
	ld := NewLogData("") // No file specified so parsing not triggered
	inputList := initTestValues()

	// Desc detected first
	err := ld.extractData(inputList[3])
	if err == nil {
		t.Errorf("Desc found first: Expected an error")
	}

	// Multiples found
	dataName := []string{"Round", "Day", "Topic"}
	for i := 0; i <= 2; i++ {
		_ = ld.extractData(inputList[i])
		err = ld.extractData(inputList[i])
		if err == nil {
			t.Errorf("%s found twice: expected an error", dataName[i])
		}
	}
}
