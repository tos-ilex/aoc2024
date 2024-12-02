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

func insertAt(index int, value int, sliceToInsertIn []int) {
	copy(sliceToInsertIn[index+1:], sliceToInsertIn[index:])
	sliceToInsertIn[index] = value
}

func getDistance(left, right []int) int {
	result := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		result += max(distance, -distance)
	}
	return result
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
	for firstEmptyIndex := 0; scanner.Scan(); firstEmptyIndex++ {
		splitLine := strings.Split(scanner.Text(), "   ")
		leftNum := parseInt(splitLine[0])
		rightNum := parseInt(splitLine[1])
		leftIndex := searchIndexForSortedInsert(leftNum, leftList[:firstEmptyIndex+1])
		rightIndex := searchIndexForSortedInsert(rightNum, rightList[:firstEmptyIndex+1])
		insertAt(leftIndex, leftNum, leftList)
		insertAt(rightIndex, rightNum, rightList)
	}

	fmt.Printf("Result distance: %d\n", getDistance(leftList, rightList))
}
