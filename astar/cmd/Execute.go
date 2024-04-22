package main

import (
	"fmt"
	. "github.com/cdelgado23/go-learning-projects/astar/pkg/pathfinding"
	. "github.com/cdelgado23/go-learning-projects/astar/pkg/utils"
)

// Generate a grid of nodes and executes the A* algorithm to find the shortest path between two nodes
func main() {
	// find the shortest path between the top-left and bottom-right nodes
	nodes := Generate(10, 10)
	start := nodes[8][9]
	end := nodes[2][2]
	path, found := AStar(start, end)
	if !found {
		fmt.Println("No path found")
	}
	Draw(nodes, path)
}
