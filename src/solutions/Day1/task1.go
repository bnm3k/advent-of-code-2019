package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	. "github.com/nagamocha3000/aoc2019/src/helpers"
)

func getModuleWeights(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var moduleWeights []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		moduleWeights = append(moduleWeights, num)
	}
	return moduleWeights, nil
}

func getFuel(moduleWeights []int) int {
	totalFuel := 0
	for _, weight := range moduleWeights {
		totalFuel += (weight / 3) - 2
	}

	return totalFuel
}

func main() {
	//read file contents
	file, err := os.Open("./input.txt")
	CheckErr(err)
	defer file.Close()

	//get module weights
	moduleWeights, err := getModuleWeights(file)
	CheckErr(err)

	//get fuel: 3150224
	totalFuel := getFuel(moduleWeights)
	fmt.Printf("Total fuel required: %d\n", totalFuel)
}
