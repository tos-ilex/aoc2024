package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := os.Open("input.txt")
	panicIfNotNil(err)
	defer input.Close()

	stat, err := input.Stat()
	panicIfNotNil(err)

	listSize := stat.Size() / 13 // input seems to be 13 chars per line so this should minimize resize ops
	leftList, rightList := make([]int, listSize), make([]int, listSize)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ { // i is the index of the first 'empty' list element in both lists
		splitLine := strings.Split(scanner.Text(), "   ")
		leftNum, err := strconv.Atoi(splitLine[0])
		panicIfNotNil(err)
		rightNum, err := strconv.Atoi(splitLine[1])
		panicIfNotNil(err)
		leftIndex := sort.Search(i+1, func(n int) bool { return leftList[n] >= leftNum })
		rightIndex := sort.Search(i+1, func(n int) bool { return rightList[n] >= rightNum })
		copy(leftList[leftIndex+1:], leftList[leftIndex:])
		copy(rightList[rightIndex+1:], rightList[rightIndex:])
		leftList[leftIndex] = leftNum
		rightList[rightIndex] = rightNum
	}

	distance := 0
	// the left and right lists now contain the entire input in a sorted manner
	for i := 0; i < len(leftList); i++ {
		difference := leftList[i] - rightList[i]
		distance += max(difference, -difference)
	}
	fmt.Printf("Result distance: %d\n", distance)
}
