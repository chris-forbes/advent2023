package day_one

import (
	"regexp"
	"strconv"
)

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
