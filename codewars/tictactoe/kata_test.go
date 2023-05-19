package tictactoe_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/panpito/youtube/codewars/tictactoe"
)

func doTest(board [3][3]int, expected int) {
	actual := tictactoe.IsSolved(board)
	Expect(actual).To(Equal(expected), "With board = %v\nExpected %d but got %d\n", board, expected, actual)
}

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		var board [3][3]int
		// not yet finished
		board = [3][3]int{
			{0, 0, 1},
			{0, 1, 2},
			{2, 1, 0},
		}
		doTest(board, -1)

		// winning row
		board = [3][3]int{
			{1, 1, 1},
			{0, 2, 2},
			{0, 0, 0},
		}
		doTest(board, 1)

		// winning column
		board = [3][3]int{
			{2, 1, 2},
			{2, 1, 1},
			{1, 1, 2},
		}
		doTest(board, 1)

		// draw
		board = [3][3]int{
			{2, 1, 2},
			{2, 1, 1},
			{1, 2, 1},
		}
		doTest(board, 0)
	})
})
