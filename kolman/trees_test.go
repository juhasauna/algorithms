package kolman

import (
	"fmt"
	"reflect"
	"testing"
)

var treesMatrixTestData map[string]treeMatrix
var treesNodesTestData map[string]testNodes

type testNodes struct {
	root   string
	nodes  nodes
	expect []string
}

func Test_trees(t *testing.T) {
	initWeightedRelationsTestData()
	initTreesTestData()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"spanningTreeMatrixPrims", spanningTreeMatrixPrimsTest},
		// {"spanningTreePrimsAlg", spanningTreePrimsAlgTest},
		{"shortestDistancePrimsAlgTest", shortestDistancePrimsAlgTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func shortestDistancePrimsAlgTest(t *testing.T) {
	tests := []struct {
		x       treePathGetter
		minDist float64
		maxDist float64
	}{
		// {treePathGetter{rootName: "A", data: weightedRelationsData["0705_fig750"].relations}, 21, 32},
		// {treePathGetter{rootName: "A", data: weightedRelationsData["0705_fig751"].relations}, 20.2, 32.7},
		// {treePathGetter{rootName: "A", data: weightedRelationsData["0705_fig760"].relations}, 31.8, 48.7},
		{treePathGetter{rootName: "A", data: weightedRelationsData["0705_fig760"].relations}, 31.8, 48.7},
	}
	for _, tt := range tests {
		// tt.x.minMaxDistance(true)
		// if !reflect.DeepEqual(tt.x.distance, tt.minDist) {
		// 	t.Errorf("FAIL: expected dist: %f, got dist %f", tt.minDist, tt.x.distance)
		// }
		// fmt.Printf("%v\n", tt.x.visited)
		tt.x.minMaxDistance(false)
		if tt.x.distance != tt.maxDist {
			t.Errorf("FAIL: expected dist: %f, got dist %f", tt.maxDist, tt.x.distance)
		}
		fmt.Printf("%v\n", tt.x.visited)
	}
}

func spanningTreeMatrixPrimsTest(t *testing.T) {
	m := treesMatrixTestData["cp0704_exp_07"]
	got := m.spanningTreeMatrixPrims("a")

	fmt.Println(got)
}

func spanningTreePrimsAlgTest(t *testing.T) {
	tests := []testNodes{
		treesNodesTestData["cp0704_ex_11"],
	}
	for _, tt := range tests {
		got := spanningTreePrimsAlg(tt.nodes, tt.root)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL: expected: %v, got: %v", tt.expect, got)
		}
	}
}

type wrelationTestItem struct {
	relations []wrelation
	expect    float64
}

type wrelationsTD map[string]wrelationTestItem

var weightedRelationsData wrelationsTD

func (x wrelationsTD) getSymmetric() wrelationsTD {
	result := make(wrelationsTD)
	for key, relations := range x {
		new := []wrelation{}
		for _, v := range relations.relations {
			new = append(new, v)
			new = append(new, wrelation{a: v.b, b: v.a, dist: v.dist})
		}
		result[key] = wrelationTestItem{new, x[key].expect}
	}
	return result
}

func initWeightedRelationsTestData() {
	weightedRelationsData = make(wrelationsTD)
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
}

func initTreesTestData() {

	// treesNodesTestData2["0705_fig750"] = testNodes2{
	// 	node2: node2{name: "A"},
	// }

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
	treesMatrixTestData = mat
}
