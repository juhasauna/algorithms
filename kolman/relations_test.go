package kolman

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var testRelationsData relationsTestData
var testBoolMatrixData booleanMatrixTestData

func Test_relations(t *testing.T) {
	initTestData()

	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"booleanMatrix_product", booleanMatrix_productTest},
		// {"relations_toMatrix", relations_toMatrixTest},
		// {"relations_isTransitive", relations_isTransitiveTest},
		// {"set_cartesianProduct", set_cartesianProductTest},
		// {"relations_isSymmetric", relations_isSymmetricTest},
		// {"relations_isReflexive", relations_isReflexiveTest},
		// {"relations_range_", relations_range_Test},
		// {"newRelations", newRelationsTest},
		// {"isReflexive", isReflexiveTest},
		// {"isSymmetricMatrix", isSymmetricMatrixTest},
		// {"newRelationsFromMatrix", newRelationsFromMatrixTest},
		// {"booleanMatrix_toRelations", booleanMatrix_toRelationsTest},
		// {"boolMatrix_isEquivalenceRelation", boolMatrix_isEquivalenceRelationTest},
		{"checkExercises", checkExercises},
		// {"integerRelations", integerRelationsTest},
		// {"boolMatrixOperations", boolMatrixOperationsTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func relations_isTransitiveTest(t *testing.T) {
	tests := []struct {
		relations relations
		expect    bool
	}{
		// {testRelationsData.trans_sym_3, true},
		// {testRelationsData.cp04_04[2], true},
		{testRelationsData.cp04_04[8], false},
		{testRelationsData.cp04_04[108], false},
	}
	for i, tt := range tests {
		fmt.Println(tt.relations.getProperties())
		got := tt.relations.isTransitive()
		if !got == tt.expect {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %v\n", i, tt.relations, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS %d\n", i)
		}
	}
}
func relations_isTransitiveTest_old(t *testing.T) {
	tests := []struct {
		relations relations
		expect    bool
	}{
		{testRelationsData.trans_sym_3, true},
		{testRelationsData.cp04_04[2], true},
		// {testRelationsData.cp04_04[8], false},
	}
	for i, tt := range tests {
		got := tt.relations.isTransitive_old()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %v\n", i, tt.relations, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS %d\n", i)
		}
	}
}
func isSymmetricMatrixTest(t *testing.T) {
	tests := []struct {
		matrix [][]int
		expect bool
	}{
		{[][]int{{0, 0}, {0, 0}}, true},
		{[][]int{{1, 1}, {1, 1}}, true},
		{[][]int{{0, 1}, {1, 1}}, true},
		{[][]int{{1, 0}, {1, 1}}, false},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, true},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 1}}, false},
	}
	for i, tt := range tests {
		got := isSymmetricMatrix(tt.matrix)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %v", i, tt.matrix, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}
func isReflexiveTest(t *testing.T) {
	tests := []struct {
		relation []relation
		expect   bool
	}{
		{[]relation{{"a", "b"}}, false},
		{[]relation{{"a", "a"}}, true},
		{[]relation{{"a", "b"}, {"b", "a"}, {"a", "a"}}, false},
		{[]relation{{"a", "b"}, {"b", "a"}, {"a", "a"}, {"b", "b"}, {"c", "b"}}, false},
		{[]relation{{"a", "b"}, {"b", "a"}, {"c", "a"}}, false},
	}
	for i, tt := range tests {
		got := isReflexive(tt.relation)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %v", i, tt.relation, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}

func booleanMatrix_productTest(t *testing.T) {
	tests := []struct {
		matrix boolMatrix
		right  boolMatrix
		expect boolMatrix
	}{
		// {boolMatrix{
		// 	{0, 0},
		// 	{0, 1},
		// }, nil, boolMatrix{
		// 	{0, 0},
		// 	{0, 1}},
		// },
		// {boolMatrix{
		// 	{0, 0},
		// 	{1, 1},
		// }, nil, boolMatrix{
		// 	{0, 0},
		// 	{1, 1}},
		// },
		{boolMatrix{
			{1, 1, 0},
			{0, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		}, boolMatrix{
			{1, 0, 0, 0},
			{0, 1, 1, 0},
			{1, 0, 1, 1},
		}, boolMatrix{
			{1, 1, 1, 0},
			{0, 1, 1, 0},
			{1, 1, 1, 0},
			{1, 0, 1, 1},
		}},
	}
	for i, tt := range tests {
		right := tt.right
		if right == nil {
			right = tt.matrix
		}
		got := tt.matrix.product(right)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %v; got %v", i, tt.matrix, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}
func relations_toMatrixTest(t *testing.T) {
	tests := []struct {
		relations relations
		expect    boolMatrix
	}{
		{newRelations([]relation{{"a", "b"}, {"c", "d"}}), boolMatrix{
			{0, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 0, 0}}},
		{newRelations([]relation{{"a", "a"}, {"a", "b"}, {"c", "c"}, {"c", "d"}}), boolMatrix{
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 0}}},
	}
	for i, tt := range tests {
		got := tt.relations.toMatrix()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %v; got %v", i, tt.relations.relationSet, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}
func relations_isReflexiveTest(t *testing.T) {
	tests := []struct {
		relations relations
		expect    bool
	}{
		{newRelations([]relation{{"a", "b"}, {"c", "d"}}), false},
		{newRelations([]relation{{"a", "a"}, {"a", "b"}, {"b", "b"}, {"c", "c"}}), true},
	}
	for i, tt := range tests {
		got := tt.relations.isReflexive()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %v", i, tt.relations.relationSet, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}
func relations_isSymmetricTest(t *testing.T) {
	tests := []struct {
		relations relations
		expect    bool
	}{
		{newRelations([]relation{{"a", "b"}, {"c", "d"}}), false},
		{newRelations([]relation{{"a", "b"}, {"b", "a"}, {"c", "d"}, {"d", "c"}, {"a", "a"}}), true},
	}
	for i, tt := range tests {
		got := tt.relations.isSymmetric()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %v", i, tt.relations.relationSet, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}
func relations_range_Test(t *testing.T) {
	tests := []struct {
		relations relations
		expect    set
	}{
		{newRelations([]relation{{"a", "b"}, {"c", "d"}}), newSet("b", "d")},
		{testRelationsData.test1, newSet("a", "b", "c", "d")},
	}
	for i, tt := range tests {
		got := tt.relations.range_()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %v; got %v", i, tt.relations.relationSet, tt.expect, got)
		} else {
			fmt.Printf("SUCCESS: %d\n", i)
		}
	}
}

func newRelationsTest(t *testing.T) {
	tests := []struct {
		x relations
	}{
		{newRelations([]relation{{"a", "b"}, {"c", "d"}})},
	}
	for _, tt := range tests {
		fmt.Println(tt.x)
	}
}
func boolMatrix_isEquivalenceRelationTest(t *testing.T) {
	tests := []struct {
		bm     boolMatrix
		expect bool
	}{
		{testRelationsData.equivalence_4.toMatrix(), true},
		{testBoolMatrixData.test1, false},
	}
	for i, tt := range tests {
		got := tt.bm.isEquivalenceRelation()
		if got != tt.expect {
			t.Errorf("FAIL %d, inputs %v, expected: %t, got: %t", i, tt.bm, tt.expect, got)
		}
	}
}
func booleanMatrix_toRelationsTest(t *testing.T) {
	tests := []struct {
		bm boolMatrix
	}{
		{testBoolMatrixData.test1},
	}
	for _, tt := range tests {
		got := tt.bm.toRelations()
		fmt.Println(got)
	}
}

type booleanMatrixTestData struct {
	test1      boolMatrix
	diagonal_3 boolMatrix
	cp04_04    map[int]boolMatrix
}

func initBooleanMatrixTestData() {
	testBoolMatrixData = booleanMatrixTestData{
		test1: boolMatrix{
			{1, 1, 1, 0},
			{0, 1, 1, 0},
			{1, 1, 1, 0},
			{1, 0, 1, 1},
		},
		diagonal_3: boolMatrix{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
	testBoolMatrixData.cp04_04 = make(map[int]boolMatrix)
	testBoolMatrixData.cp04_04[11] = boolMatrix{
		{0, 1, 0, 1},
		{1, 0, 1, 1},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	testBoolMatrixData.cp04_04[12] = boolMatrix{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	testBoolMatrixData.cp04_04[13] = boolMatrix{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
	}
	testBoolMatrixData.cp04_04[3701] = boolMatrix{ // Ex.37 a) reflexive and symmetric, but not transitive.
		{1, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	testBoolMatrixData.cp04_04[3702] = boolMatrix{ // Ex.37 b) reflexive and transitive, but not symmetric.
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	testBoolMatrixData.cp04_04[3801] = boolMatrix{ // Ex.38 a) irreflexive and transitive, but not symmetric.
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	testBoolMatrixData.cp04_04[3802] = boolMatrix{ // Ex.38 b) antisymmetric and reflexive, but not transitive.
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	testBoolMatrixData.cp04_04[3901] = boolMatrix{ // Ex.39 a) transitive, reflexive, and symmetric.
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	testBoolMatrixData.cp04_04[3902] = boolMatrix{ // Ex.39 b) asymmetric and transitive.
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

}
func initTestData() {
	initBooleanMatrixTestData()
	initRelationsTestData()
}

type relationsTestData struct {
	test1         relations
	trans_sym_3   relations
	equivalence_4 relations
	cp04_04       map[int]relations
	integerFuncs  map[string]func(int, int) bool
}

func initRelationsTestData() {
	setA := newSet("1", "2", "3", "4")
	testRelationsData = relationsTestData{
		test1:         newRelations([]relation{{"a", "a"}, {"a", "b"}, {"c", "c"}, {"c", "d"}}),
		trans_sym_3:   newRelations([]relation{{"a", "a"}, {"a", "b"}, {"a", "c"}, {"b", "c"}, {"c", "c"}}),
		equivalence_4: newRelations([]relation{{"1", "1"}, {"1", "2"}, {"2", "1"}, {"2", "2"}, {"3", "4"}, {"4", "3"}, {"3", "3"}, {"4", "4"}}),
	}
	testRelationsData.cp04_04 = make(map[int]relations)
	testRelationsData.integerFuncs = make(map[string]func(int, int) bool)
	testRelationsData.cp04_04[1] = newRelations([]relation{{"1", "1"}, {"1", "2"}, {"2", "1"}, {"2", "2"}, {"3", "3"}, {"3", "4"}, {"4", "3"}, {"4", "4"}})
	testRelationsData.cp04_04[2] = newRelations([]relation{{"1", "2"}, {"1", "3"}, {"1", "4"}, {"2", "3"}, {"2", "4"}, {"3", "4"}})
	testRelationsData.cp04_04[6] = setA.cartesianSquare()
	testRelationsData.cp04_04[8] = newRelations([]relation{{"1", "3"}, {"4", "2"}, {"2", "4"}, {"3", "1"}, {"2", "2"}})
	testRelationsData.cp04_04[108] = newRelations([]relation{{"4", "2"}, {"2", "4"}})
	testRelationsData.cp04_04[10] = newRelations([]relation{{"1", "2"}, {"1", "3"}, {"1", "4"}, {"5", "2"}, {"5", "3"}, {"5", "4"}})
	testRelationsData.cp04_04[11] = testBoolMatrixData.cp04_04[11].toRelations()
	testRelationsData.cp04_04[12] = testBoolMatrixData.cp04_04[12].toRelations()
	testRelationsData.cp04_04[13] = testBoolMatrixData.cp04_04[13].toRelations()
	testRelationsData.cp04_04[3701] = testBoolMatrixData.cp04_04[3701].toRelations()
	testRelationsData.cp04_04[3702] = testBoolMatrixData.cp04_04[3702].toRelations()
	testRelationsData.cp04_04[3801] = testBoolMatrixData.cp04_04[3801].toRelations()
	testRelationsData.cp04_04[3802] = testBoolMatrixData.cp04_04[3802].toRelations()
	testRelationsData.cp04_04[3901] = testBoolMatrixData.cp04_04[3901].toRelations()
	testRelationsData.cp04_04[3902] = testBoolMatrixData.cp04_04[3902].toRelations()
	testRelationsData.integerFuncs["cp04_04_13"] = func(from, to int) bool { return from <= (to + 1) }

}

func checkExercises(t *testing.T) {
	tests := []struct {
		key        int
		properties properties
	}{
		// Ref,Irrref,Sym,Asym,AntiSym,Trans
		{1, properties{1, 0, 1, 0, 0, 1}},
		{2, properties{0, 1, 0, 1, 1, 1}},
		{6, properties{1, 0, 1, 0, 0, 1}},
		{8, properties{0, 0, 1, 0, 0, 0}},
		{10, properties{0, 1, 0, 1, 1, 1}},
		{11, properties{0, 1, 1, 0, 0, 0}},
		{12, properties{1, 0, 1, 0, 0, 1}},
		{13, properties{1, 0, 0, 0, 0, 0}},
		{3701, properties{1, 0, 1, 0, 0, 0}},
		{3702, properties{1, 0, 0, 0, 1, 1}},
		{3801, properties{0, 1, 0, 1, 1, 1}},
		{3802, properties{1, 0, 0, 0, 1, 0}},
		{3901, properties{1, 0, 1, 0, 0, 1}},
		{3902, properties{0, 1, 0, 1, 1, 1}},
	}
	for _, tt := range tests {
		relations, ok := testRelationsData.cp04_04[tt.key]
		if !ok {
			log.Fatalf("key: %d doesn't exists in map", tt.key)
		}
		got := relations.getProperties()
		if !reflect.DeepEqual(got, tt.properties) {
			t.Errorf("FAIL %d: expected %v, got: %v", tt.key, tt.properties, got)
		}
	}
}
func set_cartesianProductTest(t *testing.T) {
	tests := []struct {
		a      set // domain
		b      set // range
		expect relations
	}{
		// {newSet("a", "b", "c"), newSet("1", "2"), newSet("a;1", "a;2", "b;1", "b;2", "c;1", "c;2")},
		{newSet("a", "b", "c"), newSet("1", "2"), newRelations([]relation{{"a", "1"}, {"a", "2"}, {"b", "1"}, {"b", "2"}, {"c", "1"}, {"c", "2"}})},
		{newSet("a", "b", "c"), nil, newRelations([]relation{{"a", "a"}, {"a", "b"}, {"a", "c"}, {"b", "a"}, {"b", "b"}, {"b", "c"}, {"c", "a"}, {"c", "b"}, {"c", "c"}})},
	}
	for i, tt := range tests {
		got := new(relations)
		if tt.b != nil {
			*got = tt.a.cartesianProduct(tt.b)
		} else {
			*got = tt.a.cartesianSquare()
		}
		if !reflect.DeepEqual(*got, tt.expect) {
			t.Errorf("FAIL %d: expected %v, got: %v", i, tt.expect, got)
		}
	}
}
func integerRelationsTest(t *testing.T) {
	tests := []struct {
		from   int
		to     int
		f      func(int, int) bool
		expect relations
	}{
		{0, 4, testRelationsData.integerFuncs["cp04_04_13"], testRelationsData.cp04_04[13]},
	}
	for i, tt := range tests {
		got := integerRelations(tt.from, tt.to, tt.f)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: expected %v, got: %v", i, tt.expect, got)
		}
	}
}

func boolMatrixOperationsTest(t *testing.T) {
	tests := []struct {
		in           boolMatrix
		complement   boolMatrix
		intersection boolMatrix
		inverse      boolMatrix
		product      boolMatrix
		transpose    boolMatrix
		union        boolMatrix
	}{
		{in: testBoolMatrixData.diagonal_3,
			complement:   boolMatrix{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}},
			intersection: boolMatrix{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			inverse:      nil,
			product:      testBoolMatrixData.diagonal_3,
			transpose:    testBoolMatrixData.diagonal_3,
			union:        testBoolMatrixData.diagonal_3,
		},
		{in: testBoolMatrixData.diagonal_3,
			complement:   nil,
			intersection: nil,
			inverse:      nil,
			product:      nil,
			transpose:    nil,
			union:        nil,
		},
	}
	for _, tt := range tests {
		gotComplement := tt.in.complement()
		if tt.complement != nil {
			if !reflect.DeepEqual(tt.complement, gotComplement) {
				t.Errorf("FAIL complement, input: %v\n, expect: %v,\n got: %v\n", tt.in, tt.complement, gotComplement)
			}
		}
		if tt.intersection != nil {
			got := tt.in.intersection(gotComplement)
			if !reflect.DeepEqual(tt.intersection, got) {
				t.Errorf("FAIL intersection, input: %v\n, expect: %v,\n got: %v\n", tt.in, tt.intersection, got)
			}
		}
		// TODO: inverse
		if tt.product != nil {
			got := tt.in.product(tt.in)
			if !reflect.DeepEqual(tt.product, got) {
				t.Errorf("FAIL intersection, input: %v\n, expect: %v,\n got: %v\n", tt.in, tt.intersection, got)
			}
		}
	}
}
