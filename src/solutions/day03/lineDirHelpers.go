package day03

import (
	"errors"
	"strconv"
	"strings"
)

//point
type point struct {
	x int
	y int
}

var errInvalidPoint = errors.New("Invalid point")

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

func getManhattanDistance(p1, p2 point) int {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func getClosestPoint(startingPt point, points []point) point {
	minDist := getManhattanDistance(startingPt, points[0])
	closestPt := points[0]
	for _, pt := range points {
		if getManhattanDistance(startingPt, pt) < minDist {
			closestPt = pt
		}
	}
	return closestPt
}
