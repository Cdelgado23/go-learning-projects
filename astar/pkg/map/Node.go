package _map

// Node represents a node in the graph
type Node struct {
	Location  *Point3D
	Neighbors []*Node
}
