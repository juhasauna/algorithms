package ut

import (
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
)

type Graph struct {
	Nodes []string // node index → name
	Adj   [][]int  // adjacency matrix
	Edges []Edge   // This has the same information as Adj, but sometimes more convenient to work with.
	Name  string
}

type sccMemo struct {
	dfsNumber int
	high      int
	component int
}

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch7_notes_b.pdf
// A directed graph is strongly connected if there is a directed path from every vertex to every other vertex.
// From Manber p.248 pseudocode
func (g Graph) TarjansStronglyConnectedComponents() map[string]int {
	memos := make(map[string]sccMemo, g.CountNodes())
	dfsN := g.CountNodes()
	stack := Stack[string]{}
	currentComponent := 0
	var scc func(node string)
	scc = func(node string) {
		memos[node] = sccMemo{dfsNumber: dfsN, high: dfsN}
		dfsN--
		stack.Push(node)
		for _, e := range g.EdgesFrom(node) {
			if memos[e.to].dfsNumber == 0 {
				scc(e.to)
				temp := memos[node]
				temp.high = Max(memos[node].high, memos[e.to].high)
				memos[node] = temp
			} else if memos[e.to].dfsNumber > memos[node].dfsNumber && memos[e.to].component == 0 {
				temp := memos[node]
				temp.high = Max(memos[node].high, memos[e.to].dfsNumber)
				memos[node] = temp
			}
			NTUSCCEXPrinter(memos)
		}
		if memos[node].high == memos[node].dfsNumber {
			currentComponent++

			x := ""
			for {
				x = stack.Pop()
				temp := memos[x]
				temp.component = currentComponent
				memos[x] = temp
				if x == node {
					break
				}
			}
		}
	}
	for {
		stop := true
		// Check that the solution does not depend on the order of the nodes.
		// ShuffleSlice(g.Nodes)
		// fmt.Println(g.Nodes)
		for _, node := range g.Nodes {
			if memos[node].dfsNumber == 0 {
				scc(node)
				stop = false
			}
		}
		if stop {
			break
		}
	}
	result := make(map[string]int)
	for key, v := range memos {
		result[key] = v.component
	}
	return result
}

func NTUSCCEXPrinter(memos map[string]sccMemo) {
	// See file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch7_notes_b.pdf
	debugHelper := []string{}
	for key, v := range memos {
		debugHelper = append(debugHelper, fmt.Sprintf("%s: %d,", key, v.high))
	}
	slices.Sort(debugHelper)
	fmt.Println(debugHelper)
}

func (g Graph) BellmanFord() {
	// Shortest path algorithm
	// 	Handles negative weights

	// 	Detects negative cycles

	// 	Time: O(VE)
}

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw7_s.pdf
func (g Graph) FindLongestPath() []string {
	type memory struct {
		weight int
		parent string
	}
	lengths := make(map[string]memory)
	for _, n := range g.Nodes {
		lengths[n] = memory{}
	}

	for _, node := range topologicalSortResultToSlice(g.TopologicalSorting()) {
		for _, e := range g.EdgesFrom(node) {
			if lengths[e.to].weight < lengths[node].weight+1 {
				newMemory := memory{lengths[node].weight + 1, node}
				lengths[e.to] = newMemory
			}
		}
	}
	maxLength := 0
	maxLengthNode := ""
	for key, v := range lengths {
		if v.weight > maxLength {
			maxLength = v.weight
			maxLengthNode = key
		}
	}
	parent := lengths[maxLengthNode].parent
	longestPath := []string{maxLengthNode}
	for parent != "" {
		longestPath = append(longestPath, parent)
		parent = lengths[parent].parent
	}
	slices.Reverse(longestPath)
	return longestPath
}

