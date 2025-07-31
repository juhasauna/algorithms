package ut

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func Test_graph(t *testing.T) {
	initMyGraphs()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"DFS", DFSTest},
		// {"TopologicalSorting", TopologicalSortingTest},
		// {"Dijkstras", DijkstrasTest},
		// {"MCST", MCSTTest},
		{"AllPairsShortestPaths", AllPairsShortestPathsTest},
		// {"FindLongestPath", FindLongestPathTest},
		// {"StronglyConnectedComponents", StronglyConnectedComponentsTest},
		// {"FindNewMCST", FindNewMCSTTtest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func StronglyConnectedComponentsTest(t *testing.T) {
	tests := []struct {
		name string
		want map[string]int
	}{
		{"triangle_cycle", map[string]int{"a": 1, "b": 1, "c": 1}},
		{"YT_WilleamFiset", map[string]int{"0": 3, "1": 3, "2": 2, "3": 2, "4": 1, "5": 1, "6": 3}},
		{"ch7_slides_scc", map[string]int{"a": 3, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1, "g": 1, "h": 3, "i": 3, "j": 3, "k": 2}},
		{"NTU_24HW9/1", map[string]int{"1": 5, "2": 5, "3": 5, "4": 2, "5": 3, "6": 1, "7": 4, "8": 5, "9": 5}},
		{"NTU_24HW9/1b", map[string]int{"1": 3, "2": 3, "3": 3, "4": 2, "5": 2, "6": 1, "7": 2, "8": 3, "9": 3}},
		{"NTU_24HW9/1_sol", map[string]int{"a": 5, "b": 5, "c": 4, "d": 1, "e": 3, "f": 2, "g": 5, "h": 5, "i": 5}},  // in alg2024hw9_s.pdf we need to reverse traversal 9 -> 1 instead of 1-> 9. This applies to nodes and edges.
		{"NTU_24HW9/1b_sol", map[string]int{"a": 3, "b": 3, "c": 2, "d": 1, "e": 2, "f": 2, "g": 3, "h": 3, "i": 3}}, // in alg2024hw9_s.pdf we need to reverse traversal 9 -> 1 instead of 1-> 9. This applies to nodes and edges.
	}
	for _, tt := range tests {
		var g Graph
		var ok bool
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		got := g.TarjansStronglyConnectedComponents()
		PrintMapAsCode(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL %+v", got)
		}
	}
}
func FindLongestPathTest(t *testing.T) {
	type rmCycles struct {
		method    string
		startNode string
	}
	tests := []struct {
		name         string
		want         []string
		removeCycles *rmCycles
	}{
		{"YT_D.Sutantyo 8.1", []string{"1", "2", "4", "3"}, &rmCycles{"bfs", "1"}},
		{"YT_D.Sutantyo 8.1", []string{"1", "2", "4", "3"}, &rmCycles{"dfs", "1"}},
		{"YT_D.Sutantyo 8.4", []string{"5", "1", "3", "2", "6"}, nil},
		{"ManberFig7.9(p.213)", []string{"a", "b", "d"}, &rmCycles{"bfs", "a"}},
		{"ManberFig7.9(p.213)", []string{"a", "b", "d", "f"}, &rmCycles{"dfs", "a"}},
		{"ManberFig7.18(p.224)", []string{"v", "a", "c", "d", "e", "h", "g", "f"}, nil},
	}
	for _, tt := range tests {
		var g Graph
		var ok bool
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		name := fmt.Sprintf("%s %v", tt.name, tt.removeCycles)
		if tt.removeCycles != nil {
			switch tt.removeCycles.method {
			case "bfs":
				g = NewGraph(name, g.BFS(tt.removeCycles.startNode))
			case "dfs":
				g = NewGraph(name, g.DFS(tt.removeCycles.startNode))
			}
		}
		got := g.FindLongestPath()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL %s got/want\n%v\n%v", name, got, tt.want)
		}
		t.Logf("%v", got)
	}
}
func AllPairsShortestPathsTest(t *testing.T) {
	tests := []struct {
		name                string
		negativeWeightCycle bool
	}{
		// {"YT_D.Sutantyo 8.1", false},
		// {"YT_D.Sutantyo 8.3 (negative weights)", false},
		// {"YT_D.Sutantyo 8.4", false},
		{"NegativeEdgeCycle1", true},
		// {"ManberFig7.18(p.224)", false},
		// {"ManberFig7.21(p.229)", false},
	}
	for _, tt := range tests {
		var g Graph
		var ok bool
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		got := g.AllPairsShortestPaths()
		// if !reflect.DeepEqual(got, tt.want) {
		// 	t.Errorf("FAIL got/want\n%v\n%v", got, tt.want)
		// }
		for _, v := range got {
			t.Logf("%v", v)
		}
		negativeWeightCycle := DetectNegativeWeightCycle(got)
		// t.Logf("%+v", negativeWeightCycle)
		if negativeWeightCycle != tt.negativeWeightCycle {
			t.Errorf("FAIL DetectNegativeWeightCycle(got). got/want: %t/%t", negativeWeightCycle, tt.negativeWeightCycle)
		}
	}
}
func FindNewMCSTTtest(t *testing.T) {
	tests := []struct {
		name string
		want []UndirectedEdge
	}{
		// {"ManberFig7.18(p.224)", []UndirectedEdge{{"a", "v", 1}, {"e", "h", 1}, {"a", "c", 2}, {"d", "e", 2}, {"f", "g", 2}, {"b", "e", 3}, {"g", "h", 3}, {"c", "d", 4}}},
		// {"ManberFig7.21(p.229)", []UndirectedEdge{{"a", "v", 1}, {"a", "c", 2}, {"b", "e", 3}, {"c", "d", 4}, {"e", "h", 5}, {"b", "v", 6}, {"c", "f", 10}, {"g", "h", 11}}},
		// {"YT_D.Sutantyo 8.1", []UndirectedEdge{{"4", "5", 1}, {"2", "3", 2}, {"2", "5", 3}, {"1", "2", 4}}},
		{"YT_WilliamFisetPrims1", []UndirectedEdge{{"1", "3", -2}, {"0", "2", 0}, {"5", "6", 1}, {"3", "5", 2}, {"1", "4", 3}, {"0", "3", 5}}},
		// {"YT_WilliamFisetPrims2", []UndirectedEdge{{"1", "4", 0}, {"0", "2", 1}, {"4", "5", 1}, {"2", "3", 2}, {"3", "5", 2}, {"5", "6", 6}, {"4", "7", 8}}},
	}
	for _, tt := range tests {
		var g Graph
		var ok bool
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		mcst := g.MCST()

		newMcst := g.FindNewMCST(mcst, UndirectedEdge{a: "2", b: "5", weight: 1})
		t.Logf("%+v", newMcst)

	}
}
func MCSTTest(t *testing.T) {
	// Prim's and Kruskal For undirected graphs
	// Kruskal = O(E*log E)
	// Prim's = O(E*log V)
	// // Thus Kurskal is better for sparse graphs and Prim's for dense.
	tests := []struct {
		name string
		want []UndirectedEdge
	}{
		{"ManberFig7.18(p.224)", []UndirectedEdge{{"a", "v", 1}, {"e", "h", 1}, {"a", "c", 2}, {"d", "e", 2}, {"f", "g", 2}, {"b", "e", 3}, {"g", "h", 3}, {"c", "d", 4}}},
		{"ManberFig7.21(p.229)", []UndirectedEdge{{"a", "v", 1}, {"a", "c", 2}, {"b", "e", 3}, {"c", "d", 4}, {"e", "h", 5}, {"b", "v", 6}, {"c", "f", 10}, {"g", "h", 11}}},
		{"YT_D.Sutantyo 8.1", []UndirectedEdge{{"4", "5", 1}, {"2", "3", 2}, {"2", "5", 3}, {"1", "2", 4}}},
		{"YT_WilliamFisetPrims1", []UndirectedEdge{{"1", "3", -2}, {"0", "2", 0}, {"5", "6", 1}, {"3", "5", 2}, {"1", "4", 3}, {"0", "3", 5}}},
		{"YT_WilliamFisetPrims2", []UndirectedEdge{{"1", "4", 0}, {"0", "2", 1}, {"4", "5", 1}, {"2", "3", 2}, {"3", "5", 2}, {"5", "6", 6}, {"4", "7", 8}}},
	}
	for _, tt := range tests {
		var g Graph
		var ok bool
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		got := g.MCST()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL (Prims) got/want\n%v\n%v", got, tt.want)
		}
		// t.Logf("%v", got)
		got = g.Kruskal()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL (Kruskal) got/want\n%v\n%v", got, tt.want)
		}
		got = g.HeapPrims()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL (Prim's) got/want\n%v\n%v", got, tt.want)
		}

		// t.Logf("%v", got)
	}
}
func DijkstrasTest(t *testing.T) {
	tests := []struct {
		name      string
		startNode string
		// want      []DijkstraNode
		want map[string]int
	}{
		{"ManberFig7.18(p.224)", "a", map[string]int{"a": 0, "c": 2, "d": 6, "e": 8, "f": 11, "g": 10, "h": 9}},
		{"ManberFig7.18(p.224)", "b", map[string]int{"b": 0, "e": 3, "f": 9, "g": 7, "h": 4}},
		{"ManberFig7.18(p.224)", "v", map[string]int{"a": 1, "b": 5, "c": 3, "d": 7, "e": 8, "f": 12, "g": 11, "h": 9, "v": 0}},
		{"YT_D.Sutantyo 8.1", "1", map[string]int{"1": 0, "2": 4, "3": 13, "4": 8, "5": 7}},
		{"YT_D.Sutantyo 8.1", "4", map[string]int{"2": 7, "3": 5, "4": 0, "5": 10}},
		{"YT_D.Sutantyo 8.4", "5", map[string]int{"1": 7, "2": 11, "3": 9, "4": 2, "5": 0, "6": 3}},
		{"YT_D.Sutantyo 8.4", "1", map[string]int{"1": 0, "2": 4, "3": 2, "4": 2, "6": 3}},
		{"YT_D.Sutantyo 8.4", "3", map[string]int{"2": 2, "3": 0, "4": 3, "6": 4}},
		{"Grok1", "a", map[string]int{"a": 0, "b": 3, "c": 2}},
		{"Claude1", "a", map[string]int{"a": 0, "b": 3, "c": 1}},
		{"Claude2", "a", map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 3}},
	}
	for _, tt := range tests {
		var g Graph
		var ok bool
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		got := g.DijkstrasNTUPseudo(tt.startNode)
		PrintMapAsCode(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL got/want\n%v\n%v", got, tt.want)
		}
		got = g.DPSSSP(tt.startNode)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL (DPSSSP) got/want\n%v\n%v", got, tt.want)
		}

		// got = g.DijkstrasSimpleButNotOptimal(tt.startNode)
		got = g.Dijkstras(tt.startNode)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL %+v != %+v", got, tt.want)
		} else {

			// t.Logf("%+v", got2)
		}
	}
}
func TopologicalSortingTest(t *testing.T) {
	tests := []struct {
		name      string
		startNode string
		bfs       bool
		want      []string
	}{
		// {"triangle_cycle", "a", true, []string{"a", "b", "c"}},
		// {"triangle_cycle", "a", false, []string{"a", "b", "c"}},
		{"ManberFig7.9(p.213)", "a", false, []string{"a", "b", "c", "d", "e", "f", "g"}},
		{"ManberFig7.9(p.213)", "a", true, []string{"a", "b", "c", "g", "d", "e", "f"}},
		{"ManberFig7.18(p.224)", "a", false, []string{"a", "c", "d", "e", "h", "g", "f"}},
		{"ManberFig7.18(p.224)", "a", true, []string{"a", "c", "d", "f", "e", "g", "h"}},
	}
	for _, tt := range tests {
		var g Graph
		{
			var temp Graph
			ok := false
			if temp, ok = MyGraphs[tt.name]; !ok {
				t.Errorf("Graph %s does not exits", tt.name)
				return
			}
			acyclicEdges := []Edge{}
			if tt.bfs {
				acyclicEdges = temp.BFS(tt.startNode)
			} else {
				acyclicEdges = temp.DFS(tt.startNode)
			}
			g = NewGraph(tt.name, acyclicEdges)
		}
		got := topologicalSortResultToSlice(g.TopologicalSorting())

		if !slices.Equal(got, tt.want) {
			t.Errorf("FAIL got/want\n%v\n%v", got, tt.want)
		}
		shortestPaths := g.AcyclicShortestPaths(tt.startNode)

		shortestPathName := tt.name + "_dfs_" + tt.startNode
		if tt.bfs {
			shortestPathName = tt.name + "_bfs_" + tt.startNode
		}
		wantShortestPaths := shortestAcyclicPaths[shortestPathName]
		if !reflect.DeepEqual(shortestPaths, wantShortestPaths) {
			t.Errorf("FAIL (shortestPahtName: %s) got/want\n%v\n%v", shortestPathName, shortestPaths, wantShortestPaths)
		}
	}
}

