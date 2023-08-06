package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ComicJsonInfo struct {
	Title      string `json:"safe_title"`
	Transcript string
}

type ComicInfo struct {
	URL  string
	JSON ComicJsonInfo
}

func main() {
	// Declare search term
	var searchTerm, sep string
	for i := 1; i < len(os.Args); i++ {
		searchTerm += sep + os.Args[i]
		sep = " "
	}

	// Open and read file
	fileData, fileErr := ioutil.ReadFile("lib/comics.txt")
	if fileErr != nil {
		fmt.Printf("error: %v\n", fileErr)
		return
	}

	// Marshal JSON to objects
	var comicLibrary []ComicInfo
	if err := json.Unmarshal(fileData, &comicLibrary); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// Loop through objects
	var matches []ComicInfo
	for _, comic := range comicLibrary {
		if strings.Contains(comic.JSON.Transcript, searchTerm) {
			matches = append(matches, comic)
		}
	}

	// print matched results of URL and transcript
	fmt.Printf("Hey great news you got %d matches\n", len(matches))
	for _, comic := range matches {
		fmt.Println()
		fmt.Println("**********************")
		fmt.Println("**********************")
		fmt.Println()
		fmt.Println(comic.URL)
		fmt.Println(comic.JSON.Transcript)
	}
}
