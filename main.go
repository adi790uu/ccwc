package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sentence, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parts := strings.Split(sentence, " ")
	
	if parts[0] == "ccwc" {
		fmt.Println("ccwc")
		if parts[1] == "-c" {
			fileSize, err := getFileSize(strings.Split(parts[len(parts) - 1], "\n")[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(fileSize)
		}
	}else {
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}