package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


var cubeLimits = map[string]int{"red": 12, "green": 13, "blue": 14}

func openFile(path string) *os.File {
	file, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}
	return file	
}

func checkReveals(reveal string) bool {
	colors := strings.Split(reveal, ",")
	for _, v := range colors {
		cubes := (strings.Split(strings.Trim(v, " "), " "))
		count, _ := strconv.Atoi(cubes[0])
		if  count > cubeLimits[cubes[1]] {
			return false
		}
	}
	return true
}

func scanLine(line string) bool {
	line = line[strings.Index(line, ":") + 1:]
	reveals := strings.Split(line, ";")
	for i := 0; i < len(reveals); i++ {
		valid := checkReveals(reveals[i])
		if !valid {
			return false
		}
	}
	return true
}


func main() {
	scanner := bufio.NewScanner(openFile("../input.txt"))

	gameIdSum := 0
	gameCount := 1
	for scanner.Scan() {
		line := scanner.Text()
		validGame := scanLine(line)
		if validGame {
			gameIdSum += gameCount
		}
		gameCount += 1
	}
	fmt.Println(gameIdSum)
}