package digraph

import (
	"container/list"
	"errors"
	"fmt"
)

var (
	// ErrCycle is returned when creating an edge between two vertices would result
	// in a cycle in the digraph
	ErrCycle = errors.New("digraph: cycle between edges")

	// ErrEdgeExists is returned when an edge between two vertices already exists
	ErrEdgeExists = errors.New("digraph: edge already exists")

	// ErrVertexExists is returned when a vertex with the same value already exists
	ErrVertexExists = errors.New("digraph: vertex already exists")
)

// Vertex represents a vertex or "node" in the digraph
type Vertex interface{}

// Digraph represents a "digraph", or directed graph data structure
type Digraph struct {
	adjList     map[Vertex]AdjacencyList
	edgeCount   int
	vertexCount int
}

// New creates a new acyclic Digraph, and initializes its adjacency list
func New() *Digraph {
	return &Digraph{
		adjList: map[Vertex]AdjacencyList{},
	}
}

// AddVertex tries to add a new vertex to the root of the adjacency list on the digraph
func (d *Digraph) AddVertex(vertex Vertex) error {
	// Check for a previous, identical vertex
	if _, found := d.adjList[vertex]; found {
		return ErrVertexExists
	}

	// Add the vertex to the adjacency list, initialize a new linked-list
	d.adjList[vertex] = AdjacencyList{list.New()}
	d.vertexCount++

	return nil
}

// AddEdge tries to add a new edge between two vertices on the adjacency list
func (d *Digraph) AddEdge(source Vertex, target Vertex) error {
	// Ensure vertices are not identical
	if source == target {
		return ErrCycle
	}

	// Add both vertices to the graph, ignoring if they already exist
	d.AddVertex(source)
	d.AddVertex(target)

	// Check if this digraph already has this edge
	if d.HasEdge(source, target) {
		// Return false, edge already exists
		return ErrEdgeExists
	}

	// Do a depth-first search from the target to the source to determine if a cycle will
	// result if this edge is created
	if d.DepthFirstSearch(target, source) {
		// Return false, a cycle will be created
		return ErrCycle
	}

	// Retrieve adjacency list
	adjList := d.adjList[source]

	// Target was not found, so add an edge between source and target
	adjList.list.PushBack(target)
	d.edgeCount++

	// Store adjacency list
	d.adjList[source] = adjList

	return nil
}

// discovered maps out which vertices have been discovered using Depth-First Search
var discovered map[Vertex]bool

// DepthFirstSearch searches the digraph for the target vertex, using the Depth-First
// Search algorithm, and returning true if a path to the target is found
func (d *Digraph) DepthFirstSearch(source Vertex, target Vertex) bool {
	// Clear discovery map
	discovered = map[Vertex]bool{}

	// Begin recursive Depth-First Search, looking for all vertices reachable from source
	d.dfs(source)

	// Check if target was discovered during Depth-First Search
	result := discovered[target]

	// Clear discovery map, return result
	discovered = map[Vertex]bool{}
	return result
}

// dfs implements a recursive Depth-First Search algorithm
func (d *Digraph) dfs(target Vertex) {
	// Get the adjacency list for this vertex
	adjList := d.adjList[target]

	// Check all adjacent vertices
	for _, v := range adjList.Adjacent() {
		// Check if vertex has not been discovered
		if !discovered[v] {
			// Mark it as discovered, recursively continue traversal
			discovered[v] = true
			d.dfs(v)
		}
	}
}

// EdgeCount returns the number of edges in the digraph
func (d *Digraph) EdgeCount() int {
	return d.edgeCount
}

// HasEdge determines if the digraph has an existing edge between source and target,
// returning true if it does, or false if it does not
func (d *Digraph) HasEdge(source Vertex, target Vertex) bool {
	// Check if the source vertex exists
	if _, found := d.adjList[source]; !found {
		return false
	}

	// Retrieve adjacency list for this source
	adjList := d.adjList[source]

	// Search for target vertex
	if v := adjList.Search(target); v != nil {
		// Vertex is adjacent, edge exists
		return true
	}

	// No result, edge does not exist
	return false
}

// Print displays a printed "tree" of the digraph to the console
func (d *Digraph) Print(root Vertex) error {
	// Check if the vertex actually exists
	if _, ok := d.adjList[root]; !ok {
		return errors.New("digraph: root node does not exist, cannot print graph")
	}

	// Begin recursive printing at the specified root vertex
	d.printRecursive(root, "")
	return nil
}

// printRecursive handles the printing of each vertex in "tree" form
func (d *Digraph) printRecursive(vertex Vertex, prefix string) {
	// Print the current vertex
	fmt.Println(prefix, "-", vertex)

	// Get the current adjacency list, get adjacent vertices
	adjList := d.adjList[vertex]
	adjacent := adjList.Adjacent()

	// Iterate all adjacent vertices
	for i, v := range adjacent {
		// If last iteration, don't add a pipe character
		if i == len(adjacent)-1 {
			d.printRecursive(v, prefix+"    ")
		} else {
			// Add pipe character to show multiple items belong to same parent
			d.printRecursive(v, prefix+"   |")
		}
	}
}

// VertexCount returns the number of vertices in the digraph
func (d *Digraph) VertexCount() int {
	return d.vertexCount
}
