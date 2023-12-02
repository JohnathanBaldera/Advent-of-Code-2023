package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type setInt struct {
	num byte
	valid bool
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func main() {	
	file, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	var total int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var first setInt
		var last setInt

		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				if first.valid {
					last.num = line[i]
					last.valid = true
				} else {
					first.num = line[i]
					first.valid = true
				}
			}			
		}

		// Edge Case: If there is only one number on the line
		if !last.valid {
			last.num = first.num
		}
	
		calibration, _ := strconv.Atoi(string(first.num) + string(last.num))

		total += calibration
	}
	
	fmt.Println(total)

}