package main

import (
	"flag"
	"github.com/chris-forbes/advent-23/internal"
	"github.com/chris-forbes/advent-23/pkg/day_one"
	"log"
)

func main() {
	log.Println("Advent 2023")
	inputFile := flag.String("input-file", "", "-input-file=<file>")
	verbose := flag.Bool("verbose", false, "-verbose")
	dayOne := flag.Bool("day-one", false, "run day one")
	flag.Parse()
	if *verbose {
		log.Println("Input file: ", *inputFile)
	}

	if *dayOne {
		sum := day_one.FindAndSum(internal.ParseToStringArray(*inputFile))
		log.Printf("Day one A answer: %d", sum)
		sumB := day_one.FindAndSumWords(internal.ParseToStringArray(*inputFile))
		log.Printf("Day one B answer: %d", sumB)
	}

}
