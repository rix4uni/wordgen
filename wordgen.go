package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define flags
	paths := flag.String("path", "", "Path or list of paths to add (comma-separated or a file)")
	extensions := flag.String("e", "", "Comma separated list of extensions")
	flag.Parse()

	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var baseURLs []string
	for scanner.Scan() {
		baseURLs = append(baseURLs, scanner.Text())
	}

	// Split paths and extensions
	var pathList []string
	if strings.Contains(*paths, ".txt") {
		// If a file is provided, read paths from the file
		file, err := os.Open(*paths)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			pathList = append(pathList, scanner.Text())
		}
	} else {
		// If comma-separated paths are provided, split them
		pathList = strings.Split(*paths, ",")
	}

	extList := strings.Split(*extensions, ",")

	// Generate output
	for _, baseURL := range baseURLs {
		for _, path := range pathList {
			if len(extList) > 0 && extList[0] != "" {
				for _, ext := range extList {
					fmt.Println(baseURL + "/" + path + ext)
				}
			} else {
				fmt.Println(baseURL + "/" + path)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
