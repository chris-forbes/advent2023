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

func TestSummationWords(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  int
	}{
		{"Two1Nine", []string{"two1nine"}, 29},
		{"EightTwoThree", []string{"eightwothree"}, 83},
		{"OneThree", []string{"abcone2threexyz"}, 13},
		{"xtwone3four", []string{"xtwone3four"}, 24},
		{"4nineeightseven2", []string{"4nineeightseven2"}, 42},
		{"zoneight234", []string{"zoneight234"}, 14},
		{"7pqrstsixteen", []string{"7pqrstsixteen"}, 76},
		{"12", []string{"1abc2"}, 12},
		{"38", []string{"pqr3stu8vwx"}, 38},
		{"15", []string{"a1b2c3d4e5f"}, 15},
		{"77", []string{"treb7uchet"}, 77},
		{"29", []string{"two1nine"}, 29},
		{"83", []string{"eightwothree"}, 83},
		{"13", []string{"abcone2threexyz"}, 13},
		{"24", []string{"xtwone3four"}, 24},
		{"97", []string{"4nineeightseven2"}, 42},
		{"14", []string{"zoneight234"}, 14},
		{"76", []string{"7pqrstsixteen"}, 76},
		{"57", []string{"9fiveajhseven"}, 97},
		{"13", []string{"xjhoneighthree"}, 13},
		{"85", []string{"eight96nxcjjddmseightxvgsixfiverrzpvmgnl"}, 85},
		{"88", []string{"8pjkm"}, 88},
		{"79", []string{"sevenbcfbpnrvkkscrjtpctdtb69bvvnvlgsmjltlvs"}, 79},
		{"23", []string{"2ninebvgdcfxtktqjxjqvxfgjdqfhv5threegqtsfhtfxg"}, 23},
		{"11", []string{"onexdchhtxmhsevenbczrslrppneightonenbnhfmbsvdcnzjx1"}, 11},
		{"45", []string{"45"}, 45},
		{"26", []string{"two86"}, 26},
		{"95", []string{"gqznine5gpg"}, 95},
		{"64", []string{"6bmltlrvrgpcfhjhmfiveqzfxptjtwo4zvsqqxgbrdlzsfmtzdd"}, 64},
		{"72", []string{"cfjgdffcgvldsnvkbjqrxhxcl7fjlxdrlrrthreeseventwo"}, 72},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := day_one.FindAndSumWords(&tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
func TestFinalSummation(t *testing.T) {
	values := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"9fiveajhseven",
		"xjhoneighthree",
		"eight96nxcjjddmseightxvgsixfiverrzpvmgnl",
		"8pjkm",
	}
	ans := day_one.FindAndSumWords(&values)
	if ans != 564 {
		t.Errorf("Expected: %d, got: %d", 564, ans)
	}
}

//values := []string{
//	"two1nine",
//	"eightwothree",
//	"abcone2threexyz",
//	"xtwone3four",
//	"4nineeightseven2",
//	"zoneight234",
//	"7pqrstsixteen",
//}
//ans := day_one.FindAndSumWords(&values)
//if ans != 281 {
//	t.Errorf("Expected: %d, got: %d", 281, ans)
//}
