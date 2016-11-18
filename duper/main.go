package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: duper <filename>")
		return
	}
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Println("Open file error:", err)
		return
	}
	foundWords := make(map[string]int, 10)
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.Trim(word, "\",")
		foundWords[word] += 1
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scan file error: ", err)
		return
	}
	log.Println("Duplicate words: ")
	for w, i := range foundWords {
		if i > 1 {
			fmt.Printf("'%s':%d ", w, i)
		}
	}
	fmt.Print("\n")
}