func DFSTest(t *testing.T) {
	tests := []struct {
		name      string
		startNode string
		want      []Tuple
		bfs       bool
	}{
		// {"triangle_cycle", "a", []Tuple{{"a", "b"}, {"b", "c"}}, false},
		// {"triangle_cycle", "a", []Tuple{{"a", "b"}, {"b", "c"}}, true},
		// {"triangle", "a", []Tuple{{"a", "b"}, {"a", "c"}}, true},
		// {"bowtie_a1", "a", []Tuple{{"a", "b"}, {"b", "m"}, {"m", "c"}, {"c", "d"}}, false},
		// {"bowtie_m1", "m", []Tuple{{"m", "a"}, {"a", "b"}, {"m", "c"}, {"c", "d"}}, false},
		// {"ManberFig7.9(p.213)", "a", []Tuple{{"a", "b"}, {"b", "d"}, {"d", "f"}, {"d", "g"}, {"a", "c"}, {"c", "e"}}, false},
		{"ManberFig7.9(p.213)", "a", []Tuple{{"a", "b"}, {"a", "c"}, {"a", "g"}, {"b", "d"}, {"c", "e"}, {"g", "f"}}, true},
	}
	for _, tt := range tests {
		var g Graph
		ok := false
		if g, ok = MyGraphs[tt.name]; !ok {
			t.Errorf("Graph %s does not exits", tt.name)
			return
		}
		got := []Edge{}
		if tt.bfs {
			got = g.BFS(tt.startNode)
		} else {
			got = g.DFS(tt.startNode)
		}
		gotGraph := NewGraph(tt.name+"solution", got)
		gotGraph.PrintMatrix()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL %s got!=want\n%v\n%v", tt.name, got, tt.want)
		}
		t.Logf("%v", got)
	}
}

