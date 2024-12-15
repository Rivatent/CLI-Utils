package count

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	UsageMsg  = "usage: ./myWc [[-l] [-w] [-m]] filename1 filename2...\n" + "only one flag can be specified"
	FlagError = "wrong flags or their combination specified"
)

func countWordsInString(str string) int {
	return len(strings.Fields(str))
}

func CountWords(filePath string) (int, error) {
	wordsCounter := 0
	file, err := os.Open(filePath)
	if err != nil {
		return wordsCounter, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wordsCounter += countWordsInString(scanner.Text())
	}
	return wordsCounter, nil
}

// Counts chars for file
func CountChars(filePath string) (int, error) {
	charsCounter := 0
	file, err := os.ReadFile(filePath)
	if err != nil {
		return charsCounter, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	for range string(file) {
		charsCounter++
	}

	return charsCounter, nil
}

// Counts newlines for file
func CountLines(filePath string) (int, error) {
	linesCounter := 0
	file, err := os.Open(filePath)
	if err != nil {
		return linesCounter, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesCounter++
	}

	return linesCounter, nil
}

func DefineFunc(linesFlag, wordsFlag, charsFlag bool) (func(string) (int, error), error) {

	var countFunc func(filename string) (int, error)

	switch {
	case linesFlag && !wordsFlag && !charsFlag:
		countFunc = CountLines
	case wordsFlag && !linesFlag && !charsFlag:
		countFunc = CountWords
	case charsFlag && !wordsFlag && !linesFlag:
		countFunc = CountChars
	case !charsFlag && !wordsFlag && !linesFlag:
		countFunc = CountWords
	default:
		return countFunc, errors.New(FlagError)
	}
	return countFunc, nil
}
