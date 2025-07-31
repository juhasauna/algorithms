package ut

import (
	"fmt"
	"math"
)

type EulerianGraph struct {
	// EulerianGraph is a connected multigraph with all vertices of even degree
	// Allows multiple edges between vertices.
	Adj  map[int]map[int]int
	Name string
}

func (g *EulerianGraph) Copy() *EulerianGraph {
	newAdj := make(map[int]map[int]int)
	for v, neighbors := range g.Adj {
		newAdj[v] = make(map[int]int)
		for u, count := range neighbors {
			newAdj[v][u] = count
		}
	}
	return &EulerianGraph{Adj: newAdj}
}
func (g *EulerianGraph) HasNode(v int) bool {
	_, exists := g.Adj[v]
	return exists
}
func (g *EulerianGraph) RemoveIfIsolated(v int) {
	neighbors, exists := g.Adj[v]
	if !exists {
		return // Node doesn't exist
	}
	if len(neighbors) == 0 {
		delete(g.Adj, v)
	}
}

// IsEulerian checks if the graph is connected and all vertices have even degree.
func (g *EulerianGraph) IsEulerian() bool {
	if len(g.Adj) == 0 {
		return false
	}

	// Check that all vertices have even degree
	for _, neighbors := range g.Adj {
		degree := 0
		for _, edgeCount := range neighbors {
			degree += edgeCount
		}
		if degree%2 != 0 {
			return false
		}
	}

	// Check connectivity
	visited := make(map[int]bool)
	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		for u := range g.Adj[v] {
			if !visited[u] {
				dfs(u)
			}
		}
	}

	// Start DFS from an arbitrary vertex with at least one edge
	var start int
	for v, neighbors := range g.Adj {
		if len(neighbors) > 0 {
			start = v
			break
		}
	}
	dfs(start)

	// Ensure all vertices with at least one edge were visited
	for v, neighbors := range g.Adj {
		if len(neighbors) > 0 && !visited[v] {
			return false
		}
	}

	return true
}

func (g *EulerianGraph) RemoveNode(v int) {
	// Step 1: Remove references to `v` in neighbors
	if neighbors, exists := g.Adj[v]; exists {
		for u := range neighbors {
			delete(g.Adj[u], v)
		}
	}

	// Step 2: Remove the node itself
	delete(g.Adj, v)
}
func (g *EulerianGraph) RemoveCircuit() bool {
	circuit := g.HierholzerFirst()
	if len(circuit) == 0 {
		return true
	}
	for i := 1; i < len(circuit); i++ {
		g.RemoveEdge(circuit[i-1], circuit[i-1])
	}
	return false
}
func (g *EulerianGraph) HierholzerFirst() []int {
	for key := range g.Adj {
		return g.Hierholzer(key)
	}
	return []int{}
}

// Hierholzer find an arbitrary Eulerian circuit
func (g *EulerianGraph) Hierholzer(start int) []int {
	var path []int
	stack := []int{start}

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		if len(g.Adj[u]) == 0 {
			// No more edges from u
			path = append(path, u)
			stack = stack[:len(stack)-1]
		} else {
			// Pick any neighbor v
			for v := range g.Adj[u] {
				g.RemoveEdge(u, v)
				stack = append(stack, v)
				break
			}
		}
	}
	return path
}

func NewEulerianGraph(name string) *EulerianGraph {
	return &EulerianGraph{Adj: make(map[int]map[int]int), Name: name}
}

func (g *EulerianGraph) AddEdge(u, v int) {
	if g.Adj[u] == nil {
		g.Adj[u] = make(map[int]int)
	}
	if g.Adj[v] == nil {
		g.Adj[v] = make(map[int]int)
	}
	g.Adj[u][v]++
	g.Adj[v][u]++ // Undirected
}

func (g *EulerianGraph) EdgeCount() int {
	count := 0
	seen := make(map[[2]int]bool)

	for u, neighbors := range g.Adj {
		for v, multiplicity := range neighbors {
			// Ensure each undirected edge is only counted once
			key := [2]int{u, v}
			if u > v {
				key = [2]int{v, u}
			}
			if !seen[key] {
				count += multiplicity
				seen[key] = true
			}
		}
	}

	return count
}
func (g *EulerianGraph) HasEdges() bool {
	return g.EdgeCount() != 0
}
func (g *EulerianGraph) RemoveEdge(u, v int) {
	if g.Adj[u][v] > 0 {
		g.Adj[u][v]--
		if g.Adj[u][v] == 0 {
			delete(g.Adj[u], v)
		}
	}
	if g.Adj[v][u] > 0 {
		g.Adj[v][u]--
		if g.Adj[v][u] == 0 {
			delete(g.Adj[v], u)
		}
	}
}

func IsValidEulerianCircuit(path []int) error {
	m := make(map[int]int)
	last := math.MaxInt32
	for _, v := range path {
		if last == v {
			return fmt.Errorf("Not valid Eulerian path: there is an edge from %d to itself.", v)
		}
		last = v
		m[v]++
	}
	for key, v := range m {
		if !IsPowerOfTwo(v) {
			return fmt.Errorf("Not valid Eulerian path: %d appears %d times, which is not a power of two.", key, v)
		}
	}
	return nil
}

func (g *EulerianGraph) Print() {
	for u, neighbors := range g.Adj {
		for v, count := range neighbors {
			fmt.Printf("%d --(%d)-- %d\n", u, count, v)
		}
	}
}

func GetEulerianGraphK3() *EulerianGraph {
	g := NewEulerianGraph("K3")
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	return g
}
func GetEulerianGraphK4() *EulerianGraph {

	g := GetEulerianGraphK3()
	g.Name = "K4"
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	return g
}
func GetEulerianGraphK5() *EulerianGraph {
	g := GetEulerianGraphK4()
	g.Name = "K5"
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	return g
}
func GetEulerianGraphK5_a() *EulerianGraph {
	g := GetEulerianGraphK4()
	g.Name = "K5_a"
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.RemoveEdge(0, 1)
	g.RemoveEdge(1, 2)
	g.RemoveEdge(0, 2)
	return g
}
func GetNotEulerianGraphK5_a() *EulerianGraph {
	g := GetEulerianGraphK4()
	g.Name = "K5_a_not_eulerian"
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.RemoveEdge(0, 1)
	return g
}
func GetEulerianGraphBowtie() *EulerianGraph {
	g := GetEulerianGraphK3()
	g.Name = "Bowtie"
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 3)
	return g
}
func GetNotEulerianGraph2by4PartialDiagonal() *EulerianGraph {
	// 0---1---4---5
	// |\ /|   |\ /|
	// | X |   | X |
	// |/ \|   |/ \|
	// 2---3---6---7
	g := GetEulerianGraphK4()
	g.Name = "2by4PartialDiagonal_not_eulerian"
	g.AddEdge(1, 4)
	g.AddEdge(3, 6)

	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(5, 6)
	g.AddEdge(5, 7)
	g.AddEdge(6, 7)
	return g
}
