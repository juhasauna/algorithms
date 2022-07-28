package kolman

var graphData map[string]graph

func initGraphData() {
	m := make(map[string]graph)
	m["082_examp_06"] = graph{
		nodes: newSet("A", "B", "C", "D", "E", "F", "G", "H"),
		edges: []tuple{
			{"A", "B"}, {"A", "C"}, {"A", "D"}, {"A", "G"},
			{"B", "C"},
			{"C", "D"}, {"C", "E"},
			{"E", "F"}, {"E", "H"}, {"E", "G"},
			{"F", "G"},
			{"G", "H"},
		},
	}
	m["082_examp_06_minus_GA"] = graph{
		nodes: newSet("A", "B", "C", "D", "E", "F", "G", "H"),
		edges: []tuple{
			{"A", "B"}, {"A", "C"}, {"A", "D"},
			{"B", "C"},
			{"C", "D"}, {"C", "E"},
			{"E", "F"}, {"E", "H"}, {"E", "G"},
			{"F", "G"},
			{"G", "H"},
		},
	}
	m["082_ex_11"] = graph{
		nodes: newSet("1", "2", "3", "4", "5", "6"),
		edges: []tuple{{"1", "2"}, {"1", "4"}, {"2", "3"}, {"2", "4"}, {"2", "5"}, {"3", "5"}, {"4", "5"}, {"4", "6"}, {"5", "6"}},
	}
	m["complete_3"] = graph{
		nodes: newSet("1", "2", "3"),
		edges: []tuple{{"1", "2"}, {"2", "3"}, {"3", "1"}}}

	graphData = m
}
