package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	targetFile       = "code"
	snippetEnding    = ".snippet"
	lineNumberEnding = ".lineno"
)

func main() {
	start := flag.String("start", "", "the starting delimiter, uncommented")
	end := flag.String("end", "", "the ending delimiter, uncommented")
	sourceFile := flag.String("file", "", "the source file")

	flag.Parse()

	// this could be parameterised at some point
	*start = "// " + *start
	*end = "// " + *end

	file, err := os.Open(*sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	codeDest, err := os.Create(targetFile + snippetEnding)
	if err != nil {
		log.Fatal(err)
	}
	defer codeDest.Close()

	lineDest, err := os.Create(targetFile + lineNumberEnding)
	if err != nil {
		log.Fatal(err)
	}
	defer lineDest.Close()

	scanner := bufio.NewScanner(file)

	// find start
	for lineno := gotoStart(scanner, *start); lineno != -1; lineno = gotoStart(scanner, *start) {
		// write line number
		_, _ = lineDest.WriteString(strconv.Itoa(lineno))
		// write source cod
		_, _ = codeDest.WriteString(strings.ReplaceAll(scanner.Text(), *start, "// ...") + "\n")

		// scan until end and write
		for scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) == *end {
				_, _ = codeDest.WriteString(strings.ReplaceAll(scanner.Text(), *end, "// ...") + "\n")
				break
			}
			_, _ = codeDest.WriteString(scanner.Text() + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	_ = codeDest.Sync()
	_ = lineDest.Sync()
}

func gotoStart(s *bufio.Scanner, start string) (linenumber int) {
	for s.Scan() {
		if strings.TrimSpace(s.Text()) == start {
			return linenumber
		}
		linenumber++
	}
	return -1
}
