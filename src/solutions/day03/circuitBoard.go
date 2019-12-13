package day03

import "fmt"

type circuitBoard struct {
	width  int
	height int
	board  [][]int
}

type markerFn func(int) int

func newCircuitBoard(width, height int) circuitBoard {
	board := make([][]int, height)
	for i := 0; i < height; i++ {
		board[i] = make([]int, width)
	}
	cb := circuitBoard{width: width, height: height, board: board}
	return cb
}

func (cb *circuitBoard) markPoint(p point, getMark markerFn) error {
	if p.x < 0 || p.x >= cb.width || p.y < 0 || p.y >= cb.height {
		return errInvalidPoint
	}
	currVal := cb.board[p.y][p.x]
	cb.board[p.y][p.x] = getMark(currVal)
	return nil
}

func (cb *circuitBoard) markLine(startingPt point, dir direction, getMark markerFn) {
	currPt := startingPt
	for i := 0; i < dir.steps; i++ {
		nextPt := getAdjacentPoint(currPt, dir)
		err := cb.markPoint(nextPt, getMark)
		if err != nil {
			panic(err)
		}
		currPt = nextPt
	}
}

func (cb *circuitBoard) markWire(centralPt point, wireDirs []direction, getMark markerFn) {
	startingPt := centralPt
	for _, dir := range wireDirs {
		cb.markLine(startingPt, dir, getMark)
		startingPt = getEndPoint(startingPt, dir)
	}
}

func (cb *circuitBoard) addWires(centralPt point, wire1Dirs, wire2Dirs []direction) error {
	wire1marker := func(i int) int {
		return 1
	}
	wire2marker := func(val int) int {
		if val == 1 || val == 3 { //intersection
			return 3
		}
		return 2 //is empty or wire2 already crossed but wire1 hadn't
	}
	cb.markWire(centralPt, wire1Dirs, wire1marker)
	cb.markWire(centralPt, wire2Dirs, wire2marker)
	return nil
}

func (cb *circuitBoard) getAllIntersections() []point {
	var intersections []point
	for y, row := range cb.board {
		for x, val := range row {
			if val == 3 {
				intersections = append(intersections, point{x: x, y: y})
			}
		}
	}
	return intersections
}

func (cb *circuitBoard) getClosestDistance(wire1Str, wire2Str string) (int, error) {
	wire1Dirs, err := parseMultipleDirections(wire1Str)
	if err != nil {
		return -1, err
	}
	wire2Dirs, err := parseMultipleDirections(wire2Str)
	if err != nil {
		return -1, err
	}
	centralPt := point{x: cb.width / 2, y: cb.height / 2}
	cb.addWires(centralPt, wire1Dirs, wire2Dirs)
	allIntersections := cb.getAllIntersections()
	closestPt := getClosestPoint(centralPt, allIntersections)
	dist := getManhattanDistance(centralPt, closestPt)
	fmt.Println(centralPt)
	return dist, nil
}

func (cb *circuitBoard) print() {
	for _, row := range cb.board {
		fmt.Println(row)
	}
}