// Floyd's algorithm AKA Floyd-Warshall. Negative edge weights allowed, but no negative cycles.
func (g Graph) AllPairsShortestPaths() [][]int {
	g.PrintMatrix()
	size := g.CountNodes()
	result := make([][]int, size)
	for i := range size {
		result[i] = make([]int, size)
	}
	for i := range size {
		for j := range size {
			if i != j {
				if g.Adj[i][j] == 0 {
					result[i][j] = INF
				} else {
					result[i][j] = g.Adj[i][j]
				}
			} else {
				result[i][j] = 0
			}
		}
	}

	for m := range size {
		for x := range size {
			for y := range size {
				w := result[x][m] + result[m][y]
				if w < result[x][y] {
					result[x][y] = w
				}
			}
		}
	}
	return result
}

// file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/alg2024hw8_s.pdf
func (g Graph) Kruskal() []UndirectedEdge {
	treeEdges := make(map[int]Edge)
	g.SortEdgesAscendingByWeight()
	seq := []int{}
	for _, v := range g.Nodes {
		runes := []rune(v)
		if len(runes) != 1 {
			log.Fatalf("use a single character node name please (not %s)", v)
		}
		seq = append(seq, int(runes[0]))
	}
	dsu := DSU{Balanced: true, UsePathCompression: true}
	dsu.DSUInit(seq)

	for i, e := range g.Edges {
		u, v := int([]rune(e.from)[0]), int([]rune(e.to)[0])
		if dsu.Find(u) != dsu.Find(v) {
			treeEdges[i] = e
			dsu.Union(u, v)
		}
		if len(treeEdges) == len(g.Edges)-1 {
			break
		}
	}

	result := []UndirectedEdge{}
	for _, e := range treeEdges {
		if e.from < e.to {
			result = append(result, UndirectedEdge{e.from, e.to, e.weight})
			continue
		}
		result = append(result, UndirectedEdge{e.to, e.from, e.weight})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].weight != result[j].weight {
			return result[i].weight < result[j].weight
		}
		if result[i].a != result[j].a {
			return result[i].a < result[j].a
		}
		return result[i].b < result[j].b
	})
	return result
}

// Minimum-Cost Spanning Tree. A variant of Prim's algorithm. For simplicity, we assume that the costs are distinct.
func (g Graph) MCST() []UndirectedEdge {
	// MCST is for undirected graph but we are treating a directecd graph as undirected by ignoring the direction of edges.

	treeEdges := make(map[string]Edge)
	marks := []string{}
	for _, v := range g.Nodes {
		treeEdges[v] = Edge{weight: INF}
	}
	minEdge := g.GetMinWeightEdge(marks)
	x := minEdge.from // Either 'from' or 'to' works
	marks = append(marks, x)
	for _, e := range g.NodeEdges(x) {
		z := GetReciprocalNode(x, e)
		treeEdges[z] = e
	}

	for len(marks) < g.CountNodes() {
		minWeight := INF
		minNode := ""
		newTreeEdge := Edge{}
		for key, e := range treeEdges {
			if !slices.Contains(marks, key) {
				if e.weight < minWeight {
					minWeight = e.weight
					minNode = key
					newTreeEdge = e
				}
			}
		}
		if minWeight == INF {
			log.Fatalf("Graph is not connected")
		}
		marks = append(marks, minNode)
		treeEdges[minNode] = newTreeEdge

		for _, e := range g.NodeEdges(minNode) {
			z := GetReciprocalNode(minNode, e)
			if !slices.Contains(marks, z) {
				if e.weight < treeEdges[z].weight {
					treeEdges[z] = e
				}
			}
		}
	}
	result := []UndirectedEdge{}
	for _, e := range treeEdges {
		if e.weight == INF {
			continue
		}
		if e.from < e.to {
			result = append(result, UndirectedEdge{e.from, e.to, e.weight})
			continue
		}
		result = append(result, UndirectedEdge{e.to, e.from, e.weight})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].weight != result[j].weight {
			return result[i].weight < result[j].weight
		}
		if result[i].a != result[j].a {
			return result[i].a < result[j].a
		}
		return result[i].b < result[j].b
	})
	return result
}

type UndirectedEdge struct {
	a      string
	b      string
	weight int
}

func GetReciprocalNode(node string, e Edge) string {
	if e.to == node {
		return e.from
	}
	if e.from != node {
		log.Fatal("GetRecipricalNode error")
	}
	return e.to
}

