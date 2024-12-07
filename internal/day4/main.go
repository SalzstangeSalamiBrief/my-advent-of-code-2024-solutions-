package main

import (
	"salzstangeSalamiBrief/adventOfCode24/pkg/fileReader"
	"strings"
)

const fileName string = "exampleInput.txt"

var wordParts []string = []string{"X", "M", "A", "S"}
var horizontalJoinedWord = strings.Join(wordParts, "")

func main() {
	multidimensionalCharacterArrayInput := fileReader.GetFileContentAsMultidimensionalCharacterArray(fileName)
	indexesOfStartingCharacterX := getIndexesOfStartingCharacterX(multidimensionalCharacterArrayInput)
}

func getIndexesOfStartingCharacterX(lines [][]string) [][2]int {
	indexes := make([][2]int, len(lines))
	for i, line := range lines {
		for j, character := range line {
			if character == "X" {
				indexes = append(indexes, [2]int{i, j})
			}
		}
	}

	return indexes
}

func getNumberOfMatchesForStartingCharacter(line []string, startingIndex [2]int) int {
	isHorizontalLeftMatching := getIsHorizontalLeftMatching(line, startingIndex[0])
	isHorizontalRightMatching := getIsHorizontalRightMatching(line, startingIndex[0])
}

func getIsHorizontalRightMatching(line []string, col int) bool {
	if col > len(line)-len(wordParts) {
		return false
	}

	isWordMatching := true
	for i, part := range wordParts {
		isCharacterMatching := line[col+i] == part
		isWordMatching = isWordMatching && isCharacterMatching
	}

	return isWordMatching
}

func getIsHorizontalLeftMatching(line []string, col int) bool {
	if col-len(wordParts)-1 < 0 {
		return false
	}

	isWordMatching := true
	for i, part := range wordParts {
		isCharacterMatching := line[col-i] == part
		isWordMatching = isWordMatching && isCharacterMatching
	}

	return isWordMatching
}
