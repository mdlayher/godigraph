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
			t.Fatalf("Unexpected result: %d -> %s", test.vertex, err.Error())
		}
	}
}
