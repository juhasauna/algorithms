package kolman

import (
	"log"
	"math"
	"sort"
)

type treeMatrix struct {
	m        map[string][]int
	keyIdx   map[string]int
	distance int
	rootName string
	visited  map[string]string
}

func (x treeMatrix) spanningTreeMatrixPrims() (result []string) {
	mat := x.m
	keyIdx := x.keyIdx
	result = append(result, x.rootName)
	for len(mat) > 1 {

		mergeKey := getMatrixSmallestWeightPosition(mat)
		mergeColIdx := keyIdx[mergeKey]

		for key, v := range mat {
			if key == x.rootName {
				continue
			}
			if v[0] == 1 {
				mergeKey = key
				mergeColIdx = keyIdx[key]
				break
			}
		}
		result = append(result, mergeKey)
		rootRow := mat[x.rootName]
		for i, v := range mat[mergeKey] {
			// Join rows
			if rootRow[i] > v {
				mat[x.rootName][i] = rootRow[i]
			} else {
				mat[x.rootName][i] = v
			}

			// Join cols
			for key, row := range mat {
				if row[mergeColIdx] > row[0] {
					mat[key][0] = 1
				} else {
					mat[key][0] = row[0]
				}
			}
		}
		mat = zeroeDiagonal(mat, keyIdx)

		delete(mat, mergeKey)
		mat = rmCol(mergeColIdx, mat)
		delete(keyIdx, mergeKey)
		for key, v := range keyIdx {
			if v > mergeColIdx {
				keyIdx[key]--
			}
		}
	}

	return result
}

func getMatrixSmallestWeightPosition(mat map[string][]int) (mergeKey string) {
	smallestWeight := 9999
	for key, row := range mat {
		for _, v := range row {
			if v == 0 {
				continue
			}
			if v < smallestWeight {
				smallestWeight = v
				mergeKey = key
			}
		}
	}
	return mergeKey
}

func zeroeDiagonal(mat map[string][]int, keyIdx map[string]int) map[string][]int {
	result := make(map[string][]int)
	for key, row := range mat {
		row[keyIdx[key]] = 0
		result[key] = row
	}
	return result
}

func (x treeMatrix) minMaxDistance(min bool) {
	// Matrix version of Prim's Algorithm. This is not done. Seems unnecessarily complicated compared to the non-matrix version. I think it's better to just convert the matrix into relation form and use the non-matrix version.
	mat := x.m
	keyIdx := x.keyIdx
	x.visited = make(map[string]string)
	x.visited[x.rootName] = "root"
	for len(mat) > 1 {
		mergeKey := getMatrixSmallestWeightPosition(mat)
		mergeColIdx := keyIdx[mergeKey]
		x.visited[mergeKey] = "?"
		rootRow := mat[x.rootName]
		for i, v := range mat[mergeKey] {
			// Join rows
			if rootRow[i] > v {
				mat[x.rootName][i] = rootRow[i]
			} else {
				mat[x.rootName][i] = v
			}

			// Join cols
			for key, row := range mat {
				if row[mergeColIdx] > row[0] {
					mat[key][0] = 1
				} else {
					mat[key][0] = row[0]
				}
			}
		}
		mat = zeroeDiagonal(mat, keyIdx)

		delete(mat, mergeKey)
		mat = rmCol(mergeColIdx, mat)
		delete(keyIdx, mergeKey)
		for key, v := range keyIdx {
			if v > mergeColIdx {
				keyIdx[key]--
			}
		}
	}

}

func rmCol(j int, mat map[string][]int) map[string][]int {
	for i, row := range mat {
		row = append(row[:j], row[j+1:]...)
		mat[i] = row
	}
	return mat
}

type nodes = map[string]set

// Finding a spanning tree for a symmetric, connected relation. N.B. there can be multiple spanning trees per relation!
func spanningTreePrimsAlg(nodes nodes, rootName string) []string {
	rootNode := nodes[rootName]
	result := []string{rootName}
	for len(nodes) > 1 {
		nextName := rootNode.getFirst()
		result = append(result, nextName)
		nextNode := nodes[nextName]
		delete(nextNode, rootName)
		for name := range nextNode {
			rootNode.add(name)
		}
		for _, v := range nodes {
			delete(v, nextName)
		}
		delete(nodes, nextName) // Delete row
	}
	return result
}

type treePath struct {
	distance float64
	visited  map[string]string
}

type wrelation struct {
	from string
	to   string
	dist float64
}

type neighbor struct {
	name     string
	distance float64
}

type treePathGetter struct {
	rootName  string
	relations []wrelation
	visited   map[string]string
	distance  float64
	edges     []wrelation // used in kruskal
}

