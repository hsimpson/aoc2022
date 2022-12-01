package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func ReadFileLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
