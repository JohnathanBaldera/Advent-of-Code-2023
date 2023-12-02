package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type setInt struct {
	num byte
	valid bool
	index int
}

type wordInt struct {
	name string
	value byte
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}


func firstNumberIndex(line string, numbers [10]wordInt) (wordInt, int) {
	var firstNum wordInt
	index := 9999
		
	for i := 0; i < len(numbers); i++ {
		numIndex := strings.Index(line, numbers[i].name)
		if numIndex != -1 && numIndex < index {
			index = numIndex
			firstNum = numbers[i]
		}
	}

	return firstNum, index
}

func lastNumberIndex(line string, numbers [10]wordInt) (wordInt, int) {
	var lastNum wordInt
	index := -1

	for i := 0; i < len(numbers); i++ {
		numIndex := strings.LastIndex(line, numbers[i].name)
		if numIndex > index {
			index = numIndex
			lastNum = numbers[i]
		}
	}

	return lastNum, index
}

func main() {	
	numbers := [10]wordInt{{"zero", 48}, {"one", 49}, {"two", 50}, {"three", 51}, {"four", 52}, {"five", 53}, {"six", 54}, {"seven", 55}, {"eight", 56}, {"nine", 57}}
	file, err := os.Open("../input.txt")
	
	if err != nil {
		panic(err)
	}
	
	
	var total int
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		
		var firstInt setInt
		var lastInt setInt
		
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				if firstInt.valid {
					lastInt.num = line[i]
					lastInt.valid = true
					lastInt.index = i
				} else {
					firstInt.num = line[i]
					firstInt.valid = true
					firstInt.index = i
				}
			}
		}	

		// Edge Case: If there is only one number on the line
		if !lastInt.valid {
			lastInt.num = firstInt.num
			lastInt.index = firstInt.index
			lastInt.valid = true
		}

		firstWord, firstWordIndex := firstNumberIndex(line, numbers)
		lastWord, lastWordIndex := lastNumberIndex(line, numbers)

		if firstInt.index > firstWordIndex && firstWordIndex != 9999 {
			firstInt.num = byte(firstWord.value)
		}

		if lastInt.index < lastWordIndex && lastWordIndex != -1 {
			lastInt.num = byte(lastWord.value)
			lastInt.valid = true
		}

	
		calibration, _ := strconv.Atoi(string(firstInt.num) + string(lastInt.num))
		total += calibration
	}
	
	fmt.Println(total)

}