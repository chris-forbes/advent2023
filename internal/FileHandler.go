package internal

import (
	"bufio"
	"log"
	"os"
)

func ParseToStringArray(path string) *[]string {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	sc := bufio.NewScanner(f)
	var fileValues []string
	for sc.Scan() {
		line := sc.Text()
		fileValues = append(fileValues, line)
	}
	return &fileValues
}
