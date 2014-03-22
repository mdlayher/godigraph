package digraph

import (
	"container/list"
	"log"
	"testing"
)

// TestAdjacent verifies that the Adjacent method is working properly
func TestAdjacent(t *testing.T) {
	log.Println("TestAdjacent()")

	// Create an adjacency list
	adjList := AdjacencyList{list.New()}

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
