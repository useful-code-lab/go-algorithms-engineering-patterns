package main

func BFS(graph map[string][]string) {

}

func main() {
	Graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "F"},
		"C": {"K", "L"},
	}
	BFS(Graph)
}
