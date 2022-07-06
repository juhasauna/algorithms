package matrix

import "log"

type matrix [][]float64

var matrixIters int

func (x matrix) rowLen() int {
	return len(x)
}
func (x matrix) colLen() int {
	return len(x[0])
}

func (left matrix) matMul(right matrix) matrix {
	// This does not work :/
	leftRowLen, leftColLen := left.rowLen(), left.colLen()
	rightRowLen, rightColLen := right.rowLen(), right.colLen()
	if leftColLen != rightRowLen {
		log.Fatalf("unable to multiply given matrices, leftColLen (%d) should equal rightRowLen (%d)", leftColLen, rightRowLen)
	}
	result := initMatrix(leftRowLen, rightColLen)
	for i, leftRow := range left {
		for j := 0; j < rightColLen; j++ {
			var productSum float64
			for rightRowIndex := range right {
				productSum += leftRow[rightRowIndex] * right[rightRowIndex][j]
				matrixIters++
			}
			result[i][j] = productSum
		}
	}
	return result
}

func (x matrix) transpose() matrix {
	result := initMatrix(x.colLen(), x.rowLen())
	rLen := x.rowLen()
	cLen := x.colLen()
	lessRows := rLen < cLen
	for r := 0; r < cLen; r++ {
		for c := 0; c < rLen; c++ {
			insertValue := func() float64 {
				if lessRows {
					return x[c%cLen][r]
				}
				return x[c][r%rLen]
			}()
			result[r][c] = insertValue
		}
	}

	return result
}

func initMatrix(rows, cols int) (result matrix) {
	for i := 0; i < rows; i++ {
		result = append(result, []float64{})
		for j := 0; j < cols; j++ {
			result[i] = append(result[i], 0)
		}
	}
	return result
}
