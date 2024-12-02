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

func countOccurrencesInSortedSlice(v int, slice []int) int {
	firstIndex := sort.Search(len(slice), func(i int) bool { return slice[i] >= v })
	if firstIndex == len(slice) || slice[firstIndex] != v || firstIndex == len(slice) {
		return 0
	}
	lastIndex := sort.Search(len(slice), func(i int) bool { return slice[i] > v })
	return lastIndex - firstIndex
}

func getSimilarity(left, right []int) int {
	result := 0
	for _, v := range left {
		occurrenceCount := countOccurrencesInSortedSlice(v, right)
		result += occurrenceCount * v
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
	firstEmptyIndex := 0
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "   ")
		leftNum := parseInt(splitLine[0])
		rightNum := parseInt(splitLine[1])
		leftIndex := searchIndexForSortedInsert(leftNum, leftList[:firstEmptyIndex])
		rightIndex := searchIndexForSortedInsert(rightNum, rightList[:firstEmptyIndex])
		insertAt(leftIndex, leftNum, leftList)
		insertAt(rightIndex, rightNum, rightList)
		firstEmptyIndex++
	}

	leftList = leftList[:firstEmptyIndex]
	rightList = rightList[:firstEmptyIndex]

	fmt.Printf("Result distance: %d\n", getDistance(leftList, rightList))
	fmt.Println(leftList)
	fmt.Printf("Result similarity: %d\n", getSimilarity(leftList, rightList))
}
