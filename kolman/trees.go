package kolman

import (
	"log"
	"math"
)

type treeMatrix struct {
	m      map[string][]int
	keyIdx map[string]int
}

func (x treeMatrix) spanningTreeMatrixPrims(rootKey string) (result []string) {
	// This should be easily modified for the weighted version. Start by replacing 1's in the matrix with the weight.
	mat := x.m
	keyIdx := x.keyIdx
	result = append(result, rootKey)
	for len(mat) > 1 {

		var mergeKey string
		mergeColIdx := 0

		for key, v := range mat {
			if key == rootKey {
				continue
			}
			if v[0] == 1 {
				mergeKey = key
				mergeColIdx = keyIdx[key]
				break
			}
		}
		result = append(result, mergeKey)
		rootRow := mat[rootKey]
		for i, v := range mat[mergeKey] {
			// Join rows
			if rootRow[i] == 1 || v == 1 {
				mat[rootKey][i] = 1
			}

			// Join cols
			for key, row := range mat {
				if row[mergeColIdx] == 1 || row[0] == 1 {
					mat[key][0] = 1
				}
			}
		}
		mat[rootKey][0] = 0

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
	a    string
	b    string
	dist float64
}

type neighbor struct {
	name     string
	distance float64
}

type treePathGetter struct {
	rootName string
	data     []wrelation
	visited  map[string]string
	distance float64
	// requiredEdges map[string]string // This is an exercise in the book but let's skip it.
}

func (x *treePathGetter) minMaxDistance(min bool) {
	x.visited = make(map[string]string)
	x.visited[x.rootName] = "root"
	for len(x.data) > 0 {
		next, nextIdx, from := x.getNeighbor(min)
		if nextIdx == -1 {
			log.Fatalf("didn't find closest neighbor")
		}
		x.visited[next.name] = from
		x.distance += next.distance
		for i := len(x.data) - 1; i >= 0; i-- {
			if _, ok := x.visited[x.data[i].b]; ok {
				x.data = append(x.data[:i], x.data[i+1:]...) // rm relation to visited
				continue
			}
		}
	}
	x.distance = math.Round(x.distance*100) / 100
}

func (x *treePathGetter) getNeighbor(closest bool) (neighbor, int, string) {
	var nbr neighbor
	neighborIdx := -1
	from := ""
	for i, v := range x.data {
		if _, ok := x.visited[v.a]; ok {
			if _, ok := x.visited[v.b]; !ok {
				candidate := neighbor{name: v.b, distance: v.dist}
				if nbr.name == "" {
					from = v.a
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
					from = v.a
				}
			}
		}
	}
	return nbr, neighborIdx, from
}
