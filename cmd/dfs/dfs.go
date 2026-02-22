package main

import "fmt"

func DFS(graph map[string][]string, current string, visited map[string]bool) {
	visited[current] = true

	fmt.Printf("%s ", current)

	for _, child := range graph[current] {
		if !visited[child] {
			DFS(graph, child, visited)
		}
	}
}

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"E": {"K", "L", "M"},
	}
	visited := map[string]bool{}

	DFS(graph, "A", visited)
}
