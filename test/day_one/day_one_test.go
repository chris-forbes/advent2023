package day_one

import (
	"github.com/chris-forbes/advent-23/pkg/day_one"
	"testing"
)

func TestDayOne(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"Test for 12", "1abc2", "12"},
		{"Test for 38", "pqr3stu8vwx", "38"},
		{"Test for 15", "a1b2c3d4e5f", "15"},
		{"Test for 77", "treb7uchet", "77"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := day_one.FirstAndLast(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestSummation(t *testing.T) {
	values := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	ans := day_one.FindAndSum(&values)
	if ans != 142 {
		t.Errorf("Expected: %d, got :%d", 142, ans)
	}
}
