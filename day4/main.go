package day4

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
)

func notOutOfBounds(x, y int, input [][]byte) bool {
	return y >= 0 && x >= 0 && y < len(input) && x < len(input[y])
}

func checkWord(x, y int, input [][]byte, word string, traversalFunction func(int, int) (int, int)) bool {
	for i := 0; notOutOfBounds(x, y, input) && word[i] == input[y][x]; i++ {
		if input[y][x] == word[len(word)-1] {
			return true
		}
		x, y = traversalFunction(x, y)
	}
	return false
}

func checkMas(x, y int, input [][]byte) bool {
	if y+1 >= len(input) || x < 1 || y < 1 || x+1 >= len(input[y]) {
		return false
	}

	if input[y][x] != 'A' {
		return false
	}

	topleft := input[y-1][x-1]
	topright := input[y+1][x-1]
	bottomleft := input[y-1][x+1]
	bottomright := input[y+1][x+1]

	return (topleft == 'M' && bottomright == 'S' || topleft == 'S' && bottomright == 'M') &&
		(topright == 'M' && bottomleft == 'S' || topright == 'S' && bottomleft == 'M')
}

func PrintSolution(path string) {
	inputFile, err := os.Open(path)
	utils.PanicIfNotNil(err)

	var input [][]byte

	scanner := bufio.NewScanner(inputFile)
	for i := 0; scanner.Scan(); i++ {
		input = append(input, []byte(scanner.Text()))
	}

	wordCount := 0
	masCount := 0
	traversalFunctions := map[string]func(int, int) (int, int){
		"up":         func(x, y int) (int, int) { return x, y - 1 },
		"up-right":   func(x, y int) (int, int) { return x + 1, y - 1 },
		"right":      func(x, y int) (int, int) { return x + 1, y },
		"down-right": func(x, y int) (int, int) { return x + 1, y + 1 },
		"down":       func(x, y int) (int, int) { return x, y + 1 },
		"down-left":  func(x, y int) (int, int) { return x - 1, y + 1 },
		"left":       func(x, y int) (int, int) { return x - 1, y },
		"up-left":    func(x, y int) (int, int) { return x - 1, y - 1 },
	}

	for y := 0; y < len(input); y++ {
		// i have debugged this for too long to care about cleaning it up at this point :)
		for x := 0; x < len(input[y]); x++ {
			if checkMas(x, y, input) {
				masCount++
				fmt.Printf("Found mas at %d %d\n", x, y)
			}
			for _, traversalFunction := range traversalFunctions {
				if checkWord(x, y, input, "XMAS", traversalFunction) {
					wordCount++
				}
			}
		}
	}
	fmt.Printf("Day4\n* Solution: %d\n", wordCount)
	fmt.Printf("\n** Solution: %d\n", masCount)
}
