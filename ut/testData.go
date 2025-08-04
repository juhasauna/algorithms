package ut

type TestData struct {
	Test2               []int
	Test3               []int
	Test4               []int
	Test5               []int
	Test6               []int
	Test7               []int
	Test8               []int
	Test9               []int
	Test10              []int
	Test11              []int
	Test12              []int
	TestDdata4813       []int
	TestNTUalg2022mid_6 []int
	// sorted []int
}

func (x *TestData) Init() {
	x.Test2 = []int{2, 1}
	x.Test3 = []int{3, 2, 1}
	x.Test4 = []int{4, 3, 2, 1}
	x.Test5 = []int{5, 4, 3, 2, 1}
	x.Test6 = []int{6, 5, 4, 3, 2, 1}
	x.Test7 = []int{7, 6, 4, 3, 2, 1}
	x.Test8 = []int{5, 4, 3, 2, 1, 0}
	x.Test9 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x.Test10 = []int{7, 3, 5, 15, 6, 1, 9, 2, 4, 12}
	x.Test11 = []int{7, 30, 5, 15, 6, 1, 9, 20, 4, 12, 16, 21, 31}
	x.Test12 = []int{7, 5, 15, 1, 9, 4, 12, 21, 31}
	x.TestDdata4813 = TestDdata4813
	x.TestNTUalg2022mid_6 = []int{8, 3, 2, 6, 5, 9, 10, 7, 1, 12, 4, 11}
}

var TDManberHeap = []int{6, 2, 8, 5, 10, 9, 12, 1, 15, 7, 3, 13, 4, 11, 16, 14}
var TDManberHeapIsMaxHeap = []int{16, 15, 13, 14, 10, 9, 12, 6, 5, 7, 3, 2, 4, 11, 8, 1}

var JuhaHeapBigger = []int{6, 2, 313, 8, 317, 318, 314, 315, 316, 5, 305, 306, 303, 304, 312, 311, 10, 9, 12, 301, 1, 15, 200, 7, 300, 301, 3, 13, 302, 4, 11, 16, 308, 307, 309, 14, 100, 101}
var NTU2024mid7 = []int{5, 2, 8, 4, 1, 15, 7, 6, 3, 11, 10, 12, 13, 14, 9}
