package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	testFiles := []string{
		"file.txt",
		"file2.txt",
	}
	c := make(chan string)
	for _, val := range testFiles {
		go processFile(val, c)
	}
	for range len(testFiles) {
		fmt.Println(<-c)
	}
}

func processFile(path string, c chan string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var final strings.Builder
	data := make([]byte, 2)
	var read int
	for ok := true; ok; ok = (read != 0) {
		read, _ = file.Read(data)
		final.Write(data[:])
		fmt.Println(string(data[:]))
	}
	c <- fmt.Sprintf("Done processing file %s %s", path, final.String())
}
