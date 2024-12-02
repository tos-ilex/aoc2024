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

func parseInt(s string) int {
	result, err := strconv.Atoi(s)
	panicIfNotNil(err)
	return result
}

func searchIndexForSortedInsert(numToInsert int, sliceToInsertIn []int) int {
	return sort.Search(len(sliceToInsertIn), func(i int) bool { return sliceToInsertIn[i] >= numToInsert })
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
		leftNum, rightNum := parseInt(splitLine[0]), parseInt(splitLine[1])

		// this will do a binary search to find the index of the first element that's >= the number to insert
		leftIndex, rightIndex := searchIndexForSortedInsert(leftNum, leftList[:i+1]), searchIndexForSortedInsert(rightNum, rightList[:i+1])

		// shift values at index one to the right to make room
		copy(leftList[leftIndex+1:], leftList[leftIndex:])
		copy(rightList[rightIndex+1:], rightList[rightIndex:])

		// then insert
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
