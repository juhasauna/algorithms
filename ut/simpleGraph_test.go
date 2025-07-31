package ut

import (
	"testing"
)

func Test_simpleGraph(t *testing.T) {
	initSimpleGraphs()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"DFSSimple", DFSSimpleTest},
		// {"IsDFSTreeTest", IsDFSTreeTest},
		// {"BuildDFSTree", BuildDFSTreeTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func DFSSimpleTest(t *testing.T) {
	tests := []struct {
		name      string
		startNode string
		want      []Tuple
	}{
		// {"triangle", "a"},
		{"bowtie", "a", []Tuple{{"a", "m"}, {"m", "c"}, {"c", "d"}, {"m", "b"}}},
	}
	for _, tt := range tests {
		g := SimpleGraphs[tt.name]
		got := g.DFS(tt.startNode)
		// sort.Slice(got)

		t.Logf("%v", got)
	}
	// tests := []struct {
	// 	name      string
	// 	startNode string
	// 	edges     []Tuple
	// 	want      []string
	// }{
	// 	{"", "a", []Tuple{{"a", "b"}, {"b", "c"}, {"b", "d"}, {"c", "e"}, {"d", "f"}}, []string{"a", "b", "c", "e", "d", "f"}},
	// }

	// for _, tt := range tests {
	// 	g := NewSimpleGraph(tt.edges)
	// 	got := g.DFS(tt.startNode)
	// 	if !slices.Equal(got, tt.want) {
	// 		t.Errorf("FAIL got/want:\n%v\n%v", got, tt.want)
	// 	}
	// 	t.Logf("%v", got)
	// }
}
func IsDFSTreeTest(t *testing.T) {
	tests := []struct {
		name           string
		startNode      string
		graphEdges     []Tuple
		candidateEdges []Tuple
		want           bool
	}{
		// {"", "a", []Tuple{{"a", "b"}, {"b", "c"}, {"c", "d"}}, []Tuple{{"a", "b"}, {"b", "c"}, {"c", "d"}}, true},
		{"", "a", []Tuple{{"a", "b"}, {"b", "c"}, {"b", "d"}, {"c", "d"}}, []Tuple{{"a", "b"}, {"b", "c"}, {"c", "d"}}, true},
		// {"", "a", []Tuple{{"a", "b"}, {"b", "c"}, {"b", "d"}, {"c", "e"}, {"d", "f"}}, []Tuple{{"a", "b"}, {"b", "c"}, {"c", "a"}}, false},
	}

	for _, tt := range tests {
		g := NewSimpleGraph(tt.graphEdges)
		candidate := NewSimpleGraph(tt.candidateEdges)
		got := g.IsDFSTree(tt.startNode, candidate)
		if got != tt.want {
			t.Errorf("FAIL got/want: %t/%t", got, tt.want)
		}
		t.Logf("%v", got)
	}
}

var SimpleGraphs map[string]SimpleGraph

func initSimpleGraphs() {
	m := make(map[string]SimpleGraph)
	m["triangle"] = NewSimpleGraph([]Tuple{{"a", "b"}, {"b", "c"}, {"c", "a"}})
	// bowtie
	// a	   c
	// | > m < |
	// b	   d
	// {"bowtie", "m", []Tuple{{"m", "a"}, {"m", "b"}, {"m", "c"}, {"m", "d"}, {"a", "b"}, {"c", "d"}}, []Tuple{{"a", "b"}, {"b", "c"}}},
	m["bowtie"] = NewSimpleGraph([]Tuple{{"m", "a"}, {"m", "b"}, {"m", "c"}, {"m", "d"}, {"a", "b"}, {"c", "d"}})
	SimpleGraphs = m
}
