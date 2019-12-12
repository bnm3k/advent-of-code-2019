package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	utils "github.com/nagamocha3000/aoc2019/src/helpers"
)

func preprocessIntCode(nums []int) []int {
	intCode := make([]int, len(nums))
	copy(intCode, nums)
	intCode[1] = 12
	intCode[2] = 2
	return intCode
}

func evaluateIntCode(nums []int) []int {
	intCode := make([]int, len(nums))
	copy(intCode, nums)
	for currPt := 0; ; {
		if intCode[currPt] == 99 {
			break
		}

		firstIndex, secondIndex, resultIndex := intCode[currPt+1], intCode[currPt+2], intCode[currPt+3]
		firstVal, secondVal := intCode[firstIndex], intCode[secondIndex]

		switch intCode[currPt] {
		case 1:
			result := firstVal + secondVal
			intCode[resultIndex] = result
			currPt += 4
		case 2:
			result := firstVal * secondVal
			intCode[resultIndex] = result
			currPt += 4
		}
	}
	return intCode
}

func getResult(intCode []int) int {
	return intCode[0]
}

func getFileContents(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	utils.CheckErr(err)
	return string(content)
}

func main() {
	content := getFileContents("input.txt")
	content = strings.Trim(content, " \n\t")
	numStrs := strings.Split(content, ",")
	var nums []int
	for _, numStr := range numStrs {
		if num, err := strconv.Atoi(numStr); err == nil {
			nums = append(nums, num)
		}
	}
	preprocessedIntCode := preprocessIntCode(nums)
	evaluatedIntCode := evaluateIntCode(preprocessedIntCode)
	result := getResult(evaluatedIntCode)
	fmt.Printf("Task 1 result: %d", result)
}
