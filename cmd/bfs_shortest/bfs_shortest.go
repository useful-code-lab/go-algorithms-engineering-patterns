package main

import "fmt"

func BFS(graph map[string][]string, start string, goal string) []string {
	visited := map[string]bool{start: true}
	queue := []string{start}
	parent := map[string]string{start: ""}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == goal {
			return collectPath(parent, goal)
		}

		for _, child := range graph[current] {
			if !visited[child] {
				parent[child] = current
				visited[child] = true
				queue = append(queue, child)
			}
		}
	}

	return nil
}

func collectPath(parent map[string]string, goal string) []string {
	path := []string{}
	for current := goal; current != ""; current = parent[current] {
		path = append([]string{current}, path...)
	}
	return path
}

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "F"},
		"D": {"K", "L"},
		"F": {"S", "T"},
	}

	result := BFS(graph, "A", "L")
	fmt.Print(result)
}
