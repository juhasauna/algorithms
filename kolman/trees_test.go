package kolman

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_trees(t *testing.T) {
	initTreesTestData()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"treeMatrixIntoRelations", treeMatrixIntoRelationsTest},
		// {"spanningTreeMatrixPrims", spanningTreeMatrixPrimsTest},
		// {"spanningTreePrimsAlg", spanningTreePrimsAlgTest},
		// {"shortestDistancePrimsAlgTest", shortestDistancePrimsAlgTest},
		// {"isCyclic", isCyclicTest},
		// {"kruskal", kruskalTest},
		{"treePython_kruskal", treePython_kruskalTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func treePython_kruskalTest(t *testing.T) {
	tt := []treePython{
		{0, 1, 8},
		{0, 2, 5},
		{1, 2, 9},
		{1, 3, 11},
		{2, 3, 15},
		{2, 4, 10},
		{3, 4, 7},
	}
	k := kruskalPython{graph: tt, nodes: 5}
	got := k.kruskal()
	fmt.Println(got)
}
func isCyclicTest(t *testing.T) {
	tests := []struct {
		startEdge int
		edges     []wrelation
		want      bool
	}{
		{0, weightedRelationsData["threeCycle"].relations, true},
		{0, []wrelation{{from: "a", to: "b"}, {from: "b", to: "c"}, {from: "c", to: "d"}}, false},
		{0, []wrelation{{from: "a", to: "b"}, {from: "b", to: "a"}}, true},
		{0, []wrelation{{from: "b", to: "c"}, {from: "d", to: "a"}, {from: "c", to: "d"}, {from: "a", to: "b"}}, true},
	}
	for _, tt := range tests {
		start := tt.edges[tt.startEdge]
		got := isCyclic(tt.edges, start, start.from, start.to)
		if got != tt.want {
			t.Errorf("FAIL: expect: %t, got: %t", tt.want, got)
		}
	}
}

func kruskalTest(t *testing.T) {
	tests := []struct {
		x       treePathGetter
		minDist float64
		maxDist float64
	}{
		// {treePathGetter{relations: weightedRelationsData["threeCycle"].relations}, 3, 5},
		{treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig750_minus_H"].relations}, 19, -123},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig750"].relations}, 21, 32},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig751"].relations}, 20.2, 32.7},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig760"].relations}, 31.8, 48.7},
		// {treePathGetter{rootName: "Abbeville", relations: treesMatrixTestData["cp0704_ex_13"].treeMatrixIntoRelations()}, 382, 1216},
	}
	for _, tt := range tests {
		tt.x.kruskal(true)
		if !reflect.DeepEqual(tt.x.distance, tt.minDist) {
			t.Errorf("FAIL: expected dist: %f, got dist %f", tt.minDist, tt.x.distance)
		}
		fmt.Printf("%v\n", tt.x.visited)
		// tt.x.minMaxDistance(false)
		// if tt.x.distance != tt.maxDist {
		// 	t.Errorf("FAIL: expected dist: %f, got dist %f", tt.maxDist, tt.x.distance)
		// }
		// fmt.Printf("%v\n", tt.x.visited)
	}
}
func shortestDistancePrimsAlgTest(t *testing.T) {
	tests := []struct {
		x       treePathGetter
		minDist float64
		maxDist float64
	}{
		{treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig750_minus_H"].relations}, 19, -123},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig750"].relations}, 21, 32},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig751"].relations}, 20.2, 32.7},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig760"].relations}, 31.8, 48.7},
		// {treePathGetter{rootName: "A", relations: weightedRelationsData["0705_fig760"].relations}, 31.8, 48.7},
		// {treePathGetter{rootName: "Abbeville", relations: treesMatrixTestData["cp0704_ex_13"].treeMatrixIntoRelations()}, 382, 1216},
	}
	for _, tt := range tests {
		tt.x.minMaxDistance(true)
		if !reflect.DeepEqual(tt.x.distance, tt.minDist) {
			t.Errorf("FAIL: expected dist: %f, got dist %f", tt.minDist, tt.x.distance)
		}
		fmt.Printf("%v\n", tt.x.visited)
		// tt.x.minMaxDistance(false)
		// if tt.x.distance != tt.maxDist {
		// 	t.Errorf("FAIL: expected dist: %f, got dist %f", tt.maxDist, tt.x.distance)
		// }
		// fmt.Printf("%v\n", tt.x.visited)
	}
}

func spanningTreeMatrixPrimsTest(t *testing.T) {
	m := treesMatrixTestData["cp0704_exp_07"]
	m.rootName = "a"
	got := m.spanningTreeMatrixPrims()

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
