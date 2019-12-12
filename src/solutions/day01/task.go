package day01

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	utils "github.com/nagamocha3000/aoc2019/src/helpers"
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

func getTotalFuel(moduleWeights []int, getFuel func(int) int) int {
	totalFuel := 0
	for _, weight := range moduleWeights {
		totalFuel += getFuel(weight)
	}
	return totalFuel
}

func getFuelTask1(weight int) int {
	return (weight / 3) - 2
}

func getFuelTask2(weight int) int {
	fuelRequired := (weight / 3) - 2
	if fuelRequired < 0 {
		return 0
	}
	return fuelRequired + getFuelTask2(fuelRequired)
}

//Answer ...
func Answer() {
	//get path to input
	_, filename, _, _ := runtime.Caller(0)
	inputFilePath := filepath.Join(filepath.Dir(filename), "input.txt")

	//read file contents
	file, err := os.Open(inputFilePath)
	utils.CheckErr(err)
	defer file.Close()

	//get module weights
	moduleWeights, err := getModuleWeights(file)
	utils.CheckErr(err)

	//Task 1, get fuel: 3150224
	totalFuel := getTotalFuel(moduleWeights, getFuelTask1)
	fmt.Printf("Task1: Total fuel required= %d\n", totalFuel)

	//Task 2, get fuel + fuel's fuel: 4722484
	totalFuel = getTotalFuel(moduleWeights, getFuelTask2)
	fmt.Printf("Task2: Total fuel required= %d\n", totalFuel)
}
