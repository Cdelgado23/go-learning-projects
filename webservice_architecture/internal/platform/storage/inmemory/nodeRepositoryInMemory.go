package inmemory

import (
	astar "github.com/cdelgado23/go-learning-projects/astar/pkg"
	node "github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/node/domain"
	"github.com/google/uuid"
)

// NodeRepositoryInMemory Implements node repository interface
type NodeRepositoryInMemory struct {
	nodes map[string]*astar.Node
}

//NewInMemoryNodeRepository creates a new in-memory node repository
func NewInMemoryNodeRepository(width, height int) NodeRepositoryInMemory {
	nodes := astar.Generate(width, height)
	nodeMap := make(map[string]*astar.Node)
	for _, row := range nodes {
		for _, n := range row {
			nodeMap[uuid.NewString()] = n
		}
	}
	return NodeRepositoryInMemory{nodes: nodeMap}
}

func (repo NodeRepositoryInMemory) GetNodeByLocation(l astar.Point3D) (*astar.Node, error) {
	//Iterate over all nodes
	for _, n := range repo.nodes {
		if node.NodeContains(n, l) {
			return n, nil
		}
	}
	return &astar.Node{}, node.ErrNodeNotFound
}
