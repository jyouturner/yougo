package yougo

import (
	"fmt"
	"testing"
)

func Test_Dfs(t *testing.T) {
	// create a graph
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D", "E"}
	graph["C"] = []string{"F"}
	graph["D"] = []string{}
	graph["E"] = []string{"F"}
	graph["F"] = []string{}

	// create a visited map
	visited := make(map[string]bool)

	// call dfs
	dfs(graph, "A", visited)

	// print visited map
	for k, v := range visited {
		fmt.Printf("level %s: %v	", k, v)
	}
}

func dfs(graph map[string][]string, node string, visited map[string]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	for _, n := range graph[node] {
		dfs(graph, n, visited)
	}
}

func Test_Dfs_Wfs(t *testing.T) {
	// create a graph
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"D", "E"}
	graph["C"] = []string{"F"}
	graph["D"] = []string{}
	graph["E"] = []string{"F"}
	graph["F"] = []string{}

	//Depth First Search iterate and print the nodes of the graph
	printDfs(graph, "A")
	//
	fmt.Println()
	fmt.Printf("%v ", "A")
	printWfs(graph, "A")

}

func printDfs(graph map[string][]string, parent string) {
	fmt.Printf("%v ", parent)
	//fmt.Println("checking children of ", parent)
	children, ok := graph[parent]
	if !ok {
		//fmt.Println("no child, edge")
		return
	}

	for _, child := range children {
		printDfs(graph, child)
	}
}

func printWfs(graph map[string][]string, parent string) {

	//fmt.Println("checking children of ", parent)
	children, ok := graph[parent]
	if !ok {
		//fmt.Println("no child, edge")
		return
	}

	for _, child := range children {
		fmt.Printf("%v ", child)
	}
	for _, child := range children {
		printWfs(graph, child)
	}
}
