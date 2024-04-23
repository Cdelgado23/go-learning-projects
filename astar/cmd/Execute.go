package main

import (
	"fmt"
	. "github.com/cdelgado23/go-learning-projects/astar/pkg/pathfinding"
	. "github.com/cdelgado23/go-learning-projects/astar/pkg/utils"
)

// Generate a grid of nodes and executes the A* algorithm to find the shortest path between two nodes
func main() {
	// find the shortest path between the top-left and bottom-right nodes
	nodes := Generate(10, 50)
	start := nodes[2][2]
	end := nodes[9][40]
	path, found := AStar(start, end)
	if !found {
		fmt.Println("No path found")
	}
	Draw(nodes, path)
}
