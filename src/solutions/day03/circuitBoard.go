package day03

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
