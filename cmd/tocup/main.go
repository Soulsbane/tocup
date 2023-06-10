package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

const CURRENT_INTERFACE_VERSION = "90002"

func getFileLines(fileName string) ([]string, error) {
	fileBytes, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}

func writeToFile(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func getTocFileName() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	return path.Join(dir, path.Base(dir)+".toc")

}
func main() {
	var outputLines []string

	tocFileName := getTocFileName()
	lines, err := getFileLines(tocFileName)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		if strings.Contains(line, "Interface:") {
			re := regexp.MustCompile(`(\w+)(\d+)`)

			replacedValue := re.ReplaceAllString(line, CURRENT_INTERFACE_VERSION)
			outputLines = append(outputLines, replacedValue)
		} else {
			outputLines = append(outputLines, line)
		}
	}

	writeToFile(outputLines)
}
