package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "aaabbcddd"
	s2 := "aaaaaabbcdddhhhhhwwwwwwwwwddddddddvvvvvvvv"
	s3 := "pppppppppownnnnnnnnnffqqqqqqqqssssssssssmvxcvud"
	fmt.Println(s1, stringCompression(s1))
	fmt.Println(s2, stringCompression(s2))
	fmt.Println(s3, stringCompression(s3))
}

func stringCompression(sequence string) string {
	data := make(map[string]int)
	for _, val := range sequence {
		data[string(val)] = data[string(val)] + 1
	}

	var output strings.Builder
	for key, val := range data {
		fmt.Fprintf(&output, "%v%d", key, val)
	}
	return output.String()
}
