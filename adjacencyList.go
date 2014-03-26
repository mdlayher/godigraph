package digraph

import (
	"container/list"
	"sync"
)

// AdjacencyList represents a linked-list of vertices connected by edges in the digraph
type AdjacencyList struct {
	sync.RWMutex
	list *list.List
}

// NewAdjacencyList returns a new AdjacencyList with its internal list initialized
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		list: list.New(),
	}
}

// Adjacent returns all vertices from the adjacency list
func (a *AdjacencyList) Adjacent() []Vertex {
	// Make sure list is not being modified while finding adjacent vertices
	a.RLock()
	defer a.RUnlock()

	// Slice of vertices to return
	vertices := make([]Vertex, 0)

	// Check for front vertex
	element := a.list.Front()
	if element == nil {
		return nil
	}
	vertices = append(vertices, element.Value)

	// Iterate all remaining vertices
	for {
		// Get next value, break if nil (end of list)
		element = element.Next()
		if element == nil {
			break
		}

		// Append vertex
		vertices = append(vertices, element.Value)
	}

	// Return all vertices
	return vertices
}

// Search traverses the adjancency list and attempts to find a specified vertex
func (a *AdjacencyList) Search(target Vertex) Vertex {
	// Make sure list is not being modified while searching
	a.RLock()
	defer a.RUnlock()

	// Ensure the list is not empty
	if a.list == nil || a.list.Len() == 0 {
		return nil
	}

	// Get front node, check immediately if it's the correct one
	element := a.list.Front()
	if element.Value == target {
		return element.Value
	}

	// Iterate from the front of the list
	for {
		// Check next until nil
		element = element.Next()
		if element == nil {
			break
		}

		// Check for result
		if element.Value == target {
			return element.Value
		}
	}

	// Not found
	return nil
}
