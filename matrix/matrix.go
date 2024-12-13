package matrix

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

type matrix [][]float64

var matrixIters int

func (x matrix) rowLen() int {
	return len(x)
}
func (x matrix) colLen() int {
	return len(x[0])
}

func Pow(matrix *mat.Dense, power int) (*mat.Dense, error) {
	r, c := matrix.Dims()
	if r != c {
		return nil, fmt.Errorf("matrix must be square")
	}
	identity := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		identity.Set(i, i, 1)
	}
	if power == 0 {
		return identity, nil
	}
	result := mat.DenseCopyOf(matrix)
	for i := 1; i < power; i++ {
		var temp mat.Dense
		temp.Mul(result, matrix)
		result.Copy(&temp)
	}
	return result, nil
}

func (m matrix) power(p int) matrix {
	if p < 1 {
		log.Fatal("cannot raise matrix to a power less than 1")
	}
	m_ := m
	i := 1
	for i < p {
		i++
		m_ = m_.matMul(m.transpose())
	}
	return m_
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
