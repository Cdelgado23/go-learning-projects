package inmemory

import (
	"errors"
	astar "github.com/cdelgado23/go-learning-projects/astar/pkg"
	node "github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/node/domain"
	"testing"
)

func TestInMemoryNodeRepository_GetNodeByLocation(t *testing.T) {
	repository := NewInMemoryNodeRepository(5, 5)
	n, err := repository.GetNodeByLocation(astar.Point3D{X: 1, Y: 1})
	if nil == n {
		t.Fatalf("Node was nil, expected not nil. Err %v", err)
	}
}

func TestInMemoryNodeRepository_GetNotExistentNode(t *testing.T) {
	repository := NewInMemoryNodeRepository(5, 5)
	_, err := repository.GetNodeByLocation(astar.Point3D{X: 10, Y: 10})
	if err == nil {
		t.Fatalf("Error was expected, found nil")
	}

	//check if the error is the ErrNodeNotFound defined in NodeRepository
	if !errors.Is(err, node.ErrNodeNotFound) {
		t.Fatalf("Was expecting %v, found %v", node.ErrNodeNotFound, err)
	}

}
