package main

import (
	"fmt"
	"log"
	"regexp"
	"salzstangeSalamiBrief/adventOfCode24/pkg/fileReader"
	"strconv"
)

const fileName string = "input.txt"

type Direction int

const (
	UNKNOWN Direction = 0
	ASC     Direction = 1
	DESC    Direction = 2
	EQ      Direction = 3
)

func main() {
	puzzleMap := getPuzzleMap(fileName)
	//PartOne(puzzleMap)
	PartTwo(puzzleMap)

}

func PartOne(puzzleMap [][]int) {
	saveEntries := filterReportsBySafety(puzzleMap, true, false)
	numberOfSaveEntries := len(saveEntries)
	fmt.Printf("Part One: %v\n", numberOfSaveEntries)
}

func PartTwo(puzzleMap [][]int) {
	saveEntries := filterReportsBySafety(puzzleMap, true, true)
	numberOfSaveEntries := len(saveEntries)
	fmt.Printf("Part Two: %v", numberOfSaveEntries)
}

func getPuzzleMap(fileName string) [][]int {
	spacingRegExp := regexp.MustCompile(`\s{1,}`)
	fileLines := fileReader.GetFileContentLines(fileName)

	if len(fileLines) == 0 || len(fileLines[0]) == 0 {
		log.Print("The input is empty")
		return [][]int{}
	}

	puzzleMap := make([][]int, len(fileLines))

	for i, line := range fileLines {
		reportStrings := spacingRegExp.Split(line, -1)
		puzzleMap[i] = make([]int, len(reportStrings))
		for j, reportString := range reportStrings {
			reportInt, err := strconv.Atoi(reportString)
			if err != nil {
				log.Panicln(err.Error())
			}

			puzzleMap[i][j] = reportInt
		}
	}

	return puzzleMap
}

func filterReportsBySafety(reportsMap [][]int, shouldSave bool, shouldUseToleranceLevel bool) [][]int {
	var saveEntries [][]int

	for _, reports := range reportsMap {
		isEntrySave := checkIfReportsAreSave(reports, shouldUseToleranceLevel)
		//if !isEntrySave {
		fmt.Printf("%v: %v\n", isEntrySave, reports)

		//}
		if isEntrySave == shouldSave {
			saveEntries = append(saveEntries, reports)
		}
	}

	return saveEntries
}

func checkIfReportsAreSave(reports []int, shouldUseToleranceLevel bool) bool {
	if len(reports) == 0 {
		return true
	}

	expectedDirection := UNKNOWN
	hasIgnoredPreviousBadLevel := shouldUseToleranceLevel == false
	currentReports := reports

	for i := 0; i < len(currentReports)-1; i += 1 {
		firstItem := currentReports[i]
		secondItem := currentReports[i+1]
		currentDirection := getDirectionOfReport(firstItem, secondItem)
		if i == 0 {
			expectedDirection = currentDirection
		}

		isSameDirection := currentDirection == expectedDirection
		areReportsInRange := areReportInRange(firstItem, secondItem)
		if isSameDirection && areReportsInRange {
			continue
		}

		if !shouldUseToleranceLevel {
			return false
		}

		if hasIgnoredPreviousBadLevel {
			return false
		}

		hasIgnoredPreviousBadLevel = true
		currentReports = append(currentReports[:i+1], currentReports[i+2:]...)
		i -= 1
	}

	return true
}

func getDirectionOfReport(a int, b int) Direction {
	if a == b {
		return EQ
	}

	if a > b {
		return DESC
	}

	return ASC
}

func areReportInRange(a int, b int) bool {
	diff := a - b
	if diff < 0 {
		diff *= -1
	}

	if 1 <= diff && diff <= 3 {
		return true
	}

	return false
}
