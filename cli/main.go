package main

import (
	"cli/cli-tool"
	"flag"
	"fmt"
	"io"
	"os"
)

func getFileContent(args []string) string {
	var filePath string
	var fileContent string

	if len(args) > 0 {
		filePath = args[0]
		fileContent = readFileContent(filePath)
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error in reading input")
			panic(err.Error())
		}
		fileContent = string(stdin)
	}

	return fileContent
}

func readFileContent(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	return string(file)
}

func main() {
	bytes := flag.Bool("c", false, "count bytes in the file")
	lines := flag.Bool("l", false, "count lines in the file")
	words := flag.Bool("w", false, "count words in the file")
	chars := flag.Bool("m", false, "count characters in the file")

	flag.Parse()

	args := flag.Args()

	fileContent := getFileContent(args)

	cli.FileReaderCLIHandler(bytes, lines, words, chars, fileContent)
}
