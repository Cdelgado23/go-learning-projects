package pathfinding

import "github.com/cdelgado23/go-learning-projects/astar/pkg/map"

type vertex struct {
	node   *_map.Node
	cost   float64
	f      float64
	g      float64
	parent *vertex
}

// AStar finds the shortest path between start and end using the A* algorithm
func AStar(start, end *_map.Node) (path []*_map.Node, found bool) {
	// open closed lists
	openList := make([]*vertex, 0)
	closedList := make([]*vertex, 0)

	//move vertex from open list to closed list
	popVertex := func(vertexIndex int) {
		v := openList[vertexIndex]
		closedList = append(closedList, v)
		openList[vertexIndex] = openList[len(openList)-1]
		openList = openList[:len(openList)-1]
	}

	// add the start node to the open list
	openList = append(openList, &vertex{node: start, f: 0, g: 0})

	// loop until the open list is empty
	for len(openList) > 0 {
		// get the node with the lowest f value
		current, index := getLowestF(openList)
		popVertex(index)

		// if the current node is the end node, we have found the path
		if current.node == end {
			path = make([]*_map.Node, 0)
			for current != nil {
				path = append([]*_map.Node{current.node}, path...)
				current = current.parent
			}
			return path, true
		}
		// get the neighbors of the current node
		for _, neighbor := range current.node.Neighbors {
			// create a new vertex for the neighbor
			v := &vertex{node: neighbor, parent: current}
			v.g = current.g + neighbor.Location.Distance(current.node.Location)
			v.f = v.g + neighbor.Location.Distance(end.Location)

			// if the neighbor is in the closed list, skip it
			if _, ok := findVertex(closedList, neighbor); ok {
				continue
			}

			// if the neighbor is in the open list, check if the new path is better
			if closedV, ok := findVertex(openList, neighbor); ok {
				if v.g < closedV.g {
					closedV.g = v.g
					closedV.f = v.f
					closedV.parent = v.parent
				}
			} else {
				openList = append(openList, v)
			}
		}
	}
	return nil, false
}

// check if a vertex is in the list
func findVertex(list []*vertex, node *_map.Node) (v *vertex, ok bool) {
	for _, v := range list {
		if v.node == node {
			return v, true
		}
	}
	return nil, false
}

// get the node with the lowes f value
func getLowestF(openList []*vertex) (vertex *vertex, index int) {
	current := openList[0]
	currentIndex := 0
	for i, v := range openList {
		if v.f < current.f {
			current = v
			currentIndex = i
		}
	}
	return current, currentIndex
}
