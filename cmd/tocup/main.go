package main

import (
	"bufio"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/duke-git/lancet/v2/fileutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

const CurrentInterfaceVersion = "100200"

func writeToFile(fileName string, lines []string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal("Failed to update ", fileName, " with the new interface version!")
	}

	writer := bufio.NewWriter(f)

	for _, line := range lines {
		writer.WriteString(line + "\n")
	}

	writer.Flush()
	fmt.Println("Updated ", fileName, " with the new interface version")
}

func getTocFileName() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	return path.Join(dir, path.Base(dir)+".toc")

}

func replaceInterfaceVersion(interfaceVersion string) {
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

				replacedValue := re.ReplaceAllString(line, interfaceVersion)
				outputLines = append(outputLines, replacedValue)
			} else {
				outputLines = append(outputLines, line)
			}
		}

		writeToFile(tocFileName, outputLines)
	} else {
		fmt.Println("Failed to find TOC file: ", tocFileName)
	}
}

func main() {
	var args programArgs
	var interfaceVersion string

	arg.MustParse(&args)

	if args.InterfaceVersion == "" {
		interfaceVersion = CurrentInterfaceVersion
	} else {
		interfaceVersion = args.InterfaceVersion
	}

	replaceInterfaceVersion(interfaceVersion)
}
