package kolman

import (
	"fmt"
	"testing"
)

func Test_graphs(t *testing.T) {
	initGraphData()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"fleurysAlg", fleurysAlgTest},
		// {"copyEdges", copyEdgesTest},
	}
	for _, v := range tests {
		t.Run(v.name, v.f)
	}
}

func fleurysAlgTest(t *testing.T) {
	tests := []struct {
		g       graph
		wantErr bool
		comment string
	}{

		{graphData["complete_3"], false, ""},
		{graphData["082_ex_11"], false, ""},
		{graphData["082_examp_06"], false, ""},
		// {graphData["082_examp_06_minus_GA"], true, "If graph has a node of odd degree, there can be no Euler circuit."},
	}
	for _, tt := range tests {
		tt.g.fleurysAlg()
		// got := tt.g.getNotBridge()
		fmt.Println(tt.g)
	}
}

func copyEdgesTest(t *testing.T) {
	tests := []struct {
		g graph
	}{
		{graphData["complete_3"]},
	}
	for _, tt := range tests {
		// copy1 := tt.g.copyEdges()
		// fmt.Printf("copy1: %v, orig: %v\n", copy1, tt.g.edges)
		// copy1 = append(copy1[:0], copy1[1:]...)
		// fmt.Printf("rmd copy1: %v, orig: %v\n", copy1, tt.g.edges)
		copy3 := tt.g.edges
		fmt.Printf("copy3: %v, orig: %v\n", copy3, tt.g.edges)
		copy3 = append(copy3[:0], copy3[1:]...)
		fmt.Printf("rmd copy3: %v, orig: %v\n", copy3, tt.g.edges)

	}
}
