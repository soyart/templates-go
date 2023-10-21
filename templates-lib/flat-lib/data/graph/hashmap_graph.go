package graph

import (
	"example.com/flatlib/data"
)

type nodeValue any

// HashMapGraphV1[T] is a GenericGraph that represents node connections as a hash map `map[Node[T]][]Node[T].
// The edges in this interface is of the same type as the node `Node[T]`, since not all use cases need real edges.
type HashMapGraphV1[T nodeValue] Graph[
	// The graph node is Node[T]
	Node[T],
	// Since there's no edge weight, this graph will use the connected nodes to represent a node's edges
	Node[T],
	// The weight can be of any types, BUT ONLY NIL IS VALID if using the default implementation of unweighted graph
	any,
]

// Read-only Node
type Node[T nodeValue] data.GetValuer[T]