var MyGraphs map[string]Graph
var shortestAcyclicPaths map[string]map[string]int
var dijkstrasResults map[string]map[string]int

func initMyGraphs() {
	shortestAcyclicPaths = make(map[string]map[string]int)
	dijkstrasResults = make(map[string]map[string]int)

	m := make(map[string]Graph)
	name := "triangle_cycle"
	m[name] = NewGraph(name, []Tuple{{"a", "b"}, {"b", "c"}, {"c", "a"}})
	name = "triangle"
	m[name] = NewGraph(name, []Tuple{{"a", "b"}, {"a", "c"}, {"b", "c"}})
	// bowtie
	// a	   c
	// | > m < |
	// b	   d
	name = "bowtie_a1"
	m[name] = NewGraph(name, []Tuple{{"a", "b"}, {"a", "m"}, {"b", "m"}, {"m", "c"}, {"c", "d"}, {"d", "m"}})
	name = "bowtie_m1"
	m[name] = NewGraph(name, []Tuple{{"m", "a"}, {"m", "b"}, {"m", "c"}, {"m", "d"}, {"c", "d"}, {"a", "b"}})
	name = "ManberFig7.9(p.213)"
	m[name] = NewGraph(name, []Tuple{{"a", "b"}, {"a", "c"}, {"a", "g"}, {"b", "d"}, {"c", "e"}, {"d", "f"}, {"d", "g"}, {"e", "g"}, {"g", "b"}, {"g", "f"}})
	shortestAcyclicPaths[name+"_dfs_a"] = map[string]int{"a": 0, "b": 1, "c": 1, "d": 2, "e": 2, "f": 3, "g": 3}
	shortestAcyclicPaths[name+"_bfs_a"] = map[string]int{"a": 0, "b": 1, "c": 1, "d": 2, "e": 2, "f": 2, "g": 1}

	name = "ManberFig7.18(p.224)"
	m[name] = NewGraph(name, []Edge{{"a", "c", 2}, {"v", "a", 1}, {"v", "b", 5}, {"v", "d", 9}, {"b", "e", 3}, {"c", "d", 4}, {"d", "e", 2}, {"c", "f", 9}, {"d", "g", 4}, {"e", "h", 1}, {"g", "f", 2}, {"h", "g", 3}})
	shortestAcyclicPaths[name+"_dfs_a"] = map[string]int{"a": 0, "c": 2, "d": 6, "e": 8, "f": 14, "g": 12, "h": 9}
	shortestAcyclicPaths[name+"_bfs_a"] = map[string]int{"a": 0, "c": 2, "d": 6, "e": 8, "f": 11, "g": 10, "h": 9}
	dijkstrasResults[name+"_a"] = map[string]int{"a": 0, "b": INF, "c": 2, "d": 6, "e": 8, "f": 11, "g": 10, "h": 9, "v": INF}
	dijkstrasResults[name+"_b"] = map[string]int{"a": INF, "b": 0, "c": INF, "d": INF, "e": 3, "f": 9, "g": 7, "h": 4, "v": INF}
	dijkstrasResults[name+"_v"] = map[string]int{"a": 1, "b": 5, "c": 3, "d": 7, "e": 8, "f": 12, "g": 11, "h": 9, "v": 0}
	name = "YT_D.Sutantyo 8.1"
	m[name] = NewGraph(name, []Edge{{"1", "2", 4}, {"1", "5", 8}, {"2", "4", 9}, {"2", "5", 3}, {"3", "2", 2}, {"4", "3", 5}, {"5", "4", 1}})
	dijkstrasResults[name+"_1"] = map[string]int{"1": 0, "2": 4, "3": 13, "4": 8, "5": 7}
	dijkstrasResults[name+"_4"] = map[string]int{"1": 32767, "2": 7, "3": 5, "4": 0, "5": 10}

	name = "YT_D.Sutantyo 8.3 (negative weights)"
	m[name] = NewGraph(name, []Edge{{"1", "2", 8}, {"1", "5", -4}, {"2", "4", 9}, {"2", "5", 3}, {"3", "2", 7}, {"4", "3", -5}, {"5", "4", 1}})

	name = "YT_D.Sutantyo 8.4"
	m[name] = NewGraph(name, []Edge{{"1", "2", 11}, {"1", "3", 2}, {"1", "4", 2}, {"2", "6", 7}, {"3", "2", 2}, {"3", "4", 3}, {"3", "6", 12}, {"4", "6", 1}, {"5", "1", 7}, {"5", "4", 2}})

	name = "ManberFig7.21(p.229)"
	m[name] = NewGraph(name, []Edge{{"a", "c", 2}, {"v", "a", 1}, {"v", "b", 6}, {"v", "d", 9}, {"b", "e", 3}, {"c", "d", 4}, {"d", "e", 7}, {"c", "f", 10}, {"d", "g", 12}, {"e", "h", 5}, {"g", "f", 13}, {"h", "g", 11}})

	name = "Grok1"
	m[name] = NewGraph(name, []Edge{{"a", "b", 4}, {"a", "c", 2}, {"c", "b", 1}})

	name = "Claude1"
	m[name] = NewGraph(name, []Edge{{"a", "b", 5}, {"a", "c", 1}, {"c", "b", 2}})
	name = "Claude2"
	m[name] = NewGraph(name, []Edge{{"a", "b", 1}, {"a", "c", 2}, {"a", "d", 3}, {"b", "e", 10}, {"c", "e", 1}, {"d", "e", 1}})

	name = "NegativeEdgeCycle1"
	m[name] = NewGraph(name, []Edge{{"a", "b", 1}, {"b", "c", 1}, {"c", "a", -10}})

	name = "YT_WilliamFisetPrims1" // https://www.youtube.com/watch?v=jsmMtJpPnhU&t=158s
	m[name] = NewGraph(name, []Edge{{"0", "1", 9}, {"0", "2", 0}, {"0", "3", 5}, {"0", "5", 7},
		{"1", "3", -2}, {"1", "4", 3}, {"1", "6", 4},
		{"2", "5", 6},
		{"3", "5", 2}, {"3", "6", 3},
		{"4", "6", 6},
		{"5", "6", 1}})
	name = "YT_WilliamFisetPrims2" // https://www.youtube.com/watch?v=jsmMtJpPnhU&t=158s
	m[name] = NewGraph(name, []Edge{{"0", "1", 10}, {"0", "2", 1}, {"0", "3", 4},
		{"1", "2", 3}, {"1", "4", 0},
		{"2", "3", 2}, {"2", "5", 8},
		{"3", "5", 2}, {"3", "6", 7},
		{"4", "5", 1}, {"4", "7", 8},
		{"5", "6", 6}, {"5", "7", 9},
		{"6", "7", 12}})

	name = "ch7_slides_scc" // file:///C:\Users\FIJUSAU\OneDrive%20-%20ABB\courses\Vaihto\TaiwanTech\algorithms_2024_material\slides\ch7_slides_scc.pptx
	m[name] = NewGraph(name, []Tuple{{"a", "b"}, {"a", "h"},
		{"b", "c"},
		{"c", "d"}, {"c", "f"},
		{"d", "e"},
		{"e", "b"},
		{"f", "g"},
		{"g", "e"},
		{"h", "c"}, {"h", "i"},
		{"i", "j"}, {"i", "k"},
		{"j", "a"}, {"j", "f"},
		{"k", "g"},
	})
	name = "YT_WilleamFiset" // https://www.youtube.com/watch?v=wUgWX0nc4NY&t=598s
	m[name] = NewGraph(name, []Tuple{{"0", "1"},
		{"1", "2"}, {"1", "4"}, {"1", "6"},
		{"2", "3"},
		{"3", "2"}, {"3", "4"}, {"3", "5"},
		{"4", "5"},
		{"5", "4"},
		{"6", "0"}, {"6", "1"},
	})
	name = "NTU_24HW9/1" // file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw9.pdf
	m[name] = NewGraph(name, []Tuple{{"1", "9"},
		{"2", "7"}, {"2", "8"},
		{"3", "1"}, {"3", "2"},
		{"4", "6"},
		{"5", "4"},
		{"7", "5"}, {"7", "6"},
		{"8", "3"}, {"8", "7"},
		{"9", "8"},
	})
	name = "NTU_24HW9/1b" // file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw9.pdf
	m[name] = NewGraph(name, []Tuple{{"1", "9"},
		{"2", "7"}, {"2", "8"},
		{"3", "1"}, {"3", "2"},
		{"4", "6"}, {"4", "7"},
		{"5", "4"},
		{"7", "5"}, {"7", "6"},
		{"8", "3"}, {"8", "7"},
		{"9", "8"},
	})
	name = "NTU_24HW9/1_sol" // file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw9.pdf
	m[name] = NewGraph(name, []Tuple{{"i", "a"},
		{"h", "c"}, {"h", "b"},
		{"g", "i"}, {"g", "h"},
		{"f", "d"},
		{"e", "f"},
		{"c", "e"}, {"c", "d"},
		{"b", "g"}, {"b", "c"},
		{"a", "b"},
	})
	name = "NTU_24HW9/1b_sol" // file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/hw9.pdf
	m[name] = NewGraph(name, []Tuple{{"i", "a"},
		{"h", "c"}, {"h", "b"},
		{"g", "i"}, {"g", "h"},
		{"f", "d"}, {"f", "c"},
		{"e", "f"},
		{"c", "e"}, {"c", "d"},
		{"b", "g"}, {"b", "c"},
		{"a", "b"},
	})

	MyGraphs = m
}