// For efficient implementation use min-heap instead.
func (g Graph) GetMinWeightEdge(marks []string) Edge {
	minWeight := INF
	result := Edge{}
	for _, e := range g.Edges {
		if !slices.Contains(marks, e.from) {
			if e.weight < minWeight {
				minWeight = e.weight
				result = e
			}
		} else if !slices.Contains(marks, e.to) {
			if e.weight < minWeight {
				minWeight = e.weight
				result = e
			}
		}
	}
	return result
}

// 2023final ex.7: file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/exams/alg2023final.pdf
// Dynamic Programming Single Source Shortest Paths
func (x Graph) DPSSSP(startNode string) map[string]int {
	D := make(map[string]int) // Lenght to the shortest path to key, from startNode with at most i edges.

	for _, node := range x.Nodes {
		D[node] = INF
		if node != startNode {
			for _, e := range x.Edges {
				if e.from == startNode {
					D[e.to] = e.weight
				}
			}
		} else {
			D[startNode] = 0
		}
	}
	for l := 1; l < x.CountNodes()-1; l++ {
		for _, node := range x.Nodes {
			if node != startNode {
				for _, e := range x.EdgesTo(node) {
					if D[e.from]+e.weight < D[e.to] {
						D[e.to] = D[e.from] + e.weight
					}
				}
			}
		}
	}
	for key, v := range D {
		if v == INF {
			delete(D, key)
		}
	}
	return D
}

// See DijkstrasNTUPseudo for a more accurate duplication of the NTU-pseudo code.
// AKA. Single Source Shortest Paths. Dijkstras assumes no negative weights. When the graph is not acyclic, there is no such thing as a topological order.
func (g Graph) Dijkstras(startNode string) map[string]int {
	// The problem with this implementation is that we're inserting duplicate values to the heap. It still works since Extract() always removes the min. But for an efficient implementation we should be updating the heap element if it already exists. The method for this is commonly called 'DecreaseKey()'
	minHeap := NewDijkstraHeap([]*DijkstraNode{{Name: startNode, Weight: 0}})
	results := make(map[string]int)

	for !minHeap.IsEmpty() {
		w := minHeap.Extract()
		if _, ok := results[w.Name]; !ok {
			fmt.Printf("Extracted: %+v\n", *w)
			results[w.Name] = w.Weight
		}
		fmt.Printf("Extracted: %+v\n", *w)
		for _, e := range g.EdgesFrom(w.Name) {
			if _, visited := results[e.to]; !visited {
				fmt.Printf("inserting %s: %d\n", e.to, w.Weight+e.weight)
				kid := &DijkstraNode{Name: e.to, Weight: w.Weight + e.weight}
				if existingNode := minHeap.GetNode(e.to); existingNode != nil {
					if existingNode.Weight > kid.Weight {
						minHeap.DecreaseKey(kid)
					}
				} else {
					minHeap.Insert(kid)
				}
			}
		}
	}
	fmt.Printf("DIJKSTRAS RESULT %+v\n", results)
	return results
}
func (g Graph) DijkstrasSimpleButNotOptimal(startNode string) map[string]int {
	// The problem with this implementation is that we're inserting duplicate values to the heap. It still works since Extract() always removes the min. But for an efficient implementation we should be updating the heap element if it already exists. The method for this is commonly called 'DecreaseKey()'
	minHeap := NewDijkstraHeap([]*DijkstraNode{{Name: startNode, Weight: 0}})
	results := make(map[string]int) // This also functions as "marks".
	for !minHeap.IsEmpty() {
		w := minHeap.Extract()
		if _, ok := results[w.Name]; !ok {
			results[w.Name] = w.Weight
		}
		for _, e := range g.EdgesFrom(w.Name) {
			if _, visited := results[e.to]; !visited {
				// fmt.Printf("inserting %s: %d\n", e.to, w.Weight+e.weight)
				kid := &DijkstraNode{Name: e.to, Weight: w.Weight + e.weight}
				minHeap.Insert(kid)
			}
		}
	}
	fmt.Printf("DIJKSTRAS RESULT %+v\n", results)
	return results
}

