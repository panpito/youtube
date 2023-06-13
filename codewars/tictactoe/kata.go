package tictactoe

import "log"

func IsSolved(grid [3][3]int) int {
	// build the board
	cells := make([]Cell, 0)
	byRow := make(map[int][]Cell)
	byCol := make(map[int][]Cell)

	for i, row := range grid {
		log.Print("row: ", row)
		for j, cellValue := range row {
			c := Cell{
				value: cellValue,
				x:     i,
				y:     j,
			}

			rowCells := byRow[i]
			if rowCells == nil {
				byRow[i] = []Cell{c}
			} else {
				byRow[i] = append(rowCells, c)
			}

			colCells := byCol[j]
			if colCells == nil {
				byCol[j] = []Cell{c}
			} else {
				byCol[j] = append(colCells, c)
			}

			cells = append(cells, c)
		}
	}

	board := Board{
		cells: cells,
		byRow: byRow,
		byCol: byCol,
		byDiag: map[int][]Cell{
			0: {
				{value: grid[0][0]},
				{value: grid[1][1]},
				{value: grid[2][2]},
			},
			1: {
				{value: grid[0][2]},
				{value: grid[1][1]},
				{value: grid[2][0]},
			},
		},
	}
	log.Print(board)

	// build sets of rows, columns and diagonals
	results := make([]Result, 0)

	for _, i := range []int{0, 1, 2} {
		rowResult := make(Result)
		for _, c := range board.byRow[i] {
			rowResult[c.value] = true
		}
		log.Printf("result for row %v, %v:\t\t %v, %v", i, board.byRow[i], rowResult, len(rowResult))
		results = append(results, rowResult)

		colResult := make(Result)
		for _, c := range board.byCol[i] {
			colResult[c.value] = true
		}
		log.Printf("result for col %v, %v:\t\t %v", i, board.byCol[i], colResult)
		results = append(results, colResult)
	}

	for _, i := range []int{0, 1} {
		diagResult := make(Result)
		for _, c := range board.byDiag[i] {
			diagResult[c.value] = true
		}
		log.Printf("result for dia %v, %v:\t\t %v", i, board.byDiag[i], diagResult)
		results = append(results, diagResult)
	}

	//map[int]bool{
	//	1: true,
	//}

	// determine the state of the board
	boardState := BoardState{}
	for _, result := range results {
		if len(result) == 1 {
			var key int
			for k := range result {
				key = k
			}
			log.Print(result, " ", key)

			switch key {
			case 1:
				boardState.xWon = true
			case 2:
				boardState.oWon = true
			case 0:
				boardState.emptySlot = true
			}
		}

		for key := range result {
			if key == 0 {
				boardState.emptySlot = true
			}
		}
	}

	log.Print(results)
	log.Print(boardState)
	// build the final response

	if boardState.xWon && boardState.oWon {
		return 0
	}

	if !boardState.xWon && !boardState.oWon && !boardState.emptySlot {
		return 0
	}

	if boardState.xWon {
		return 1
	}

	if boardState.oWon {
		return 2
	}

	return -1
}

type Board struct {
	cells  []Cell
	byRow  map[int][]Cell
	byCol  map[int][]Cell
	byDiag map[int][]Cell
}

type Cell struct {
	value int
	x, y  int
}

type Result map[int]bool

type BoardState struct {
	xWon, oWon, emptySlot bool
}
