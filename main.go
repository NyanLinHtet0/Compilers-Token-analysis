package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.py")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	var counter int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			println(string(ch))
		}
		println(string(' '))
		counter += 1
	}
}