// Inefficient method for finding min weight. file:///C:/Users/FIJUSAU/OneDrive%20-%20ABB/courses/Vaihto/TaiwanTech/algorithms_2024_material/slides/ch7_notes_a.pdf
func (x Graph) DijkstrasNTUPseudo(startNode string) map[string]int {
	lengths := make(map[string]int)

	marks := []string{}
	for _, v := range x.Nodes {
		lengths[v] = INF
	}
	lengths[startNode] = 0
	for len(marks) < x.CountNodes() {
		w := FindMinWeight(marks, lengths)
		marks = append(marks, w)
		// fmt.Printf("w: %s, marks: %v\n", w, marks)
		for _, e := range x.EdgesFrom(w) {
			if !slices.Contains(marks, e.to) {
				if lengths[w]+e.weight < lengths[e.to] {
					lengths[e.to] = lengths[w] + e.weight
				}
			}
		}
	}
	for key, v := range lengths {
		if v == INF {
			delete(lengths, key)
		}
	}
	return lengths
}

// Helper for Dijkstras. Efficient implementation would be done with minheap.
func FindMinWeight(marks []string, lengths map[string]int) string {
	min := math.MaxInt
	result := ""
	for key, length := range lengths {
		if !slices.Contains(marks, key) {
			if length < min {
				min = length
				result = key
			}
		}
	}
	return result
}

// Manber p.221
func (x Graph) ImprovedAcyclicShortestPaths(startNode string) map[string]int {
	return nil
}

// Manber p.220
func (x Graph) AcyclicShortestPaths(startNode string) map[string]int {
	topoSort := x.TopologicalSorting()
	sortedNodes := topologicalSortResultToSlice(topoSort)
	sp := make(map[string]int)
	for _, n := range sortedNodes {
		sp[n] = 999
	}

	var recurse func(n int)
	recurse = func(n int) {
		z := sortedNodes[n]
		sortedNodes = slices.Delete(sortedNodes, n, n+1)

		if z != startNode {
			edges := x.EdgesTo(z)
			for _, e := range edges {
				x.RemoveEdge(e)
			}
			recurse(n - 1)
			for _, e := range edges {
				if sp[e.from]+e.weight < sp[z] {
					sp[z] = sp[e.from] + e.weight
				}
			}
		} else {
			sp[z] = 0
		}
	}
	recurse(len(x.Nodes) - 1)
	return sp
}

type TopologicalOrder struct {
	inDegree  int
	outDegree int
	label     int
}

func (x Graph) GetNodeDegrees() map[string]TopologicalOrder {
	m := make(map[string]TopologicalOrder)
	for _, v := range x.Nodes {
		m[v] = TopologicalOrder{}
	}
	for _, e := range x.Edges {
		temp := m[e.from]
		temp.outDegree++
		m[e.from] = temp
		temp = m[e.to]
		temp.inDegree++
		m[e.to] = temp
	}
	return m
}

// TopologicalSorting. ONLY DEFINED FOR DAG's.
func (x *Graph) TopologicalSorting() map[string]TopologicalOrder {
	// Manber Lemma 7.8: A DAG always contains a vertex with indegree 0.
	{
		sinks := x.GetSinks()
		if len(sinks) == 0 {
			log.Fatalf("TopologicalSorting: no sinks in graph")
		}
	}

	queue := x.GetSources()
	if len(queue) == 0 {
		log.Fatalf("TopologicalSorting: no sources in graph")
	}
	G_label := 0
	nodes := x.GetNodeDegrees()
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		G_label++
		temp := nodes[from]
		temp.label = G_label
		nodes[from] = temp
		for _, e := range x.EdgesFrom(from) {
			temp := nodes[e.to]
			temp.inDegree--
			nodes[e.to] = temp
			if temp.inDegree <= 0 {
				queue = append(queue, e.to)
			}
		}
	}
	return nodes
}

