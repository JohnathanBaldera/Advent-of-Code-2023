package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openFile(path string) *os.File {
	file, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}
	return file	
}

func checkReveals(reveal string, cubeLimits map[string]int) {
	colors := strings.Split(reveal, ",")
	for _, v := range colors {
		cubes := (strings.Split(strings.Trim(v, " "), " "))
		count, _ := strconv.Atoi(cubes[0])
		if count > cubeLimits[cubes[1]] {
			cubeLimits[cubes[1]] = count
		}
	}
}

func scanLine(line string) int {
	var cubeLimits = map[string]int{"red": 0, "green": 0, "blue": 0}
	line = line[strings.Index(line, ":") + 1:]
	reveals := strings.Split(line, ";")
	for i := 0; i < len(reveals); i++ {
		checkReveals(reveals[i], cubeLimits)
	}
	return cubeLimits["red"] * cubeLimits["green"] * cubeLimits["blue"]
}


func main() {
	scanner := bufio.NewScanner(openFile("../input.txt"))
	gamesSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameTotal := scanLine(line)
		gamesSum += gameTotal
	}
	fmt.Println(gamesSum)
}