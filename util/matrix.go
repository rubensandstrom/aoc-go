package util

func MajorDiagonals[T any](matrix [][]T) [][]T {
	diagonals := [][]T{}
	for k := 0; k < len(matrix) + len(matrix[0]) - 1; k++ {
		diagonal := []T{}
		for i := 0; i < len(matrix); i++ {
			j := k - i
			if j >= 0 && j < len(matrix[0]) {
				diagonal = append(diagonal, matrix[i][j])
			}
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}
	return diagonals
}

func MinorDiagonals[T any](matrix [][]T) [][]T{
	diagonals := [][]T{}
	for k := 0; k < len(matrix) + len(matrix[0]) - 1; k++ {
		diagonal := []T{}
		for i := 0; i < len(matrix); i++ {
			j := k - (len(matrix) - 1 - i)
			if j >= 0 && j < len(matrix[0]) {
				diagonal = append(diagonal, matrix[i][j])
			}
		}
		if len(diagonal) > 0 {
			diagonals = append(diagonals, diagonal)
		}
	}
	return diagonals
}

func Cols[T any](matrix [][]T) [][]T{
	cols := [][]T{}
	for i := 0; i < len(matrix[0]); i++ {
		tmp := []T{}
		for j := 0; j < len(matrix); j++ {
			tmp = append(tmp, matrix[j][i])
		}
		cols = append(cols, tmp)
	}
	return cols
}

func Window[T any](matrix [][]T, x, y int) [][][]T {

	submatrixes := [][][]T{}
	for i := 0; i < len(matrix) - x + 1; i++ {
		for j := 0; j < len(matrix[0]) - y + 1; j++ {
			tmp := [][]T{}
			for k := i; k < i + x; k++ {
				tmp = append(tmp, matrix[k][j:j+y])
			}
			submatrixes = append(submatrixes, tmp)
		}
	}
	return submatrixes
}
