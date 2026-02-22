package main

func BFS(start string, graph map[string][]string) {
	startNode = graph[start]
	visited := map[string]bool{start: true,}
	queue := []string{start,}

}

func main() {
	Graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "F"},
		"C": {"K", "L"},
	}
	BFS("A", Graph)
}
