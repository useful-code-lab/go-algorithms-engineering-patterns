package main

import "fmt"

func BFS(start string, graph map[string][]string) {
	visited := map[string]bool{start: true}
	queue := []string{start}

	fmt.Println("Посещённые элементы:")
	if len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		for _, el := range graph[element] {
			if !visited[el] {
				visited[el] = true
				queue = append(queue, el)
				fmt.Printf("%s ", el)
			}
		}
	}
}

func main() {
	Graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "F"},
		"C": {"K", "L"},
	}
	BFS("A", Graph)
}
