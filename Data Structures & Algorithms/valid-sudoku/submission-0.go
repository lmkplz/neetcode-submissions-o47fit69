func isValidSudoku(board [][]byte) bool {
	size := len(board)

	rowHash := make([]map[byte]byte, size)
	colHash := make([]map[byte]byte, size)
	squareHash := make([]map[byte]byte, size)

	for i := 0; i < size; i++ {
		rowHash[i] = make(map[byte]byte)
		colHash[i] = make(map[byte]byte)
		squareHash[i] = make(map[byte]byte)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val := board[i][j]
			if val != '.' {
				rowHash[i][val]++
				colHash[j][val]++

				squareIndex := (i / 3) * 3 + j / 3
				squareHash[squareIndex][val]++

				if rowHash[i][val] > 1 || colHash[j][val] > 1 || squareHash[squareIndex][val] > 1 {
					return false
				}
			}
		}
	}

	return true
}
