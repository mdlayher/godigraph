package digraph

import (
	"log"
	"testing"
)

// TestAddVertex verifies that the AddVertex method is working properly
func TestAddVertex(t *testing.T) {
	log.Println("TestAddVertex()")

	// Create a digraph
	graph := New()

	// Create a table of tests and expected error results
	var tests = []struct{
		vertex interface{}
		result error
	}{
		// Add vertices which do not exist
		{1, nil},
		{2, nil},
		{3, nil},
		// Add vertices which already exist
		{1, ErrVertexExists},
		{2, ErrVertexExists},
		{3, ErrVertexExists},
	}

	// Iterate test table, check results
	for _, test := range tests {
		if err := graph.AddVertex(test.vertex); err != test.result {
			t.Fatalf("graph.AddVertex(%d) - unexpected result: %s", test.vertex, err.Error())
		}
	}
}

// TestAddEdge verifies that the AddEdge method is working properly
func TestAddEdge(t *testing.T) {
	log.Println("TestAddEdge()")

	// Create a digraph, add root vertex
	graph := New()

	// Create a table of tests and expected error results
	var tests = []struct {
		source interface{}
		target interface{}
		result error
	}{
		// Add edges which do not exist
		{1, 2, nil},
		{1, 3, nil},
		{2, 3, nil},
		{3, 4, nil},
		// Add edges which already exist
		{1, 2, ErrEdgeExists},
		{3, 4, ErrEdgeExists},
		// Add edges which create a cycle
		{1, 1, ErrCycle},
		{4, 1, ErrCycle},
	}

	// Iterate test table, check results
	for _, test := range tests {
		if err := graph.AddEdge(test.source, test.target); err != test.result {
			t.Fatalf("graph.AddEdge(%d, %d) - unexpected result: %s", test.source, test.target, err.Error())
		}
	}
}
