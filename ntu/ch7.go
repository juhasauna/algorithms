package ntu

import (
	"algorithms/ut"
	"log"
)

type CH7 struct {
	iters int
}

type Edge struct {
	a int
	b int
}

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw7_s.pdf
// Ex. 1
func (x *CH7) FindEulerianCircuitPseudo(eGraph *ut.EulerianGraph) ([]int, []Edge) {
	if !eGraph.IsEulerian() {
		log.Fatal("graph is not Eulerian")
	}
	// eGraph.Print()
	edgeStack := []Edge{}
	vertexPathStack := []int{}
	unMarkedEdges := eGraph.Copy()
	// fmt.Printf("unmarked: %+v\n", unMarkedEdges)
	var f func(u int)
	f = func(u int) {
		for unMarkedEdges.EdgeCount() > 0 {
			neighbors, ok := unMarkedEdges.Adj[u]
			if !ok {
				log.Fatalf("something went wrong with node: %d", u)
				continue
			} else if len(neighbors) == 0 {
				log.Fatalf("something went wrong with node: %d, neighbors", u)
				unMarkedEdges.RemoveNode(u)
				continue
			}
			maxNeighbor := -1
			for key := range neighbors {
				if key > maxNeighbor {
					// We're taking the max to avoid randomness.
					maxNeighbor = key
				}
			}
			if maxNeighbor < 0 {
				log.Fatalf("something went wrong with aNeighbor")
			}
			unMarkedEdges.RemoveEdge(u, maxNeighbor)
			unMarkedEdges.RemoveIfIsolated(u)
			unMarkedEdges.RemoveIfIsolated(maxNeighbor)
			if unMarkedEdges.HasNode(maxNeighbor) {
				f(maxNeighbor)
				// fmt.Printf("edgeStack: %+v\n", edgeStack)
			}
			edgeStack = append(edgeStack, Edge{a: u, b: maxNeighbor})
		}
		vertexPathStack = append(vertexPathStack, u)
	}
	f(0)
	return vertexPathStack, edgeStack
}

// func (x *CH7) ConstructEulerCircuit(eGraph ut.EulerianGraph) {
// 	eGraph.RemoveCircuit()
// 	for eGraph.HasEdges() {
// 		subCircuit := eGraph.HierholzerFirst()
// 		l := len(subCircuit)
// 		lastSubCircuitNode := subCircuit[l-1]
// 		//
// 		// g.RemoveIsolatedNodes()
// 	}
// }
