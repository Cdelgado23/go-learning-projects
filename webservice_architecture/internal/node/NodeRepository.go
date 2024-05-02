package node

import (
	"errors"
	astar "github.com/cdelgado23/go-learning-projects/astar/pkg"
)

//Contains indicates weather or not a point is contained in the node
func NodeContains(n *astar.Node, p astar.Point3D) bool {
	return n.Location.Distance(&p) <= 0.02
}

type NodeRepository interface {
	GetNodeByLocation(d astar.Point3D) (astar.Node, error)
}

var ErrNodeNotFound = errors.New("node not found")
