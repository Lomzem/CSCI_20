/***************************************
File Merge
Author: Lawjay Lee
Date Completed: 2023-11-16
Description: Reads two files, concats a line
from each, then prints concated string to stdout
***************************************/

package main

import (
	"bufio"
	"os"
)

func getFileLine(filepath string, lineNumber int) string {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		panic("Unable to open " + filepath)
	}

	scanner := bufio.NewScanner(file)

	for i := 0; i < lineNumber; i++ {
		scanner.Scan()
	}

	return scanner.Text()
}

func main() {
	var file1_line string = getFileLine("./File1.txt", 3)
	var file2_line string = getFileLine("./File2.txt", 2)

	var newString string = file1_line + " and " + file2_line

	newFile, err := os.Create("Merged.txt")
	defer newFile.Close()

	if err != nil {
		panic("Unable to create Merged.txt")
	}

	writer := bufio.NewWriter(newFile)
	_, err = writer.WriteString(newString)

	if err != nil {
		panic("Unable to fully write to Merged.txt buffer")
	}

	err = writer.Flush()
	if err != nil {
		panic("Unable to write to Merged.txt file")
	}
}
