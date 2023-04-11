package datastructures

import "fmt"

// Graph data structure represented as an adjacency list
type Graph struct {
	AdjList map[int][]int
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}

// DFSRecursive implements depth-first search recursively
func (g *Graph) DFSRecursive(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Printf("%d ", v)

	for _, i := range g.AdjList[v] {
		if !visited[i] {
			g.DFSRecursive(i, visited)
		}
	}
}