func (x *treePathGetter) minMaxDistance(min bool) {
	x.visited = make(map[string]string)
	x.visited[x.rootName] = "root"
	for len(x.relations) > 0 {
		next, nextIdx, from := x.getNeighbor(min)
		if nextIdx == -1 {
			log.Fatalf("didn't find closest neighbor")
		}
		x.visited[next.name] = from
		x.distance += next.distance
		for i := len(x.relations) - 1; i >= 0; i-- {
			if _, ok := x.visited[x.relations[i].to]; ok {
				x.relations = append(x.relations[:i], x.relations[i+1:]...) // rm relation to visited
			}
		}
	}
	x.distance = math.Round(x.distance*100) / 100
}
func (x *treePathGetter) getMinEdge() (result wrelation) {
	var minEdge float64 = 999999.0
	// for i := len(x.relations) - 1; i >= 0; i-- {
	for _, v := range x.relations {
		// v := x.relations[i]
		if v.dist < minEdge {
			if x.isUsedEdge(v) {
				continue
			}
			if isCyclic(x.edges, v, v.from, v.to) {
				continue
			}
			minEdge = v.dist
			result = v
		}
	}
	return result
}
func isCyclic(edges []wrelation, r wrelation, to, from string) bool {
	if r.to == to || r.from == from {
		return true
	}
	if len(edges) == 0 {
		return false
	}
	var copyEdges []wrelation
	for _, v := range edges {
		copyEdges = append(copyEdges, v)
	}

	for i := len(copyEdges) - 1; i >= 0; i-- {
		v := copyEdges[i]
		if v.to == r.from || v.from == r.to {
			copyEdges = append(copyEdges[:i], copyEdges[i+1:]...)
			if isCyclic(copyEdges, v, to, from) {
				return true
			}
		}
	}
	return false
}

func (x *treePathGetter) isUsedEdge(r wrelation) bool {
	for _, v := range x.edges {
		if v.to == r.to && v.from == r.from {
			return true
		}
		if v.from == r.to && v.to == r.from {
			return true
		}
	}
	return false
}

func (x *treePathGetter) rmUsedEdges() {
	for i := len(x.relations) - 1; i >= 0; i-- {
		if from, okTo := x.visited[x.relations[i].to]; okTo {
			if _, okFrom := x.visited[from]; okFrom {
				x.relations = append(x.relations[:i], x.relations[i+1:]...)
			}
		}
	}
}

func (x *treePathGetter) countNodes() int {
	m := make(map[string]struct{})
	for _, relation := range x.relations {
		m[relation.from] = struct{}{}
		m[relation.to] = struct{}{}
	}
	return len(m)
}

func (x *treePathGetter) kruskal(min bool) {
	// This doesn't always work. It fails to produce a connected tree with some inputs.
	// TODO: Make a simplified version that takes doesn't include directions. I.e. the input is represented as bidirectional. The current version is complicated by having a separate relation both ways.
	x.visited = make(map[string]string)
	minEdge := x.getMinEdge()
	x.edges = append(x.edges, minEdge)
	x.visited[minEdge.to] = minEdge.from
	x.distance = minEdge.dist
	for len(x.edges) < x.countNodes()-1 {
		minEdge := x.getMinEdge()
		x.visited[minEdge.to] = minEdge.from
		x.edges = append(x.edges, minEdge)
		x.distance += minEdge.dist
	}
	x.distance = math.Round(x.distance*100) / 100
}

func (x *treePathGetter) getNeighbor(closest bool) (neighbor, int, string) {
	var nbr neighbor
	neighborIdx := -1
	from := ""
	for i, v := range x.relations {
		if _, ok := x.visited[v.from]; ok {
			if _, ok := x.visited[v.to]; !ok {
				candidate := neighbor{name: v.to, distance: v.dist}
				if nbr.name == "" {
					from = v.from
					nbr = candidate
					neighborIdx = i
					continue
				}
				updateCandidate := func() bool {
					if closest {
						return nbr.distance > candidate.distance
					}
					return nbr.distance < candidate.distance
				}

				if updateCandidate() {
					nbr = candidate
					neighborIdx = i
					from = v.from
				}
			}
		}
	}
	return nbr, neighborIdx, from
}

type treePython struct {
	u int
	v int
	w int
}

type kruskalPython struct {
	nodes int
	graph []treePython
}

// https://www.pythonpool.com/kruskals-algorithm-python/
func (k *kruskalPython) kruskal() (result []treePython) {
	// Haven't the slightest clue of how this works. But it seems to work.
	i, e := 0, 0
	sort.Slice(k.graph, func(i2, j int) bool {
		return k.graph[i2].w < k.graph[j].w
	})
	parent := []int{}
	rank := []int{}
	for node := 0; node < k.nodes; node++ {
		parent = append(parent, node)
		rank = append(rank, 0)
	}
	for e < k.nodes-1 {
		treeItem := k.graph[i]
		i++
		x := k.search(parent, treeItem.u)
		y := k.search(parent, treeItem.v)
		if x != y {
			e++
			result = append(result, treeItem)
			k.union(parent, rank, x, y)
		}
	}
	return result
}
func (k *kruskalPython) search(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return k.search(parent, parent[i])
}
func (k *kruskalPython) union(parent, rank []int, x, y int) {
	xroot := k.search(parent, x)
	yroot := k.search(parent, y)
	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}
