package day3

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
)

func indexOf(value byte, input []byte, start int) int {
	for i := start; i < len(input)-1; i++ {
		if input[i] == value {
			return i
		}
	}

	return -1
}

func getMulResult(openingBracketIndex int, input []byte) int {
	commaIndex := indexOf(',', input, openingBracketIndex)

	if commaIndex < 0 {
		return 0
	}

	leftNumber, err := strconv.Atoi(string(input[openingBracketIndex:commaIndex]))
	if err != nil {
		return 0
	}

	closingBracketIndex := indexOf(")"[0], input, commaIndex+1)
	if closingBracketIndex < 0 {
		return 0
	}

	rightNumber, err := strconv.Atoi(string(input[commaIndex+1 : closingBracketIndex]))
	if err != nil {
		return 0
	}

	return rightNumber * leftNumber
}

func PrintSolution(path string) {
	result := 0
	input, err := os.ReadFile(path)
	utils.PanicIfNotNil(err)

	enabled := true
	for i := 0; i+6 < len(input); i++ {
		if enabled && string(input[i:i+4]) == "mul(" {
			result += getMulResult(i+4, input)
		} else if string(input[i:i+4]) == "do()" {
			enabled = true
		} else if string(input[i:i+7]) == "don't()" {
			enabled = false
		}
	}

	fmt.Printf("\nDay3:\nResult: %d\n", result)
}
