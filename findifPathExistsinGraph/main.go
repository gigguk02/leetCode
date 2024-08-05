package main

import (
	"fmt"
)

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {2, 5}, {5, 4}, {4, 3}}
	n := 6
	source := 0
	destination := 3
	fmt.Println(validPath(n, edges, source, destination))
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)

	}
	visited := make([]bool, n)
	return dfs(graph, source, destination, visited)
}

func dfs(graph [][]int, current int, destination int, visited []bool) bool {
	if current == destination {
		return true
	}
	visited[current] = true
	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			if dfs(graph, neighbor, destination, visited) {
				return true
			}

		}
	}
	return false

}
