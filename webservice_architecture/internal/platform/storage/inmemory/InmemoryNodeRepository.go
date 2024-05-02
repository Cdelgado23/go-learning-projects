package inmemory

import (
	astar "github.com/cdelgado23/go-learning-projects/astar/pkg"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/node"
	"github.com/google/uuid"
)

// InmemoryNodeRepository Implements node repository interface
type InmemoryNodeRepository struct {
	nodes map[string]astar.Node
}

//NewInmemoryNodeRepository creates a new in-memory node repository
func NewInmemoryNodeRepository(width, height int) InmemoryNodeRepository {
	nodes := astar.Generate(width, height)
	nodeMap := make(map[string]astar.Node)
	for _, row := range nodes {
		for _, n := range row {
			nodeMap[uuid.NewString()] = *n
		}
	}
	return InmemoryNodeRepository{nodes: nodeMap}
}

func (repo InmemoryNodeRepository) GetNodeByLocation(x float64, y float64, z float64) (astar.Node, error) {
	//Iterate over all nodes
	for _, n := range repo.nodes {
		if node.NodeContains(&n, astar.Point3D{X: x, Y: y, Z: z}) {
			return n, nil
		}
	}
	return astar.Node{}, node.ErrNodeNotFound
}
