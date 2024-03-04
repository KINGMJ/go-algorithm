package n_queens

import (
	"strings"
)

func SolveNQueens(n int) [][]string {
	// 最终结果
	var res [][]string
	// 递归调用函数
	var callNQueens func(board []string, row int)
	// 初始化棋盘
	board := make([]string, n)
	for i := range board {
		board[i] = strings.Repeat(".", n)
	}

	callNQueens = func(board []string, row int) {
		// 递归结束条件
		if row == n {
			tmp := make([]string, n)
			copy(tmp, board)
			res = append(res, tmp)
			return
		}
		// 选择路径
		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			// 做选择
			board[row] = board[row][:col] + "Q" + board[row][col+1:]
			// 递归进入下一行选择
			callNQueens(board, row+1)
			// 撤销选择，恢复棋盘
			board[row] = strings.Repeat(".", n)
		}
	}

	callNQueens(board, 0)
	return res
}

func isValid(board []string, row, col int) bool {
	// 该棋子的左上角和右上角
	leftup := col - 1
	rightup := col + 1

	// 逐行向上检查是否可以放置棋子
	for i := row - 1; i >= 0; i-- {
		// 正上方不能放置
		if board[i][col] == 'Q' {
			return false
		}
		// 左对角线
		if leftup >= 0 && board[i][leftup] == 'Q' {
			return false
		}
		// 右对角线
		if rightup < len(board) && board[i][rightup] == 'Q' {
			return false
		}
		// 每一行对角线位置更新
		leftup--
		rightup++
	}
	return true
}
