package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func countBytes(r io.Reader) (int64, error) {
	var count int64
	buf := make([]byte, 8192)
	for {
		n, err := r.Read(buf)
		count += int64(n)
		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}
	}
}

func countLines(r io.Reader) (int64, error) {
	scanner := bufio.NewScanner(r)
	var lines int64
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

func countWords(r io.Reader) (int64, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var words int64
	for scanner.Scan() {
		words++
	}
	return words, scanner.Err()
}

func countCharacters(r io.Reader) (int64, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	var chars int64
	for scanner.Scan() {
		chars++
	}
	return chars, scanner.Err()
}

func processInput(r io.Reader, flag string) (int64, error) {
	switch flag {
	case "-c":
		return countBytes(r)
	case "-l":
		return countLines(r)
	case "-w":
		return countWords(r)
	case "-m":
		return countCharacters(r)
	default:
		return 0, fmt.Errorf("invalid flag: %s", flag)
	}
}

func main() {
	var flag, filename string
	var input io.Reader
	var err error

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		input = os.Stdin
	}

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-c" || args[i] == "-l" || args[i] == "-w" || args[i] == "-m" {
			flag = args[i]
		} else {
			filename = args[i]
			break
		}
	}

	if filename != "" && input == nil {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	if input == nil {
		fmt.Println("Usage: ccwc [-c|-l|-w|-m] [filename]")
		fmt.Println("Or: cat file | ccwc [-c|-l|-w|-m]")
		os.Exit(1)
	}

	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	reader := bytes.NewReader(content)

	if flag == "" {
		lines, _ := countLines(bytes.NewReader(content))
		words, _ := countWords(bytes.NewReader(content))
		byteCount, _ := countBytes(bytes.NewReader(content))
		fmt.Printf("%d %d %d", lines, words, byteCount)
	} else {
		// Process based on flag
		count, err := processInput(reader, flag)
		if err != nil {
			fmt.Println("Error processing input:", err)
			os.Exit(1)
		}
		fmt.Printf("%d", count)
	}

	if filename != "" {
		fmt.Printf(" %s", filename)
	}
	fmt.Println()
}