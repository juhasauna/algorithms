package kolman

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test_relations(t *testing.T) {
	initTestData()

	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"boolMatrix_transitiveClosure", boolMatrix_transitiveClosureTest},
		// {"booleanMatrix_product", booleanMatrix_productTest},
		// {"relations_toMatrix", relations_toMatrixTest},
		// {"relations_isTransitive", relations_isTransitiveTest},
		// {"set_cartesianProduct", set_cartesianProductTest},
		{"relations_isSymmetric", relations_isSymmetricTest},
		// {"relations_isReflexive", relations_isReflexiveTest},
		// {"relations_range_", relations_range_Test},
		// {"newRelations", newRelationsTest},
		// {"isReflexive", isReflexiveTest},
		// {"isSymmetricMatrix", isSymmetricMatrixTest},
		// {"booleanMatrix_toRelations", booleanMatrix_toRelationsTest},
		// {"booleanMatrix_isPoset", booleanMatrix_isPosetTest},
		// {"boolMatrix_isEquivalenceRelation", boolMatrix_isEquivalenceRelationTest},
		// {"checkPropertyExercises", checkPropertyExercises},
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

func booleanMatrix_isPosetTest(t *testing.T) {
	tests := []struct {
		matrix boolMatrix
		expect bool
	}{
		{boolMatrixData.data["cp06.1_ex11"], true},
	}
	for i, tt := range tests {
		got := tt.matrix.isPoset()
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: with inputs %v, expected %t; got %t", i, tt.matrix, tt.expect, got)
		}
	}
}
func booleanMatrix_productTest(t *testing.T) {
	tests := []struct {
		matrix boolMatrix
		right  boolMatrix
		expect boolMatrix
	}{
		{boolMatrix{
			{0, 0},
			{0, 1},
		}, nil, boolMatrix{
			{0, 0},
			{0, 1}},
		},
		{boolMatrix{
			{0, 0},
			{1, 1},
		}, nil, boolMatrix{
			{0, 0},
			{1, 1}},
		},
		{boolMatrixData.data["cp01.5_ex13_A"], boolMatrixData.data["cp01.5_ex13_B"], boolMatrixData.data["cp01.5_ex13_product"]},
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
		{boolMatrixData.data["lessThan4x4"].toRelations(), false},
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
	cp04_04_matrix := func(key int) boolMatrix {
		x, ok := testRelationsData.cp04_05[key]
		if !ok {
			log.Fatalf("key: %d doesn't exists in map", key)
		}
		return x.toMatrix()
	}
	tests := []struct {
		bm     boolMatrix
		expect bool
	}{
		{testRelationsData.equivalence_4.toMatrix(), true},
		{boolMatrixData.test1, false},
		{cp04_04_matrix(5), false},
		{cp04_04_matrix(7), false},
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
		{boolMatrixData.test1},
	}
	for _, tt := range tests {
		got := tt.bm.toRelations()
		fmt.Println(got)
	}
}

func checkPropertyExercises(t *testing.T) {
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
		{in: boolMatrixData.data["diag3"],
			complement:   boolMatrix{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}},
			intersection: boolMatrix{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			inverse:      nil,
			product:      boolMatrixData.data["diag3"],
			transpose:    boolMatrixData.data["diag3"],
			union:        boolMatrixData.data["diag3"],
		},
		{in: boolMatrixData.data["diag3"],
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

func boolMatrix_transitiveClosureTest(t *testing.T) {
	tests := []struct {
		in     boolMatrix
		expect boolMatrix
	}{
		{boolMatrixData.get("cp04.8_examp02"), boolMatrixData.get("cp04.8_examp02_answer")},
		{boolMatrixData.get("cp04.8_examp03"), boolMatrixData.get("cp04.8_examp03_answer")},
	}
	for _, tt := range tests {
		got := tt.in.transitiveClosure(0)
		if !reflect.DeepEqual(tt.expect, got) {
			t.Errorf("FAIL: expected: %v, got: %v", tt.expect, got)
		}
	}
}
