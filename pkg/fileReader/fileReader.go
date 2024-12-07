package fileReader

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func GetFileContentAsString(fileName string) string {
	if fileName == "" {
		log.Panicln("fileName is missing")
	}

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicln(err.Error())
	}

	return string(bytes)
}

func GetFileContentLines(fileName string) []string {
	fileContent := GetFileContentAsString(fileName)
	lineEndRegExp := regexp.MustCompile(`\r\n`)
	return lineEndRegExp.Split(fileContent, -1)
}

func GetFileContentAsMultidimensionalCharacterArray(fileName string) [][]string {
	fileContent := GetFileContentAsString(fileName)
	lineEndRegExp := regexp.MustCompile(`\r\n`)
	lines := lineEndRegExp.Split(fileContent, -1)

	result := make([][]string, len(lines))
	for _, line := range lines {
		currentCharacterLine := strings.Split(line, "")
		result = append(result, currentCharacterLine)
	}

	return result
}
