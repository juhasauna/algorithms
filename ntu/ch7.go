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

// Ex. 1
func (x *CH7) FindEulerianCircuitPseudo(eGraph *ut.EulerianGraph) ([]int, []Edge) {
	if !eGraph.IsEulerian() {
		log.Fatalf("graph '%s' is not Eulerian", eGraph.Name)
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

// 2 (7.28, p.269) A binary de Bruijn sequence
func (x *CH7) BinaryDeBruijinSequence() {
	// Prove that G is a directed Eulerian graph.
}
