package day02

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	utils "github.com/nagamocha3000/aoc2019/src/helpers"
)

//mutates intCode
func preprocessIntCode(intCode []int, noun, verb int) {
	intCode[1] = noun
	intCode[2] = verb
}

//mutates intCode
func evaluateIntCode(intCode []int) {
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
}

//creates copy of nums, does not mutate nums
func getResult(nums []int, noun, verb int) int {
	intCode := make([]int, len(nums))
	copy(intCode, nums)
	preprocessIntCode(intCode, noun, verb)
	evaluateIntCode(intCode)
	return intCode[0]
}

func getNounVerbCombo(nums []int, target int) int {
	noun, verb := 0, 0
outer:
	for ; noun <= 99; noun++ {
		for ; verb <= 99; verb++ {
			result := getResult(nums, noun, verb)
			if result == target {
				break outer
			}
		}
		verb = 0
	}
	return (100 * noun) + verb
}

func task1(nums []int) {
	result := getResult(nums, 12, 2)
	fmt.Printf("Task 1 result: %d\n", result)
}

func task2(nums []int) {
	result := getNounVerbCombo(nums, 19690720)
	fmt.Printf("Task 2 result: %d\n", result)
}

//Answer ...
func Answer() {
	//get path to input
	_, filename, _, _ := runtime.Caller(0)
	inputFilePath := filepath.Join(filepath.Dir(filename), "input.txt")

	//read input content
	content := utils.GetFileContentsAsStr(inputFilePath)
	content = strings.Trim(content, " \n\t")
	numStrs := strings.Split(content, ",")
	var nums []int
	for _, numStr := range numStrs {
		if num, err := strconv.Atoi(numStr); err == nil {
			nums = append(nums, num)
		}
	}

	//answer
	task1(nums)
	task2(nums)
}
