package digraph

import (
	"log"
	"testing"
)

// TestAdjacent verifies that the Adjacent method is working properly
func TestAdjacent(t *testing.T) {
	log.Println("TestAdjacent()")

	// Create an adjacency list
	adjList := NewAdjacencyList()

	// Generate some adjacent vertices for the list
	elements := []Vertex{1, 2, 3, 4, 5, 6}
	for _, e := range elements {
		adjList.list.PushBack(e)
	}

	// Fetch all adjacent elements from the list
	for i, a := range adjList.Adjacent() {
		if elements[i] != a {
			t.Fatalf("adjList.Adjacent() - unexpected result: %d != %d", elements[i], a)
		}
	}
}

// TestSearch verifies that the Search method is working properly
func TestSearch(t *testing.T) {
	log.Println("TestSearch()")

	// Create an adjacency list
	adjList := NewAdjacencyList()

	// Generate some adjacent vertices for the list
	elements := []Vertex{1, 2, 3}
	for _, e := range elements {
		adjList.list.PushBack(e)
	}

	// Create a table of tests and expected element results
	var tests = []struct {
		vertex interface{}
		result interface{}
	}{
		// Existing vertices
		{1, 1},
		{2, 2},
		{3, 3},
		// Non-existant vertices
		{4, nil},
		{5, nil},
		{6, nil},
	}

	// Iterate test table, check results
	for _, test := range tests {
		// Check for element
		element := adjList.Search(test.vertex)

		// If element is nil and it should not be, test fails
		if element == nil && test.result != nil {
			t.Fatalf("adjList.Search(%d) - unexpected result: nil", test.vertex)
		} else if element != nil && element != test.result {
			// If element is not nil, but is value does not match, test fails
			t.Fatalf("adjList.Search(%d) - unexpected result: %v", test.vertex, element)
		}
	}
}
