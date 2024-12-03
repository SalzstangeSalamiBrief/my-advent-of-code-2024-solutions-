package main

import (
	"fmt"
	"log"
	"regexp"
	"salzstangeSalamiBrief/adventOfCode24/pkg/fileReader"
	"strconv"
	"strings"
)

const fileName string = "input.txt"

func main() {
	input := getCorruptInstructionString(fileName)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	instructions := getMulInstructions(input)
	sum := getSumOfMulInstructions(instructions)
	fmt.Printf("Part One: %v\n", sum)
}

func partTwo(input string) {
	cleanedInstructions := getDoInstructions(input)
	instructions := getMulInstructions(cleanedInstructions)
	sum := getSumOfMulInstructions(instructions)
	fmt.Printf("Part Two: %v\n", sum)

}

func getCorruptInstructionString(fileName string) string {
	lines := fileReader.GetFileContentLines(fileName)
	if len(lines) != 1 {
		log.Panicln("The number of lines is not equal to 1")
	}

	return lines[0]
}

func getMulInstructions(line string) []string {
	mulRegexp := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	directions := mulRegexp.FindAllString(line, -1)
	return directions
}

func getSumOfMulInstructions(instructions []string) int {
	sum := 0
	for _, instruction := range instructions {
		sum += performMulInstruction(instruction)
	}

	return sum
}

func performMulInstruction(instruction string) int {
	numberRegexp := regexp.MustCompile(`\d{1,3}`)
	numberStrings := numberRegexp.FindAllString(instruction, 2)

	if len(numberStrings) != 2 {
		log.Panicf("The number of numbers is not equal to 2: '%v'\n", numberStrings)
	}

	numberOne, err := strconv.Atoi(numberStrings[0])
	if err != nil {
		log.Panicln(err.Error())
	}

	numberTwo, err := strconv.Atoi(numberStrings[1])
	if err != nil {
		log.Panicln(err.Error())
	}

	fmt.Printf("Get sum of '%v * %v'\n", numberOne, numberTwo)
	return numberOne * numberTwo
}

func getDoInstructions(input string) string {
	doRegexp := regexp.MustCompile(`do\(\)`)
	doMatches := doRegexp.Split(input, -1)
	dontRegexp := regexp.MustCompile(`don't\(\).+`)
	for i, doMatch := range doMatches {
		doMatches[i] = dontRegexp.ReplaceAllString(doMatch, "")
	}
	return strings.Join(doMatches, "")
}
