package dijkstra_test

import (
	"reflect"
	"testing"

	"github.com/flemeur/go-shortestpath/dijkstra"
)

func TestFindsShortestPath(t *testing.T) {
	vA := &vertex{id: "A"}
	vB := &vertex{id: "B"}
	vC := &vertex{id: "C"}
	vD := &vertex{id: "D"}
	vE := &vertex{id: "E"}
	vF := &vertex{id: "F"}

	// Connect A to B, C and D
	connect(vA, vB, 1)
	connect(vA, vC, 3)
	connect(vA, vD, 2)

	// Connect B to E and F
	connect(vB, vE, 1)
	connect(vB, vF, 2)

	// Connect C to D and F
	connect(vC, vD, 1)
	connect(vC, vF, 6)

	// Connect D to F
	connect(vD, vF, 4)

	// Connect E to F
	connect(vE, vF, 1)

	expPath := []*vertex{vA, vB, vF}

	path, err := dijkstra.ShortestPath(vA, vF)
	if err != nil {
		t.Errorf("expected nil error, got %+v\n", err)
	}
	if path == nil {
		t.Errorf("expected non-nil path, got %+v\n", path)
	}

	vertexPath := make([]*vertex, len(path))
	for i := range path {
		vertexPath[i] = path[i].(*vertex)
	}

	if !reflect.DeepEqual(expPath, vertexPath) {
		t.Errorf("expected %+v, got %+v\n", expPath, vertexPath)
	}
}

func TestErrorsOnNoShortestPath(t *testing.T) {
	vA := &vertex{id: "A"}
	vG := &vertex{id: "G"} // This vertex is unreachable

	path, err := dijkstra.ShortestPath(vA, vG)
	if err == nil {
		t.Errorf("expected error, got %+v\n", err)
	}
	if path != nil {
		t.Errorf("expected nil path, got %+v\n", path)
	}
}

// dijkstra.Node implementation
type vertex struct {
	id    string
	edges []edge
}

func (v *vertex) Edges() []dijkstra.Edge {
	edges := make([]dijkstra.Edge, len(v.edges))
	for i := range v.edges {
		edges[i] = v.edges[i]
	}
	return edges
}

// dijkstra.Edge implementation
type edge struct {
	destination *vertex
	weight      float64
}

func (e edge) Destination() dijkstra.Node {
	return e.destination
}

func (e edge) Weight() float64 {
	return e.weight
}

// Connect two vertices both ways
func connect(a, b *vertex, dist float64) {
	a.edges = append(a.edges, edge{destination: b, weight: dist})
	b.edges = append(b.edges, edge{destination: a, weight: dist})
}
