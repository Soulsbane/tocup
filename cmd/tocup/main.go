package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

const CURRENT_INTERFACE_VERSION = "90002"

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

func replaceInterfaceVersion() {
	var outputLines []string

	tocFileName := getTocFileName()

	if fileutil.IsExist(tocFileName) {
		lines, err := fileutil.ReadFileByLine(tocFileName)

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
	} else {
		fmt.Println("Failed to find TOC file: ", tocFileName)
	}
}

func main() {
	replaceInterfaceVersion()
}
