package astar

// Generate generates a map with a given width and height
func Generate(height int, width int) [][]*Node {
	nodes := make([][]*Node, height)
	for i := range nodes {
		nodes[i] = make([]*Node, width)
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
			if i < height-1 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i+1][j])
			}
			if j > 0 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i][j-1])
			}
			if j < width-1 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i][j+1])
			}
		}
	}
	return nodes
}
