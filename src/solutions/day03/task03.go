package day03

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

//Dir
type towards int

const (
	up towards = iota
	down
	left
	right
)

type direction struct {
	towards
	steps int
}

var errInvalidDirStr = errors.New("Invalid direction string")

func parseDirection(singleDirStr string) (direction, error) {
	dir := direction{}
	num, err := strconv.Atoi(singleDirStr[1:])
	dir.steps = num
	switch singleDirStr[0] {
	case 'R':
		dir.towards = right
	case 'L':
		dir.towards = left
	case 'U':
		dir.towards = up
	case 'D':
		dir.towards = down
	default:
		err = errInvalidDirStr
	}
	return dir, err
}

func parseMultipleDirections(dirStr string) ([]direction, error) {
	var dirs []direction
	dirStrSlice := strings.Split(dirStr, ",")
	for _, dirStr := range dirStrSlice {
		dir, err := parseDirection(dirStr)
		if err != nil {
			return dirs, err
		}
		dirs = append(dirs, dir)
	}
	return dirs, nil
}

type circuitBoard struct {
	width  int
	height int
	board  [][]int
}

func newCircuitBoard(width, height int) circuitBoard {
	board := make([][]int, height)
	for i := 0; i < height; i++ {
		board[i] = make([]int, width)
	}
	cb := circuitBoard{width: width, height: height, board: board}
	return cb
}

type point struct {
	x int
	y int
}

var errInvalidPoint = errors.New("Invalid point")

func (cb *circuitBoard) markPoint(p point) error {
	if p.x < 0 || p.x > cb.width || p.y < 0 || p.y >= cb.height {
		return errInvalidPoint
	}
	cb.board[p.y][p.x] = 1
	return nil
}

func (cb *circuitBoard) markLine(startingPt point, dir direction) {
	currPt := startingPt
	for i := 0; i < dir.steps; i++ {
		nextPt := getAdjacentPoint(currPt, dir)
		err := cb.markPoint(nextPt)
		if err != nil {
			panic(err)
		}
		currPt = nextPt
	}
}

func getManhattanDistance(p1, p2 point) int {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func getNextPoint(currPt point, t towards, steps int) point {
	nextPt := currPt
	switch t {
	case up:
		nextPt.y -= steps
	case down:
		nextPt.y += steps
	case left:
		nextPt.x -= steps
	case right:
		nextPt.x += steps
	}
	return nextPt
}

func getEndPoint(startingPt point, dir direction) point {
	return getNextPoint(startingPt, dir.towards, dir.steps)
}

func getAdjacentPoint(startingPt point, dir direction) point {
	return getNextPoint(startingPt, dir.towards, 1)
}

//Answer ..
func Answer() {
	//get path to input
	_, filename, _, _ := runtime.Caller(0)
	inputFilePath := filepath.Join(filepath.Dir(filename), "input.txt")
	fmt.Println("Day 3", inputFilePath)
}
