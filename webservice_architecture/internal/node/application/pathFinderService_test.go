package node

import (
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/platform/storage/inmemory"
	"math"
	"math/rand"
	"testing"
	"time"
)

//Test pathfinder application path from to
func TestPathFinderService_PathFromTo(t *testing.T) {
	repository := inmemory.NewInMemoryNodeRepository(5, 5)
	//Create a new pathfinder application
	pathFinderService := NewPathFinderService(repository)

	//Create a start and end location
	start := NewLocation(0, 0, 0)
	end := NewLocation(1, 1, 0)

	//Get the path
	path := pathFinderService.PathFromTo(start, end)

	//Check if the path is correct
	if len(path) != 3 {
		t.Errorf("Path length was %d, expected 2", len(path))
	}
	//First element is the start
	if path[0] != start {
		t.Errorf("First element in path was %v, expected %v", path[0], start)
	}
	//Last element is the end
	if path[len(path)-1] != end {
		t.Errorf("Last element in path was %v, expected %v", path[len(path)-1], end)
	}
}

//Test pathfinder when start node is not found
func TestPathFinderService_PathFromToStartNotFound(t *testing.T) {
	repository := inmemory.NewInMemoryNodeRepository(5, 5)
	//Create a new pathfinder application
	pathFinderService := NewPathFinderService(repository)

	//Create a start and end location
	start := NewLocation(10, 10, 0)
	end := NewLocation(1, 1, 0)

	//Get the path
	path := pathFinderService.PathFromTo(start, end)

	//Check if the path is correct
	if len(path) != 0 {
		t.Errorf("Path length was %d, expected 0", len(path))
	}
}

// Table-driven test PathFromTo with two test cases
// Test case 1: start node not found
// Test case 2: end node not found
func TestPathFinderService_PathFromToNotFound(t *testing.T) {
	repository := inmemory.NewInMemoryNodeRepository(5, 5)
	//Create a new pathfinder application
	pathFinderService := NewPathFinderService(repository)

	var tests = []struct {
		name  string
		start Location
		end   Location
	}{
		{"Start not found", NewLocation(10, 10, 0), NewLocation(1, 1, 0)},
		{"End not found", NewLocation(0, 0, 0), NewLocation(10, 10, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Get the path
			path := pathFinderService.PathFromTo(tt.start, tt.end)

			//Check if the path is correct
			if len(path) != 0 {
				t.Errorf("Path length was %d, expected 0", len(path))
			}
		})
	}
}

// Benchmark PathFromTo
func BenchmarkPathFinderService_PathFromTo(b *testing.B) {
	repository := inmemory.NewInMemoryNodeRepository(5, 5)
	//Create a new pathfinder application
	pathFinderService := NewPathFinderService(repository)
	//Benchmark the path
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())

		//Get a random start
		start := NewLocation(math.Round(rand.Float64()*4), math.Round(rand.Float64()*4), 0)
		//Get a random end
		end := NewLocation(math.Round(rand.Float64()*4), math.Round(rand.Float64()*4), 0)

		p := pathFinderService.PathFromTo(start, end)
		if len(p) == 0 {
			b.Errorf("Path was empty for start %v and end %v", start, end)
		}
	}
}
