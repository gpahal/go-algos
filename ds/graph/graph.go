package graph

// Node represents a node or a vertex in a graph.
type Node struct {
	ID    int
	Value int
}

// Edge represents an edge or a relationship in a graph.
type Edge struct {
	SourceID int
	TargetID int
	Weight   int
}

// Graph represents a graph consisting of nodes and edges. Nodes in the graph are identified using
// a unique auto-incrementing integer id.
type Graph struct {
	nodes             map[int]int
	edges             map[int]map[int]int
	edgesReverseIndex map[int]map[int]int

	currID int
}

// New return a new graph instance.
func New() *Graph {
	return &Graph{
		nodes:             make(map[int]int),
		edges:             make(map[int]map[int]int),
		edgesReverseIndex: make(map[int]map[int]int),
		currID:            1,
	}
}

// Len returns the number of nodes in the graph.
func (g *Graph) Len() int {
	return len(g.nodes)
}

// Empty checks whether the graph is empty.
func (g *Graph) Empty() bool {
	return len(g.nodes) == 0
}

// Clear deletes all the items from the graph.
func (g *Graph) Clear() {
	*g = *New()
	return
}

// Node returns the node with the given id. If such a node doesn't exist, nil is returned.
func (g *Graph) Node(id int) *Node {
	val, ok := g.nodes[id]
	if ok {
		return &Node{ID: id, Value: val}
	}

	return nil
}

// HasNode checks if a node with the given id exists.
func (g *Graph) HasNode(id int) bool {
	return g.Node(id) != nil
}

// AddNode adds a new node to the graph and returns the id of this new node.
func (g *Graph) AddNode(value int) int {
	id := g.currID
	g.nodes[g.currID] = value
	g.currID++
	return id
}

// UpdateNode updates the value of the node with the given id.
func (g *Graph) UpdateNode(id, value int) bool {
	_, ok := g.nodes[id]
	if !ok {
		return false
	}

	g.nodes[id] = value
	return true
}

// DeleteNode deletes the node with the given id.
func (g *Graph) DeleteNode(id int) bool {
	_, ok := g.nodes[id]
	if ok {
		delete(g.nodes, id)
		return true
	}

	return false
}

// Edge returns the edge with the given source and target ids. If such an edge doesn't exist, nil
// is returned.
func (g *Graph) Edge(sourceID, targetID int) *Edge {
	ett, ok := g.edges[sourceID]
	if ok {
		w, ok := ett[targetID]
		if ok {
			return &Edge{
				SourceID: sourceID,
				TargetID: targetID,
				Weight:   w,
			}
		}
	}

	return nil
}

// HasEdge checks if an edge exists with the diven source and target ids.
func (g *Graph) HasEdge(sourceID, targetID int) bool {
	return g.Edge(sourceID, targetID) != nil
}

// AddEdge adds a new edge to the graph.
func (g *Graph) AddEdge(sourceID, targetID, weight int) bool {
	_, ok := g.nodes[sourceID]
	if !ok {
		return false
	}

	_, ok = g.nodes[targetID]
	if !ok {
		return false
	}

	ett, ok := g.edges[sourceID]
	if ok {
		_, ok = ett[targetID]
		if ok {
			return false
		}

		ett[targetID] = weight
	} else {
		ett := make(map[int]int, 1)
		ett[targetID] = weight

		g.edges[sourceID] = ett
	}

	ets, ok := g.edgesReverseIndex[targetID]
	if ok {
		ets[sourceID] = weight
	} else {
		ets := make(map[int]int, 1)
		ets[sourceID] = weight

		g.edgesReverseIndex[targetID] = ets
	}

	return true
}