func topologicalSortResultToSlice(nodes map[string]TopologicalOrder) []string {
	sortedNodes := make([]string, len(nodes))
	for key, v := range nodes {
		sortedNodes[v.label-1] = key
	}
	return sortedNodes
}

func (x *Graph) BFS(startNode string) []Edge {
	result := []Edge{}
	marks := []string{startNode}
	queue := []string{startNode}
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:] // {first in queue, ..., last in queue}
		for _, e := range x.Edges {
			if e.from == from {
				if !slices.Contains(marks, e.to) {
					marks = append(marks, e.to)
					result = append(result, Edge{from, e.to, e.weight})
					queue = append(queue, e.to)
				}
			}
		}
	}
	return result

}

func (x *Graph) HaveCycle() bool { // NOT TESTED SO PROPABLY NOT CORRECT.
	startNode := x.Edges[0].from
	dfsEdges := x.DFS(startNode)
	visited := make(map[string]struct{})
	for _, v := range dfsEdges {
		visited[v.from] = struct{}{}
	}
	for _, v := range dfsEdges {
		if _, ok := visited[v.to]; ok {
			fmt.Printf("cycle found from %s to itself", v.to)
			return true
		}
	}
	return false
}

func (x *Graph) DFS(startNode string) []Edge {
	result := []Edge{}
	marks := []string{}
	var recurse = func(node string) {}
	recurse = func(from string) {
		marks = append(marks, from)
		for _, e := range x.Edges {
			if e.from == from {
				if !slices.Contains(marks, e.to) {
					result = append(result, Edge{from, e.to, e.weight})
					recurse(e.to)
				}
			}

		}
	}
	recurse(startNode)
	return result
}
func (x *Graph) RemoveEdge(edge Edge) {
	for i, e := range x.Edges {
		if e.from == edge.from && e.to == edge.to {
			x.Edges = slices.Delete(x.Edges, i, i+1)
			u, v := x.Index(e.from), x.Index(e.to)
			x.Adj[u][v] = 0
			return
		}
	}
}

func (x *Graph) AddEdge(edge Edge) {
	if !slices.Contains(x.Nodes, edge.from) {
		log.Fatalf("Cannot add edge from node %s because it's not in the nodes list", edge.from)
	}
	if !slices.Contains(x.Nodes, edge.to) {
		log.Fatalf("Cannot add edge to node %s because it's not in the nodes list", edge.to)
	}
	x.Edges = append(x.Edges, edge)
	x.Adj[x.Index(edge.from)][x.Index(edge.to)] = edge.weight
}

func NewUndirectedGraph(name string, edges_ any, unconnectedNodes ...string) Graph {
	g := NewGraph(name, edges_, unconnectedNodes...)
	for i := range g.CountNodes() {
		for j := range g.CountNodes() {
			if g.Adj[i][j] != 0 {
				g.Adj[j][i] = g.Adj[i][j]
			} else if g.Adj[j][i] != 0 {
				g.Adj[i][j] = g.Adj[j][i]
			}
		}
	}
	directedEdgesCount := len(g.Edges)
	for _, e := range g.Edges {
		g.Edges = append(g.Edges, Edge{from: e.to, to: e.from, weight: e.weight})
	}
	edgeCount := len(g.Edges)
	if 2*directedEdgesCount != edgeCount {
		log.Fatalf("inconsistent edge counts %d/%d", directedEdgesCount, edgeCount)
	}
	return g
}
func NewGraph(name string, edges_ any, unconnectedNodes ...string) Graph {
	var edges []Edge
	switch value := edges_.(type) {
	case []Edge:
		edges = value
	case []Tuple:
		for _, e := range value {
			edges = append(edges, Edge{from: e[0], to: e[1], weight: 1})
		}
	default:
		log.Fatal("Unsupported type")
	}
	g := Graph{Name: name}
	nodes := unconnectedNodes
	for _, e := range edges {
		if !slices.Contains(nodes, e.from) {
			nodes = append(nodes, e.from)
		}
		if !slices.Contains(nodes, e.to) {
			nodes = append(nodes, e.to)
		}
	}
	slices.Sort(nodes)
	g.Nodes = nodes
	adjSize := g.CountNodes()
	adj := make([][]int, adjSize)
	for i := range adj {
		adj[i] = make([]int, adjSize)
	}
	g.Adj = adj
	for _, e := range edges {
		g.AddEdge(e)
	}

	// This is just to make testcases more convenient to work with.
	slices.SortFunc(g.Edges, func(a, b Edge) int {
		if a.from > b.from {
			return 1
		} else if a.from != b.from {
			return -1
		}
		if a.to > b.to {
			return 1
		} else if a.to != b.to {
			return -1
		}
		if a.weight > b.weight {
			return 1
		} else if a.weight != b.weight {
			return -1
		}
		return 0
	})
	return g
}

