package _map

import "math"

type Point3D struct {
	X, Y, Z float64
}

// Distance returns the Euclidean distance between two points
func (p *Point3D) Distance(p2 *Point3D) float64 {
	return math.Sqrt((p.X-p2.X)*(p.X-p2.X) + (p.Y-p2.Y)*(p.Y-p2.Y) + (p.Z-p2.Z)*(p.Z-p2.Z))
}