// UpdateEdge updates the weight of the edge with the given source and target ids.
func (g *Graph) UpdateEdge(sourceID, targetID, weight int) bool {
	_, ok := g.nodes[sourceID]
	if !ok {
		return false
	}

	_, ok = g.nodes[targetID]
	if !ok {
		return false
	}

	ett, ok := g.edges[sourceID]
	if !ok {
		return false
	}
	w, ok := ett[targetID]
	if !ok {
		return false
	}
	if w == weight {
		// new weight same as the previous value - no update required
		return true
	}
	ett[targetID] = weight

	g.edgesReverseIndex[targetID][sourceID] = weight

	return true
}

// AddOrUpdateEdge adds a new edge or updates the weight of an existing edge of one exists with the
// given source and target ids.
func (g *Graph) AddOrUpdateEdge(sourceID, targetID, weight int) bool {
	_, ok := g.nodes[sourceID]
	if !ok {
		return false
	}

	_, ok = g.nodes[targetID]
	if !ok {
		return false
	}

	ett, ok := g.edges[sourceID]
	if ok {
		var w int
		w, ok = ett[targetID]
		if ok {
			// Edge already present.

			if w != weight {
				// New weight not the same as the previous value. Update required.
				g.edgesReverseIndex[targetID][sourceID] = weight
			}

			return true
		}

		ett[targetID] = weight
	} else {
		ett := make(map[int]int, 1)
		ett[targetID] = weight

		g.edges[sourceID] = ett
	}

	ets, ok := g.edgesReverseIndex[targetID]
	if ok {
		ets[sourceID] = weight
	} else {
		ets := make(map[int]int, 1)
		ets[sourceID] = weight

		g.edgesReverseIndex[targetID] = ets
	}

	return true
}

// DeleteEdge deletes the edge with the given source and target ids.
func (g *Graph) DeleteEdge(sourceID, targetID int) bool {
	ett, ok := g.edges[sourceID]
	if !ok {
		return false
	}
	delete(ett, targetID)

	ets, ok := g.edgesReverseIndex[targetID]
	if !ok {
		return false
	}
	delete(ets, sourceID)

	return true
}

// NodeOutgoingEdges returns all the outgoing edges from the node with the given id.
// NOTE: The returned map should not be mutated as it's used internally. It also changes as the
// graph is mutated.
func (g *Graph) NodeOutgoingEdges(id int) map[int]int {
	ett, ok := g.edges[id]
	if !ok {
		return nil
	}

	return ett
}

// NodeIncomingEdges returns all the incoming edges from the node with the given id.
// NOTE: The returned map should not be mutated as it's used internally. It also changes as the
// graph is mutated.
func (g *Graph) NodeIncomingEdges(id int) map[int]int {
	ets, ok := g.edgesReverseIndex[id]
	if !ok {
		return nil
	}

	return ets
}

// DeleteNodeOutgoingEdges deletes all the outgoing edges from the node with the given id.
func (g *Graph) DeleteNodeOutgoingEdges(id int) bool {
	_, ok := g.nodes[id]
	if !ok {
		return false
	}

	ett, ok := g.edges[id]
	if !ok {
		return true
	}

	for targetID := range ett {
		if ets, ok := g.edgesReverseIndex[targetID]; ok {
			delete(ets, id)
		}
	}

	delete(g.edges, id)
	return true
}

// DeleteNodeIncomingEdges deletes all the incoming edges to the node with the given id.
func (g *Graph) DeleteNodeIncomingEdges(id int) bool {
	_, ok := g.nodes[id]
	if !ok {
		return false
	}

	ets, ok := g.edgesReverseIndex[id]
	if !ok {
		return true
	}

	for sourceID := range ets {
		if ett, ok := g.edges[sourceID]; ok {
			delete(ett, id)
		}
	}

	delete(g.edgesReverseIndex, id)
	return true
}

// DeleteNodeEdges deletes all the outgoing and incoming edges from the node with the given id.
func (g *Graph) DeleteNodeEdges(id int) bool {
	return g.DeleteNodeOutgoingEdges(id) && g.DeleteNodeIncomingEdges(id)
}
