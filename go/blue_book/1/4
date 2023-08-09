// Modify dup2 to print the  names of all files in which each suplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles)
			f.Close()
		}
	}
	for line, files := range dupFiles {
		if len(files) > 1 {
			fmt.Printf("%s\t%s\n", files, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		found := false
		for _, file := range dupFiles[input.Text()] {
			if file == f.Name() {
				found = true
			}
		}
		if !found {
			dupFiles[input.Text()] = append(dupFiles[input.Text()], f.Name())
		}
	}
}
