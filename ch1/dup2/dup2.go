package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {

		for _, arg := range files {
			fd, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(fd, counts, arg)
			fd.Close()
			fmt.Printf("Evaluating file %s\n", arg)
			printLines(counts)
			fmt.Printf("Done\n")
			counts = make(map[string]map[string]int)
		}
	}
}
func printLines(counts map[string]map[string]int) {
	for file, lineNumberMap := range counts {
		for line, n := range lineNumberMap {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file, n, line)
			}
		}
	}
}
func countLines(fd *os.File, counts map[string]map[string]int, file string) {
	input := bufio.NewScanner(fd)
	for input.Scan() {
		if counts[file] == nil {
			counts[file] = map[string]int{}
		}
		counts[file][input.Text()]++ //save results in counts
	}

}