func (g *Graph) SortEdgesAscendingByWeight() {
	slices.SortFunc(g.Edges, func(a, b Edge) int {
		if a.weight > b.weight {
			return 1
		} else if a.weight != b.weight {
			return -1
		}
		if a.from > b.from {
			return 1
		} else if a.from != b.from {
			return -1
		}
		if a.to > b.to {
			return 1
		} else if a.to != b.to {
			return -1
		}
		return 0
	})
}

func (x Graph) Index(node string) int {
	// For O(1) index lookup, add field nodeIdx map[string]int to type Graph
	// This is O(n)
	if len(x.Nodes) == 0 {
		log.Fatalf("NO NODES IN GRAPH %s -> node %s does not exist", x.Name, node)
	}
	i := slices.Index(x.Nodes, node)
	if i == -1 {
		log.Fatalf("node %s does not exist in graph %s", node, x.Name)
	}
	return i
}
func (x Graph) CountNodes() int {
	return len(x.Nodes)
}

func (x Graph) PrintMatrix() {
	fmt.Printf("  ")
	for k := range x.CountNodes() {
		fmt.Printf("%s ", x.Nodes[k])
	}
	fmt.Printf("\n")

	for i, row := range x.Adj {
		fmt.Printf("%s ", x.Nodes[i])
		for j, weight := range row {

			if j > 0 {
				fmt.Print(" ")
			}

			if weight == 0 {
				fmt.Print("-")
			} else {
				fmt.Printf("%d", weight)
			}
		}
		fmt.Println()
	}
}

func (g Graph) NodeEdges(node string) []Edge {
	result := []Edge{}
	for _, e := range g.Edges {
		if e.from == node || e.to == node {
			result = append(result, e)
		}
	}
	return result
}

func (x *Graph) EdgesFrom(from string) []Edge {
	result := []Edge{}
	for _, e := range x.Edges {
		if e.from == from {
			result = append(result, e)
		}
	}
	return result
}
func (x *Graph) EdgesTo(to string) []Edge {
	result := []Edge{}
	for _, e := range x.Edges {
		if e.to == to {
			result = append(result, e)
		}
	}
	return result
}

func DrawDirectedGraphASCII(edges []Edge) {
	adj := make(map[string][]string)

	// Build adjacency list
	for _, e := range edges {
		adj[e.from] = append(adj[e.from], e.to)
		// Ensure the 'to' node appears even if it has no outgoing edges
		if _, exists := adj[e.to]; !exists {
			adj[e.to] = []string{}
		}
	}

	// Print nodes and their outgoing edges
	for node, targets := range adj {
		fmt.Printf("%s", node)
		if len(targets) > 0 {
			for _, target := range targets {
				fmt.Printf(" --> %s", target)
			}
		}
		fmt.Println()
	}
}

type Edge struct {
	from   string
	to     string
	weight int
}

// Source is a node with indegree zero.
func (x Graph) GetSources() []string {
	sources := x.Nodes
	for _, e := range x.Edges {
		sources = RemoveSliceValue(sources, e.to)
	}
	return sources
}
func (x Graph) GetSinks() []string {
	sinks := x.Nodes
	for _, e := range x.Edges {
		sinks = RemoveSliceValue(sinks, e.from)
	}
	return sinks
}

const INF = math.MaxInt16
