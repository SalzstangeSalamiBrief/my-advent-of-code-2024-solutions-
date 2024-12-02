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
	saveEntries := filterReportsBySafety(puzzleMap, true, false)
	numberOfSaveEntries := len(saveEntries)
	fmt.Printf("Part One: %v\n", numberOfSaveEntries)
	saveEntries = filterReportsBySafety(puzzleMap, true, true)
	numberOfSaveEntries = len(saveEntries)
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

	for _, entry := range reportsMap {
		isEntrySave := checkIfReportsAreSave(entry, shouldUseToleranceLevel)
		if isEntrySave == shouldSave {
			saveEntries = append(saveEntries, entry)
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
	isSave := true
	i := 0

	for i < len(reports)-1 {
		currentDirection := getDirectionOfReport(reports[i], reports[i+1])
		if expectedDirection == UNKNOWN {
			expectedDirection = currentDirection
		}

		isSameDirection := currentDirection == expectedDirection
		areReportsInRange := areReportInRange(reports[i], reports[i+1])
		isSave = isSave && isSameDirection && areReportsInRange

		if !isSave {
			if shouldUseToleranceLevel {
				if hasIgnoredPreviousBadLevel {
					return false
				}

				reports = append(reports[:i], reports[i+1:]...)
				hasIgnoredPreviousBadLevel = true
				isSave = true
				if i > 0 {
					i -= 1
				}
				continue
			}

			return false
		}

		i += 1
	}

	return isSave
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

	if 0 <= diff && diff <= 3 {
		return true
	}

	return false
}
