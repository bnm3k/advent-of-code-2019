package day03

import (
	"fmt"
	"path/filepath"
	"runtime"
)

//Answer ..
func Answer() {
	//get path to input
	_, filename, _, _ := runtime.Caller(0)
	inputFilePath := filepath.Join(filepath.Dir(filename), "input.txt")
	fmt.Println("Day 3", inputFilePath)
}
