package digraph

import (
	"container/list"
)

// AdjacencyList represents a linked-list of vertices connected by edges in the digraph
type AdjacencyList struct {
	list *list.List
}

// Adjacent returns all vertices from the adjacency list
func (a *AdjacencyList) Adjacent() []Vertex {
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

// Search traverses the adjancency list and attempts to find a specified vertex,
// returning true if the vertex is found, or false if it is not
func (a *AdjacencyList) Search(target Vertex) *list.Element {
	// Ensure the list is not empty
	if a.list == nil || a.list.Len() == 0 {
		return nil
	}

	// Get front node, check immediately if it's the correct one
	element := a.list.Front()
	if element.Value == target {
		return element
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
			return element
		}
	}

	// Not found
	return nil
}
