package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise4.9:%v\n", err)
			continue
		}
		wordFreq(f)
		f.Close()
	}
}

func wordFreq(file *os.File) {
	freq := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		freq[word]++
	}
	if input.Err() != nil {
		fmt.Fprintln(os.Stderr, input.Err())
		os.Exit(1)
	}
	for word, n := range freq {
		fmt.Printf("%s is %d\n", word, n)
	}
}
