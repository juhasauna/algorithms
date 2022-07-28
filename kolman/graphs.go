package kolman

import (
	"fmt"
	"log"
)

type tuple struct {
	a string
	b string
}

type graph struct {
	nodes   set
	edges   []tuple
	circuit []string
}

func (g *graph) fleurysAlg() {
	edge := g.getNotBridge()
	if edge == nil {
		log.Fatal("No non-bridge edge found!")
	}
	first := true
	for len(g.edges) > 0 {
		g.rmEdge(*edge)
		if g.isIsolated(edge.a) {
			if first {
				g.circuit = append(g.circuit, edge.b)
				first = false
			}
			g.circuit = append(g.circuit, edge.a)
			*edge = g.getIsolatedEdge(edge.a)
		} else if g.isIsolated(edge.b) {
			if first {
				g.circuit = append(g.circuit, edge.a)
				first = false
			}
			g.circuit = append(g.circuit, edge.b)
			*edge = g.getIsolatedEdge(edge.b)
		} else {
			if first {
				g.circuit = append(g.circuit, edge.a)
				g.circuit = append(g.circuit, edge.b)
				first = false
			} else {
				lastNode := edge.other(g.getLastNode())
				g.circuit = append(g.circuit, lastNode)
			}
			// lastNode := g.getLastNode()
			// if e := g.getNotBridgeFromNode(lastNode); e != nil {
			// 	edge = e
			// 	// g.circuit = append(g.circuit, edge.other(lastNode))
			// 	first = true
			// } else {
			*edge = g.getNextEdge()
			// first = true
			// g.circuit = append(g.circuit, edge.other(lastNode))
			// }
		}
		if len(g.edges) == 1 {
			break
		}
	}
	fmt.Println(edge)
	finalNode := edge.other(g.getLastNode())
	g.circuit = append(g.circuit, finalNode)
}

func (g *graph) getIsolatedEdge(node string) tuple {
	for _, v := range g.edges {
		if v.a == node || v.b == node {
			return v
		}
	}
	log.Fatalf("getIsolatedEdge didn't find edge for node: %s\n", node)
	return tuple{}
}

func (g *graph) appendCircuit(node string) {
	if len(g.circuit) == 0 {
		g.circuit = append(g.circuit, node)
		return
	}
	lastNode := g.getLastNode()
	if lastNode == node {
		return
	}
	g.circuit = append(g.circuit, node)

}

func (g *graph) getLastNode() string {
	if len(g.circuit) == 0 {
		log.Fatal("empty circuit")
	}
	return g.circuit[len(g.circuit)-1]
}

func (x *tuple) other(v string) string {
	if x.a == v {
		return x.b
	}
	return x.a
}

func (g *graph) getNextEdge() tuple {
	node := g.getLastNode()
	for _, v := range g.edges {
		if v.a == node || v.b == node {
			return v
		}
	}
	log.Fatalf("no edges for node: %s\n", node)
	return tuple{}
}
func (g *graph) isIsolated(node string) bool {
	found := false
	for _, v := range g.edges {
		if v.a == node || v.b == node {
			if found {
				return false
			}
			found = true
		}
	}
	if !found {
		fmt.Printf("isIsolated didn't find node: %s\n", node)
	}
	return false
}

func (g *graph) rmEdge(e tuple) {
	for i, v := range g.edges {
		if (v.a == e.a && v.b == e.b) || (v.a == e.b && v.b == e.a) {
			g.edges = append(g.edges[:i], g.edges[i+1:]...)
			return
		}
	}
	log.Fatalf("Invalid use of rmEdge. Edge: %v doesn't exist.\n", e)
}

// Returns an edge that has been confirmed not to be a bridge.
func (g *graph) getNotBridge() *tuple {
	edgesCopy := g.copyEdges()
	for i, start := range edgesCopy {
		edgesMinusStart := append(edgesCopy[:i], edgesCopy[i+1:]...)
		if v := getNotBridgeInner(start.a, start, edgesMinusStart); v != nil {
			return v
		}
		reverseStart := tuple{a: start.b, b: start.a}
		if v := getNotBridgeInner(reverseStart.a, reverseStart, edgesMinusStart); v != nil {
			return v
		}
	}
	return nil
}
func (g *graph) getNotBridgeFromNode(node string) *tuple {
	edgesCopy := g.copyEdges()
	for i, start := range edgesCopy {
		if start.a != node && start.b != node {
			continue
		}
		edgesMinusStart := append(edgesCopy[:i], edgesCopy[i+1:]...)
		if v := getNotBridgeInner(start.a, start, edgesMinusStart); v != nil {
			return v
		}
		reverseStart := tuple{a: start.b, b: start.a}
		if v := getNotBridgeInner(reverseStart.a, reverseStart, edgesMinusStart); v != nil {
			return v
		}
	}
	return nil
}
func getNotBridgeInner(startNode string, startEdge tuple, edges []tuple) *tuple {
	cand := startEdge
	for i := len(edges) - 1; i >= 0; i-- {
		v := edges[i]
		if v.a == cand.b {
			cand = v
			edgesMinusStart := append(edges[:i], edges[i+1:]...) // Remove used edge
			if foundIt := getNotBridgeInner(startNode, cand, edgesMinusStart); foundIt != nil {
				return foundIt
			}
		} else if v.b == cand.b {
			cand = tuple{a: v.b, b: v.a}
			edgesMinusStart := append(edges[:i], edges[i+1:]...) // Remove used edge
			if foundIt := getNotBridgeInner(startNode, cand, edgesMinusStart); foundIt != nil {
				return foundIt
			}
		}
		if cand.b == startNode {
			return &cand
		}
	}

	return nil
}

func (g *graph) copyEdges() []tuple {
	return append([]tuple{}, g.edges[0:]...)
}
