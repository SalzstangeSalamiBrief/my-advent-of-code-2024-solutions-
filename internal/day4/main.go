package main

import (
	"fmt"
	"salzstangeSalamiBrief/adventOfCode24/pkg/fileReader"
)

const fileName string = "exampleInput.txt"

var wordParts []string = []string{"X", "M", "A", "S"}

func main() {
	multidimensionalCharacterArrayInput := fileReader.GetFileContentAsMultidimensionalCharacterArray(fileName)
	indexesOfStartingCharacterX := getIndexesOfStartingCharacterX(multidimensionalCharacterArrayInput)

	sum := 0
	for _, tupleOfIndexes := range indexesOfStartingCharacterX{
		sum += getNumberOfMatchesForStartingCharacter(multidimensionalCharacterArrayInput, tupleOfIndexes)
	}
	fmt.Printf("sum '%v'", sum)
}

func getIndexesOfStartingCharacterX(lines [][]string) [][2]int {
	indexes := make([][2]int, 0)
	for i, line := range lines {
		for j, character := range line {
			if character == "X" {
				indexes = append(indexes, [2]int{i, j})
			}}}
		return indexes}

func getNumberOfMatchesForStartingCharacter(puzzleInput [][]string, startingIndex [2]int) int {
	isHorizontalLeftMatching := getIsHorizontalLeftMatching(puzzleInput[startingIndex[0]], startingIndex[1])
	isHorizontalRightMatching := getIsHorizontalRightMatching(puzzleInput[startingIndex[0]], startingIndex[1])
	isVerticalTopMatching := getIsVerticalTopMatching(puzzleInput, startingIndex[0], startingIndex[1])
	isVerticalBottomMatching := getIsVerticalBottomMatching(puzzleInput, startingIndex[0], startingIndex[1])

	result := isHorizontalLeftMatching + isHorizontalRightMatching+isVerticalTopMatching+isVerticalBottomMatching
	if result > 0{
		fmt.Printf("[row, col]: '[%v, %v]'\n", startingIndex[0], startingIndex[1])
		fmt.Printf("left '%v', right '%v', top '%v', bottom '%v'\n", isHorizontalLeftMatching, isHorizontalRightMatching, isVerticalTopMatching, isVerticalBottomMatching)
	}

	return result
}

func getIsHorizontalRightMatching(line []string, col int) int {
	if getIsHorizontalRightOutOfBound(line, col){
		return 0
	}

	for i, part := range wordParts {
		isCharacterMatching := line[col+i] == part
		if !isCharacterMatching{
			return 0;
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
		if !isCharacterMatching{
			return 0
		}
	}

	return 1
}

func getIsVerticalBottomMatching(puzzleInput [][]string, row int, col int) int{
	if getIsVerticalBottomOutOfBound(puzzleInput, row){
		return 0
	}

		for i, part := range wordParts{
			isCharacterMatching := puzzleInput[row+i][col] == part
			if !isCharacterMatching{
				return 0
			}		
	}

	return 1
}

func getIsVerticalTopMatching(puzzleInput [][]string, row int, col int) int{
	if getIsVerticalTopOutOfBound(row){
		return 0
	}

		for i, part := range wordParts{
			isCharacterMatching := puzzleInput[row-i][col] == part
			if !isCharacterMatching{
				return 0
			}		
	}

	return 1
}

func getIsHorizontalLeftOutOfBound(col int) bool{
	return col - len(wordParts) <= 0
}

func getIsHorizontalRightOutOfBound(line []string, col int) bool{
	return col + len(wordParts) >= len(line)
}

func getIsVerticalBottomOutOfBound(puzzleInput [][]string, row int) bool{
	return row + len(wordParts) >= len(puzzleInput)
}

func getIsVerticalTopOutOfBound(row int) bool{
	return row - len(wordParts) <= 0
}