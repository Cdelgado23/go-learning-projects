package utils

import . "github.com/cdelgado23/go-learning-projects/astar/pkg/map"

// Generate generates a map with a given width and height
func Generate(width int, height int) [][]*Node {
	nodes := make([][]*Node, width)
	for i := range nodes {
		nodes[i] = make([]*Node, height)
		for j := range nodes[i] {
			nodes[i][j] = &Node{Location: &Point3D{X: float64(i), Y: float64(j)}}
		}
	}
	// set the neighbors of each node
	for i := range nodes {
		for j := range nodes[i] {
			if i > 0 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i-1][j])
			}
			if i < width-1 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i+1][j])
			}
			if j > 0 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i][j-1])
			}
			if j < height-1 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i][j+1])
			}
		}
	}
	return nodes
}
