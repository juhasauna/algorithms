package kolman

import (
	"fmt"
	"testing"
)

var treesMatrixTestData map[string]treeMatrix
var treesNodesTestData map[string]testNodes

type wrelationTestItem struct {
	relations []wrelation
	expect    float64
}

type testNodes struct {
	root   string
	nodes  nodes
	expect []string
}
type wrelationsTD map[string]wrelationTestItem

var weightedRelationsData wrelationsTD

func (x wrelationsTD) getSymmetric() wrelationsTD {
	result := make(wrelationsTD)
	for key, relations := range x {
		new := []wrelation{}
		for _, v := range relations.relations {
			new = append(new, v)
			new = append(new, wrelation{from: v.to, to: v.from, dist: v.dist})
		}
		result[key] = wrelationTestItem{new, x[key].expect}
	}
	return result
}

func initTreesTestData() {
	weightedRelationsData = make(wrelationsTD)
	weightedRelationsData["threeCycle"] = wrelationTestItem{relations: []wrelation{
		{"a", "b", 1},
		{"b", "c", 2},
		{"c", "a", 3},
	}, expect: 3}
	weightedRelationsData["0705_fig750_minus_H"] = wrelationTestItem{relations: []wrelation{
		{"A", "B", 3},
		{"A", "C", 2},
		{"B", "C", 3},
		{"B", "D", 6},
		{"C", "E", 5},
		{"C", "F", 5},
		{"D", "E", 2},
		{"E", "F", 4},
		{"E", "G", 3},
		{"F", "G", 4},
	}, expect: -123}
	weightedRelationsData["0705_fig750"] = wrelationTestItem{relations: []wrelation{
		{"A", "B", 3},
		{"A", "C", 2},
		{"B", "C", 3},
		{"B", "D", 6},
		{"C", "E", 5},
		{"C", "F", 5},
		{"D", "E", 2},
		{"D", "H", 2},
		{"E", "F", 4},
		{"E", "G", 3},
		{"F", "G", 4},
		{"G", "H", 6},
	}, expect: 21}
	weightedRelationsData["0705_fig751"] = wrelationTestItem{relations: []wrelation{
		{"A", "B", 2.7},
		{"A", "I", 2.4},
		{"A", "H", 3.4},
		{"B", "C", 2.6},
		{"B", "I", 2.1},
		{"C", "E", 4.2},
		{"C", "I", 3.6},
		{"C", "J", 3.3},
		{"C", "D", 2.9},
		{"D", "E", 2.2},
		{"D", "G", 5.3},
		{"D", "F", 2.1},
		{"E", "F", 4.4},
		{"F", "J", 2.5},
		{"G", "J", 3.2},
		{"G", "H", 1.7},
		{"H", "I", 1.8},
		{"I", "J", 2.8},
	}, expect: 20.2}
	weightedRelationsData["0705_fig760"] = wrelationTestItem{relations: []wrelation{
		{"A", "B", 3.7},
		{"A", "G", 3.5},
		{"A", "H", 6.8},
		{"B", "C", 3.7},
		{"B", "F", 2.8},
		{"B", "G", 3.5},
		{"C", "D", 4.3},
		{"D", "E", 1.7},
		{"D", "L", 7.2},
		{"D", "N", 2.7},
		{"E", "F", 1.9},
		{"F", "I", 2.1},
		{"G", "I", 2.7},
		{"H", "I", 4.9},
		{"H", "J", 2.2},
		{"I", "K", 2.6},
		{"J", "K", 2.3},
		{"K", "L", 2.3},
		{"L", "M", 2.3},
		{"M", "N", 1.7},
	}, expect: 31.8}
	weightedRelationsData = weightedRelationsData.getSymmetric()

	nod := make(map[string]testNodes)
	nod["cp0704_exp_07"] = testNodes{"a", nodes{
		"a": newSet("c", "d"),
		"b": newSet("c", "d"),
		"c": newSet("a", "b"),
		"d": newSet("a", "b"),
	}, []string{"a", "c", "b", "d"}}
	nod["cp0704_ex_07"] = testNodes{"e", nodes{
		"e": newSet("d", "f"),
		"a": newSet("b", "f"),
		"b": newSet("a", "c"),
		"c": newSet("b", "d"),
		"d": newSet("e", "c"),
		"f": newSet("a", "e"),
	}, []string{"e", "d", "c", "b", "a", "f"}}
	nod["cp0704_ex_09"] = testNodes{"c", nodes{
		"a": newSet("b", "c", "d"),
		"b": newSet("a", "c", "d"),
		"c": newSet("b", "a", "d"),
		"d": newSet("b", "c", "a"),
	}, []string{"c", "a", "b", "d"}}
	nod["cp0704_ex_11"] = testNodes{"e", nodes{
		"e": newSet("a", "c"),
		"a": newSet("e", "b", "c"),
		"b": newSet("a", "d", "c"),
		"c": newSet("b", "d", "e"),
		"d": newSet("a", "b", "c"),
	}, []string{"e", "a", "b", "c", "d"}}
	nod["cp0704_ex_11"] = testNodes{"e", nodes{
		"e": newSet("a", "c"),
		"a": newSet("e", "b", "c"),
		"b": newSet("a", "d", "c"),
		"c": newSet("b", "d", "e"),
		"d": newSet("a", "b", "c"),
	}, []string{"e", "a", "b", "c", "d"}}
	treesNodesTestData = nod

	mat := make(map[string]treeMatrix)
	mat["cp0704_exp_07"] = treeMatrix{m: map[string][]int{
		"a": {0, 0, 1, 1},
		"b": {0, 0, 1, 1},
		"c": {1, 1, 0, 0},
		"d": {1, 1, 0, 0},
	},
		keyIdx: map[string]int{"a": 0, "b": 1, "c": 2, "d": 3},
	}
	mat["cp0704_ex_07"] = treeMatrix{m: map[string][]int{
		"e": {0, 0, 0, 0, 1, 1},
		"a": {0, 0, 1, 0, 0, 1},
		"b": {0, 1, 0, 1, 0, 0},
		"c": {0, 0, 1, 0, 1, 0},
		"d": {1, 0, 0, 1, 0, 0},
		"f": {1, 1, 0, 0, 0, 0},
	},
		keyIdx: map[string]int{"e": 0, "a": 1, "b": 2, "c": 3, "d": 4, "f": 5},
	}
	mat["cp0704_ex_13"] = treeMatrix{m: map[string][]int{
		"Abbeville": {0, 69, 121, 30, 113, 70, 135, 63},
		"Aiken":     {69, 0, 52, 97, 170, 117, 163, 16},
		"Allendale": {121, 52, 0, 149, 222, 160, 206, 59},
		"Anderson":  {30, 97, 149, 0, 92, 63, 122, 93},
		"Asheville": {113, 170, 222, 92, 0, 155, 204, 174},
		"Athens":    {70, 117, 160, 63, 155, 0, 66, 101},
		"Atlanta":   {135, 163, 206, 122, 204, 66, 0, 147},
		"Augusta":   {63, 16, 59, 93, 174, 101, 147, 0},
	},
		keyIdx: map[string]int{"Abbeville": 0, "Aiken": 1, "Allendale": 2, "Anderson": 3, "Asheville": 4, "Athens": 5, "Atlanta": 6, "Augusta": 7},
	}
	treesMatrixTestData = mat
}

func treeMatrixIntoRelationsTest(t *testing.T) {
	tt := treesMatrixTestData["cp0704_ex_13"]
	relations := tt.treeMatrixIntoRelations()
	fmt.Println(relations)
}

func (x treeMatrix) treeMatrixIntoRelations() []wrelation {
	result := []wrelation{}
	for toKey, toIdx := range x.keyIdx {
		for fromKey, row := range x.m {
			v := float64(row[toIdx])
			if v == 0 {
				continue
			}
			result = append(result, wrelation{from: fromKey, to: toKey, dist: v})
		}
	}
	return result
}
