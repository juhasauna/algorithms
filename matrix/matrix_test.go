package matrix

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func Test_matrix(t *testing.T) {
	initTestData()

	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"transpose", transposeTest},
		// {"matMul", matMulTest},
		// {"power", powerTest},
		{"Pow", powTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func powTest(t *testing.T) {
	tests := []struct {
		m     *mat.Dense
		power int
	}{
		// {testMatrices.get("markov"), 1},
		// {testMatrices.get("markov"), 2},
		// {mat.NewDense(2, 2, []float64{.8, .1, .2, .9}), 100},
		// {mat.NewDense(2, 2, []float64{.8, .2, .5, .5}), 4},
		{mat.NewDense(2, 2, []float64{.6, .4, .8, .2}), 2},
		// {mat.NewDense(2, 2, []float64{.9, .6, .1, .4}), 100},
		// {mat.NewDense(3, 3, []float64{.8, .2, .3, .1, .7, .1, .1, .1, .6}), 100},
		// {mat.NewDense(3, 3, []float64{.8, .2, .1, .1, .6, .2, .1, .2, .7}), 100},
	}
	for _, tt := range tests {
		printMatrix(tt.m)
		got, err := Pow(tt.m, tt.power)
		if err != nil {
			t.Error(err)
		}
		printMatrix(got)
	}
}
func printMatrix(dense *mat.Dense) {
	s := ""
	r, _ := dense.Dims()
	for i := 0; i < r; i++ {
		vec := dense.RawRowView(i)
		s += fmt.Sprintf("%v\n", vec)
	}
	fmt.Println(s)
}
func powerTest(t *testing.T) {
	tests := []struct {
		m     matrix
		power int
	}{
		// {testMatrices.get("markov"), 1},
		// {testMatrices.get("markov"), 2},
		{testMatrices.get("markov"), 100},
	}
	for _, tt := range tests {
		got := tt.m.power(tt.power)
		t.Logf("%v", got)
	}
}
func matMulTest(t *testing.T) {
	tests := []struct {
		left   matrix
		right  matrix
		expect matrix
	}{
		// {testMatrices.get("diag_4"), testMatrices.get("diag_4"), testMatrices.get("diag_4")},
		// {testMatrices.get("triangular_lower_4"), testMatrices.get("triangular_upper_4"), testMatrices.get("triangular_matMul_4")},
		// {testMatrices.get("triangular_lower_4x3"), testMatrices.get("triangular_upper_3x4"), testMatrices.get("triangular_4x3_matMul_3x4")},
		// {testMatrices.get("a_3x3"), testMatrices.get("a_3x2"), testMatrices.get("a_3x3*a_3x2")},
		{testMatrices.get("a_2x4"), testMatrices.get("a_2x4.T"), testMatrices.get("a_2x4_matMult")},
	}
	for i, tt := range tests {
		matrixIters = 0
		got := tt.left.matMul(tt.right)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: iters: %d, inputs %v %v, expected %v; got %v\n", i, matrixIters, tt.left, tt.right, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS %d: iters: %d\n", i, matrixIters)
		}
	}
}
func transposeTest(t *testing.T) {
	tests := []struct {
		m      matrix
		expect matrix
	}{
		{testMatrices.get("diag_4"), testMatrices.get("diag_4")},
		{testMatrices.get("triangular_lower_4"), testMatrices.get("triangular_upper_4")},
		{testMatrices.get("triangular_lower_4x3"), testMatrices.get("triangular_upper_3x4")},
		{testMatrices.get("triangular_upper_3x4"), testMatrices.get("triangular_lower_4x3")},
		{testMatrices.get("a_2x4"), testMatrices.get("a_2x4.T")},
		{testMatrices.get("a_2x4.T"), testMatrices.get("a_2x4")},
	}
	for i, tt := range tests {
		got := tt.m.transpose()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %v; got %v\n", i, tt.m, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS %d\n", i)
		}
	}
}

var testMatrices matrixTestData

type matrixTestData struct {
	data map[string]matrix
}

func (x matrixTestData) get(key string) matrix {
	item, ok := x.data[key]
	if !ok {
		log.Fatalf("key: %s doesn't exists in matrixTestData map", key)
	}
	return item
}

func initTestData() {
	x := matrixTestData{
		data: make(map[string]matrix),
	}
	x.data["markov"] = matrix{
		{.8, .1},
		{.2, .9},
	}
	x.data["diag_4"] = matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	x.data["triangular_lower_4"] = matrix{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
		{1, 1, 1, 1},
	}
	x.data["triangular_upper_4"] = matrix{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
		{0, 0, 0, 1},
	}
	x.data["triangular_matMul_4"] = matrix{
		{1, 1, 1, 1},
		{1, 2, 2, 2},
		{1, 2, 3, 3},
		{1, 2, 3, 4},
	}
	x.data["triangular_lower_4x3"] = matrix{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}
	x.data["triangular_upper_3x4"] = matrix{
		{1, 1, 1},
		{0, 1, 1},
		{0, 0, 1},
		{0, 0, 0},
	}
	x.data["triangular_4x3_matMul_3x4"] = matrix{
		{1, 1, 1},
		{1, 2, 2},
		{1, 2, 3},
	}
	x.data["triangular_4x3_matMul_3x4"] = matrix{
		{1, 1, 1},
		{1, 2, 2},
		{1, 2, 3},
	}
	x.data["a_3x3"] = matrix{
		{1, 2, 1},
		{0, 1, 0},
		{2, 3, 4},
	}
	x.data["a_3x2"] = matrix{
		{2, 5},
		{6, 7},
		{1, 8},
	}
	x.data["a_3x3*a_3x2"] = matrix{
		{15, 27},
		{6, 7},
		{26, 63},
	}
	x.data["a_2x4"] = matrix{
		{1, 2, 3, 44},
		{1, 2, 33, 4},
	}
	x.data["a_2x4.T"] = matrix{
		{1, 1},
		{2, 2},
		{3, 33},
		{44, 4},
	}
	x.data["a_2x4_matMult"] = matrix{
		{1950, 280},
		{280, 1110},
	}
	testMatrices = x
}
