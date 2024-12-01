package main

import (
	"fmt"
	"log"
	"regexp"
	"salzstangeSalamiBrief/adventOfCode24/pkg/fileReader"
	"strconv"
)

const fileName = "input.txt"

func main() {
	fileLines := fileReader.GetFileContentLines(fileName)
	firstList, secondList := getLists(fileLines)
	sortableFirstList := SortableIntSlice(firstList)
	(&sortableFirstList).SortListAscBubbleSort()
	sortableSecondList := SortableIntSlice(secondList)
	(&sortableSecondList).SortListAscBubbleSort()
	PartOne(sortableFirstList, sortableSecondList)
	PartTwo(sortableFirstList, sortableSecondList)
}

func PartOne(firstList SortableIntSlice, secondList SortableIntSlice) {
	distances := getDistancesBetweenLists(firstList, secondList)
	sumOfDistances := getSumOfDistances(distances)
	fmt.Printf("Part one: %v\n", sumOfDistances)
}

func PartTwo(firstList SortableIntSlice, secondList SortableIntSlice) {
	weightedValues := getWeightedValues(firstList, secondList)
	sumOfWeights := getSumOfDistances(weightedValues)
	fmt.Printf("Part two: %v\n", sumOfWeights)
}

func getLists(lines []string) ([]int, []int) {
	spacingRegExp := regexp.MustCompile(`\s{1,}`)

	firstList := make([]int, len(lines))
	secondList := make([]int, len(lines))

	for i, line := range lines {
		items := spacingRegExp.Split(line, 2)
		if len(items) != 2 {
			log.Panicln("The number of lines is not equal to two.")
		}

		firstItemString, secondItemString := items[0], items[1]

		firstItemInt, err := strconv.Atoi(firstItemString)
		if err != nil {
			log.Panicln(err.Error())
		}

		secondItemInt, err := strconv.Atoi(secondItemString)
		if err != nil {
			log.Panicln(err.Error())
		}

		firstList[i] = firstItemInt
		secondList[i] = secondItemInt
	}

	return firstList, secondList

}

func getDistancesBetweenLists(firstList SortableIntSlice, secondList SortableIntSlice) []int {
	if len(firstList) != len(secondList) {
		log.Panicf("The length of the lists does not match: %v != %v\n", len(firstList), len(secondList))
	}

	distances := make([]int, len(firstList))
	for i := 0; i < len(firstList); i += 1 {
		distances[i] = firstList[i] - secondList[i]
		if distances[i] < 0 {
			distances[i] = distances[i] * -1
		}
	}
	return distances
}

func getSumOfDistances(distances []int) int {
	sum := 0

	for _, distance := range distances {
		sum += distance
	}

	return sum
}

func getWeightedValues(firstList SortableIntSlice, secondList SortableIntSlice) []int {
	weights := make([]int, len(firstList))
	for i, value := range firstList {
		occurrences := getNumberOfOccurrenceInList(value, secondList)
		weights[i] = value * occurrences
	}

	return weights
}

func getNumberOfOccurrenceInList(target int, referenceList []int) int {
	var occurence int

	for _, reference := range referenceList {
		if target == reference {
			occurence += 1
		}
	}

	return occurence
}

type SortableIntSlice []int

func (input *SortableIntSlice) SortListAscBubbleSort() {
	maxLength := len(*input)
	for i := 0; i < maxLength; i += 1 {
		for j := i + 1; j < maxLength; j += 1 {
			if (*input)[i] > (*input)[j] {
				(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
			}
		}
	}
}
