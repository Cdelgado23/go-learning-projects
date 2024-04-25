package path

import (
	"fmt"
	astar "github.com/cdelgado23/go-learning-projects/astar/pkg"
)

type Location struct {
	X, Y, Z float64
}

func NewLocation(x, y, z float64) Location {
	return Location{X: x, Y: y, Z: z}
}

func newLocation(node astar.Node) Location {
	return Location{X: node.Location.X, Y: node.Location.Y, Z: node.Location.Z}
}

type PathFinderService struct {
	nodes [][]*astar.Node
}

func NewPathFinderService(height int, width int) PathFinderService {
	nodes := astar.Generate(height, width)
	return PathFinderService{nodes: nodes}
}

func (p *PathFinderService) PathFromTo(start Location, end Location) []Location {
	startNode, err := p.NodeFromLocation(start)
	if err != nil {
		fmt.Println("error", err)
		return []Location{}
	}
	endNode, err := p.NodeFromLocation(end)
	if err != nil {
		fmt.Println("error", err)
		return []Location{}
	}

	path, _ := astar.AStar(startNode, endNode)

	var locations []Location
	for _, node := range path {
		locations = append(locations, newLocation(*node))
	}

	return locations
}

// NodeFromLocation Node from Location, finding the closest node in the nodes array
// There is a tolerance of +-0.02 for distance between the location and the node
func (p *PathFinderService) NodeFromLocation(location Location) (*astar.Node, error) {
	for _, row := range p.nodes {
		for _, node := range row {
			distance := node.Location.Distance(&astar.Point3D{X: location.X, Y: location.Y, Z: location.Z})
			if distance <= 0.02 {
				return node, nil
			}
		}
	}
	return nil, fmt.Errorf("node not found for location %v", location)
}
