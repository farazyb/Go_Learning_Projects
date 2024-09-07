package main

import (
	"testing"
)

type testpair struct {
	roman    string
	expected int
}

var tests = []testpair{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomansToInteger(t *testing.T) {
	for _, pair := range tests {
		if output := romanToInt(pair.roman); output != pair.expected {
			t.Errorf("Output %q not equal to expected %#q", output, pair.expected)
		}
	}
}
