package ut

import (
	"fmt"
	"log"
	"slices"
)

// func (g SimpleGraph) BuildDFSTree(startNode string) SimpleGraph {
// 	tree := NewSimpleGraph([]Tuple{})
// 	marks := []string{}
// 	path := g.DFS(startNode)
// 	// return g.buildDFSTree(startNode, []string{}, &tree)
// 	var buildDFSTree func(node string)
// 	buildDFSTree = func(node string) {
// 		marks = append(marks, node)
// 		fmt.Printf("%v\t\t%v\n", path, marks)
// 		tree.Print()
// 		fmt.Println()

// 		for i := 0; i+2 <= len(path); i++ {
// 			v, neighbor := path[i], path[i+1]
// 			if node == v && !slices.Contains(marks, neighbor) {
// 				tree.AddEdge(node, neighbor)
// 				buildDFSTree(neighbor)
// 			}
// 		}
// 	}
// 	buildDFSTree(startNode)
// 	return tree
// }
// func (g SimpleGraph) buildDFSTree_(node string, marks []string, tree *SimpleGraph) SimpleGraph {
// 	path := []string{}
// 	for _, v := range g.DFS(node) {
// 		if !slices.Contains(marks, v) {
// 			path = append(path, v)
// 		}
// 	}

// 	marks = append(marks, node)
// 	fmt.Printf("%v\t\t%v\n", path, marks)
// 	tree.Print()
// 	fmt.Println()
// 	for _, node2 := range path {
// 		if !slices.Contains(marks, node2) {
// 			tree.AddEdge(node, node2)
// 			g.buildDFSTree_(node2, marks, tree)
// 		}
// 	}
// 	return *tree
// }

func (g *SimpleGraph) Print() {
	marks := []string{}
	for node, neighbors := range *g {
		for key := range neighbors {
			if !slices.Contains(marks, key) {
				fmt.Printf("%s %v\t", node, key)
			}
		}
		marks = append(marks, node)
	}
}
func (g *SimpleGraph) AddEdge(u, v string) {
	if _, ok := (*g)[u]; !ok {
		(*g)[u] = make(map[string]struct{})
	}
	if (*g)[v] == nil {
		(*g)[v] = make(map[string]struct{})
	}
	(*g)[u][v] = struct{}{}
	(*g)[v][u] = struct{}{}
}
func (g SimpleGraph) IsDFSTree(startNode string, candidate SimpleGraph) bool {
	return g.isDFSTree(startNode, candidate, []string{})
}
func (g SimpleGraph) isDFSTree(node string, candidate SimpleGraph, marks []string) bool {
	marks = append(marks, node)
	if neighbors, ok := candidate[node]; !ok {
		log.Fatalf("SimpleGraph (tree candidate) node %s isn't there", node)
	} else {
		for neighbor := range neighbors {
			if slices.Contains(marks, neighbor) {
				continue
			}
			return g.isDFSTree(neighbor, candidate, marks)
		}
	}
	if neighbors, ok := g[node]; !ok {
		log.Fatalf("SimpleGraph node %s isn't there", node)
	} else {
		for neighbor := range neighbors {
			if !slices.Contains(marks, neighbor) {
				return false
			}
		}
	}
	return true
}

func (g SimpleGraph) DFSRecursive(node string, visited map[string]bool, result *[]Tuple) {
	visited[node] = true
	for neighbor := range g[node] {
		if !visited[neighbor] {
			*result = append(*result, Tuple{node, neighbor})
			g.DFSRecursive(neighbor, visited, result)
		}
	}
}

func (g *SimpleGraph) DFS(startNode string) []Tuple {
	visited := make(map[string]bool)
	result := []Tuple{}
	g.DFSRecursive(startNode, visited, &result)
	return result
}

type SimpleGraph map[string]map[string]struct{} // Using maps does not maintain order and adding new edges we have to add/remove the edge 'both' ways.

// NewSimpleGraph creates an undirected connected graph with no parallel edges.
func NewSimpleGraph(edges []Tuple) SimpleGraph {
	g := make(SimpleGraph)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if g[a] == nil {
			g[a] = make(map[string]struct{})
		}
		if g[b] == nil {
			g[b] = make(map[string]struct{})
		}
		g[a][b] = struct{}{}
		g[b][a] = struct{}{}
	}
	return g
}

func (g SimpleGraph) Head() string {
	minKey := ""
	for key := range g {
		if minKey == "" {
			minKey = key
		} else {
			if minKey > key {
				minKey = key
			}
		}
		break
	}
	if minKey == "" {
		log.Fatal("bad SimpleGraph key")
	}
	return minKey
}
