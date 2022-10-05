package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	readFile, err := os.Open("File.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()

	fmt.Println(lines[0])
	fmt.Println(lines[len(lines)-1])
	for index, mots := range lines {
		if mots == "before" {
			v := lines[index+1]
			a, _ := strconv.Atoi(v)
			fmt.Println(lines[a-1])
		}
		if mots == "now " {
			b := lines[index-1]
			c := int(b[1]) / len(lines)
			fmt.Println(lines[c-1])

		}

	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}

// generate a random nunber
