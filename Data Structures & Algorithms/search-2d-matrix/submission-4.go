func searchMatrix(matrix [][]int, target int) bool {
	b, t := 0, len(matrix) - 1

	for t >= b {
		m := (t + b)/2
		if matrix[m][0] > target {
			t = m - 1
		} else if matrix[m][len(matrix[0])-1] < target {
			b = m + 1
		} else {
			break
		}
	}

	fmt.Println(t, b)
	if !(b <= t) {
		return false
	}

	row := (b+t)/2
	l, r := 0, len(matrix[0])-1


	for r >= l {
		fmt.Println(r, l)
		m := (r+l)/2
		if matrix[row][m] > target {
			r = m - 1
		} else if matrix[row][m] < target {
			l = m + 1
		} else {
			return true
		}
	}

	return false
}
