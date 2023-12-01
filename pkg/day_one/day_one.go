package day_one

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var _smallNumbers = []string{
	"one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

func FirstAndLast(value string) string {
	if value == "" {
		return ""
	}
	re := regexp.MustCompile("[0-9+]")
	numbers := re.FindAllString(value, -1)
	if len(numbers) == 1 {
		value = numbers[0] + numbers[0]
		return value
	}
	max := numbers[len(numbers)-1]
	min := numbers[0]
	return min + max
}

func FindAndSum(values *[]string) int {
	sum := 0
	for _, v := range *values {
		valueString := FirstAndLast(v)
		intVal, _ := strconv.Atoi(valueString)
		sum += intVal
	}
	return sum
}

func FindAndSumWords(values *[]string) int {
	var updatedValues []string
	for _, v := range *values {
		var replacedValue = replaceWordWithNumber(v)
		updatedValues = append(updatedValues, replacedValue)
	}
	return FindAndSum(&updatedValues)
}
func replaceWordWithNumber(line string) string {
	var replaced = line
	var firstMatchIndex = math.MaxInt
	var lastMatch = math.MaxInt
	for i, num := range _smallNumbers {
		// Check if the current element is a substring of the target string
		if strings.Contains(replaced, num) {
			var index = strings.Index(replaced, num)
			if index < lastMatch {
				lastMatch = index
				firstMatchIndex = i
			}
		}
	}

	for i, v := range _smallNumbers {
		if i == firstMatchIndex {
			replaced = strings.Replace(replaced, v, strconv.Itoa(i+1), 1)
			break
		}
	}

	for i, v := range _smallNumbers {
		replaced = strings.Replace(replaced, v, strconv.Itoa(i+1), 1)
	}

	return replaced
}
