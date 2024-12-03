package day2

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPairSafe(left, right int, isIncreasing bool) bool {
	difference := left - right
	if isIncreasing && difference >= 0 {
		return false
	}
	if !isIncreasing && difference <= 0 {
		return false
	}
	if difference > 3 {
		return false
	}
	if difference < -3 {
		return false
	}
	return true
}

func removeElementAt(slice []int, i int) []int {
	result := make([]int, len(slice)-1)
	offset := 0
	for j, v := range slice {
		if j == i {
			offset = 1
			continue
		}
		result[j-offset] = v
	}
	return result
}

func isReportSafe(report string) bool {
	if len(report) <= 1 {
		return false
	}

	reportSplitBySpaces := strings.Split(report, " ")
	reportAsIntegers := make([]int, len(reportSplitBySpaces))

	for i, v := range reportSplitBySpaces {
		var err error
		reportAsIntegers[i], err = strconv.Atoi(v)
		utils.PanicIfNotNil(err)
	}

	for i := 0; i < len(reportAsIntegers); i++ {
		pissShit := removeElementAt(reportAsIntegers, i)

		fmt.Println("unsafe:")
		for _, v := range pissShit {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
		fmt.Println()

		if isIntReportSafe(pissShit) {
			return true
		}
	}
	return false
}

func isIntReportSafe(reportAsIntegers []int) bool {
	isIncreasing := reportAsIntegers[0] < reportAsIntegers[1]
	for i := 0; i < len(reportAsIntegers)-1; i++ {
		left, right := reportAsIntegers[i], reportAsIntegers[i+1]
		if !isPairSafe(left, right, isIncreasing) {
			return false
		}
	}

	fmt.Println("safe:")
	for _, v := range reportAsIntegers {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Println()

	return true

}

func PrintSolution(inputPath string) {
	inputFile, err := os.Open(inputPath)
	utils.PanicIfNotNil(err)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	safeReportCount := 0
	for scanner.Scan() {
		if isReportSafe(scanner.Text()) {
			safeReportCount++
		}
	}

	fmt.Println("Day 2")
	fmt.Printf("Number of safe reports: %d", safeReportCount)
}
