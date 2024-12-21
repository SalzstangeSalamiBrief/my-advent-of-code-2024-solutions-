package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"salzstangeSalamiBrief/adventOfCode24/pkg/fileReader"
	"strings"
)

const fileName string = "exampleInput.txt"

var wordParts []string = []string{"X", "M", "A", "S"}

func main() {
	multidimensionalCharacterArrayInput := fileReader.GetFileContentAsMultidimensionalCharacterArray(fileName)
	indexesOfStartingCharacterX := getIndexesOfStartingCharacterX(multidimensionalCharacterArrayInput)

	sum := 0
	matches := [][2]int{}
	for _, tupleOfIndexes := range indexesOfStartingCharacterX {
		result := getNumberOfMatchesForStartingCharacter(multidimensionalCharacterArrayInput, tupleOfIndexes)
		if result > 0 {
			matches = append(matches, tupleOfIndexes)
		}
		sum += result
	}
	fmt.Printf("sum '%v'", sum)
	writeMatchesToMarkdownTable(multidimensionalCharacterArrayInput, matches, fileName)
}

func getIndexesOfStartingCharacterX(lines [][]string) [][2]int {
	indexes := make([][2]int, 0)
	for i, line := range lines {
		for j, character := range line {
			if character == "X" {
				indexes = append(indexes, [2]int{i, j})
			}
		}
	}
	return indexes
}

func getNumberOfMatchesForStartingCharacter(puzzleInput [][]string, startingIndex [2]int) int {
	isHorizontalLeftMatching := getIsHorizontalLeftMatching(puzzleInput[startingIndex[0]], startingIndex[1])
	isHorizontalRightMatching := getIsHorizontalRightMatching(puzzleInput[startingIndex[0]], startingIndex[1])
	isVerticalTopMatching := getIsVerticalTopMatching(puzzleInput, startingIndex[0], startingIndex[1])
	isVerticalBottomMatching := getIsVerticalBottomMatching(puzzleInput, startingIndex[0], startingIndex[1])
	isDiagonalToTopLeftMatching := getIsDiagonalToTopLeftMatching(puzzleInput, startingIndex[0], startingIndex[1])
	isDiagonalToTopRightMatching := getIsDiagonalToRopRightMatching(puzzleInput, startingIndex[0], startingIndex[1])
	isDiagonalToBottomleftMatching := getIsDiagonalToBottomLeftMatching(puzzleInput, startingIndex[0], startingIndex[1])
	isDiagonalToBottomRightMatching := getIsDiagonalToBottomRightMatching(puzzleInput, startingIndex[0], startingIndex[1])
	result := isHorizontalLeftMatching +
		isHorizontalRightMatching +
		isVerticalTopMatching +
		isVerticalBottomMatching +
		isDiagonalToBottomRightMatching +
		isDiagonalToBottomleftMatching +
		isDiagonalToTopLeftMatching +
		isDiagonalToTopRightMatching
	if result > 0 {
		fmt.Printf("[row, col]: '[%v, %v]'\n", startingIndex[0], startingIndex[1])
		fmt.Printf("left '%v', right '%v', top '%v', bottom '%v'\n", isHorizontalLeftMatching, isHorizontalRightMatching, isVerticalTopMatching, isVerticalBottomMatching)
		fmt.Printf("top-left '%v', top-right '%v', bottom-left '%v', bottom-right '%v'\n", isDiagonalToTopLeftMatching, isDiagonalToTopRightMatching, isDiagonalToBottomleftMatching, isDiagonalToBottomRightMatching)
	}

	return result
}

func getIsHorizontalRightMatching(line []string, col int) int {
	if getIsHorizontalRightOutOfBound(line, col) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := line[col+i] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsHorizontalLeftMatching(line []string, col int) int {
	if getIsHorizontalLeftOutOfBound(col) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := line[col-i] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsVerticalBottomMatching(puzzleInput [][]string, row int, col int) int {
	if getIsVerticalBottomOutOfBound(puzzleInput, row) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := puzzleInput[row+i][col] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsVerticalTopMatching(puzzleInput [][]string, row int, col int) int {
	if getIsVerticalTopOutOfBound(row) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := puzzleInput[row-i][col] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsDiagonalToTopLeftMatching(puzzleInput [][]string, row int, col int) int {
	if getIsHorizontalLeftOutOfBound(col) || getIsVerticalTopOutOfBound(row) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := puzzleInput[row-i][col-i] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsDiagonalToRopRightMatching(puzzleInput [][]string, row int, col int) int {
	if getIsHorizontalRightOutOfBound(puzzleInput[row], col) || getIsVerticalTopOutOfBound(row) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := puzzleInput[row-i][col+i] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsDiagonalToBottomLeftMatching(puzzleInput [][]string, row int, col int) int {
	if getIsHorizontalLeftOutOfBound(col) || getIsVerticalBottomOutOfBound(puzzleInput, row) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := puzzleInput[row+i][col-i] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsDiagonalToBottomRightMatching(puzzleInput [][]string, row int, col int) int {
	if getIsHorizontalRightOutOfBound(puzzleInput[row], col) || getIsVerticalBottomOutOfBound(puzzleInput, row) {
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := puzzleInput[row+i][col+i] == part
		if !isCharacterMatching {
			return 0
		}
	}

	return 1
}

func getIsHorizontalLeftOutOfBound(col int) bool {
	return col-len(wordParts) < 0
}

func getIsHorizontalRightOutOfBound(line []string, col int) bool {
	return col+len(wordParts) > len(line)
}

func getIsVerticalBottomOutOfBound(puzzleInput [][]string, row int) bool {
	return row+len(wordParts) > len(puzzleInput)
}

func getIsVerticalTopOutOfBound(row int) bool {
	return row-len(wordParts) < 0
}

func writeMatchesToMarkdownTable(puzzleInput [][]string, matches [][2]int, filename string) {
	lines := make([][]string, len(puzzleInput)+2)

	for i := 0; i < len(lines); i += 1 {
		line := make([]string, len(puzzleInput[0])+1)
		for j := range len(line[0]) {
			if j == 0 {
				line[j] = fmt.Sprintf("%v", j-1)
			} else {
				line[j] = " "
			}
		}

		lines[i] = line
	}

	for i := range len(lines[0]) {
		if i == 0 {
			lines[0][0] = "X"
		} else {
			lines[0][i] = fmt.Sprintf("%v", i-1)
		}

		lines[1][i] = "-"
	}

	for i := 2; i < len(lines); i += 1 {
		lines[i][0] = fmt.Sprintf("%v", i-2)
	}

	for _, match := range matches {
		lines[match[0]+2][match[1]+1] = wordParts[0]
		// TODO ADD DIRECTIONS?
	}

	var stringToWriteToFile string
	for _, line := range lines {
		stringToWriteToFile += fmt.Sprintf("|%v|\n", strings.Join(line, "|"))
	}

	bytes := []byte(stringToWriteToFile)

	err := os.WriteFile(fmt.Sprintf("%v.result.table.md", strings.TrimSuffix(filename, filepath.Ext(filename))), bytes, 0644)
	if err != nil {
		log.Panic(err.Error())
	}
}
