package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)

func main() {
	// Check file exists
	file, err := os.Open("/Users/swaschni/Projekte/GO-IDEAS/books.csv")
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	// 10 -> byte representing delimiter aka NewLine
	line, err := reader.ReadString(10)
	for err != io.EOF {

		formatAndPrintLine(line)

		// Read next line
		line, err = reader.ReadString(10)
	}
}

func formatAndPrintLine(line string) {
	split := strings.Split(line, ";")
	for i := 0; i < len(split); i++ {
		fmt.Print(makeNice(split[i]))
	}

	fmt.Print("\n")
}

func makeNice(text string) string {
	text = strings.TrimSpace(text);
	i := utf8.RuneCountInString(text);
	for i != 20 {
		text += " ";
		i = utf8.RuneCountInString(text);
	}
	return text;
}

func check(myError error) {
	if myError != nil {
		panic(myError)
	}
}