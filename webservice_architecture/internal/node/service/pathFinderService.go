package node

import (
	"fmt"
	astar "github.com/cdelgado23/go-learning-projects/astar/pkg"
	node "github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/node/domain"
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
	repository node.NodeRepository
}

func NewPathFinderService(repository node.NodeRepository) PathFinderService {
	return PathFinderService{repository: repository}
}

func (p *PathFinderService) PathFromTo(start Location, end Location) []Location {
	startNode, err := p.repository.GetNodeByLocation(locationTo3DPoint(start))
	if err != nil {
		fmt.Println("error", err)
		return []Location{}
	}
	endNode, err := p.repository.GetNodeByLocation(locationTo3DPoint(end))
	if err != nil {
		fmt.Println("error", err)
		return []Location{}
	}

	path, _ := astar.AStar(startNode, endNode)

	var locations []Location
	for _, n := range path {
		locations = append(locations, newLocation(*n))
	}

	return locations
}

func locationTo3DPoint(l Location) astar.Point3D {
	return astar.Point3D{X: l.X, Y: l.Y, Z: l.Z}
}
