package graph_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/graph"
)

func TestNew(t *testing.T) {
	newGraph := graph.New()
	newGraph.AddNode(5)
	if newGraph.Len() != 1 {
		t.Errorf("New: expected Len to be 1, got %d", newGraph.Len())
	}
}

func TestGraph_Len(t *testing.T) {
	newGraph := graph.New()
	if newGraph.Len() != 0 {
		t.Errorf("Len: expected Len to be 0, got %d", newGraph.Len())
	}

	id1 := newGraph.AddNode(5)
	id2 := newGraph.AddNode(6)
	newGraph.AddEdge(id1, id2, 1)
	if newGraph.Len() != 2 {
		t.Errorf("Len: expected Len to be 2, got %d", newGraph.Len())
	}
}

func TestGraph_Empty(t *testing.T) {
	newGraph := graph.New()
	if !newGraph.Empty() {
		t.Errorf("Empty: expected Empty to be true, got false")
	}

	id1 := newGraph.AddNode(5)
	id2 := newGraph.AddNode(6)
	newGraph.AddEdge(id1, id2, 1)
	if newGraph.Empty() {
		t.Errorf("Empty: expected Empty to be false, got true")
	}
}

func TestGraph_Clear(t *testing.T) {
	newGraph := graph.New()
	newGraph.Clear()
	if newGraph.Len() != 0 {
		t.Errorf("Clear: expected Len to be 0, got %d", newGraph.Len())
	}
}

func TestGraph_Node(t *testing.T) {
	newGraph := graph.New()
	id1 := newGraph.AddNode(5)
	id2 := newGraph.AddNode(6)
	newGraph.AddEdge(id1, id2, 1)

	el := newGraph.Node(id1)
	if el == nil || el.Value != 5 {
		t.Errorf("Node: expected Node to return Node with Value 5, got %v", el)
	}

	el = newGraph.Node(id1 + id2 + 1)
	if el != nil {
		t.Errorf("Node: expected Node to return nil, got %v", el)
	}
}

func TestGraph_HasNode(t *testing.T) {
	newGraph := graph.New()
	id1 := newGraph.AddNode(5)
	id2 := newGraph.AddNode(6)
	newGraph.AddEdge(id1, id2, 1)

	ok := newGraph.HasNode(id1)
	if !ok {
		t.Errorf("HasNode: expected HasNode to be true, got false")
	}

	ok = newGraph.HasNode(id1 + id2 + 1)
	if ok {
		t.Errorf("HasNode: expected HasNode to be false, got true")
	}
}

func TestGraph_AddNode(t *testing.T) {
	newGraph := graph.New()
	id1 := newGraph.AddNode(5)
	if newGraph.Len() != 1 {
		t.Errorf("AddNode 5: expected Len to be 1, got %d", newGraph.Len())
	}

	ok := newGraph.HasNode(id1)
	if !ok {
		t.Errorf("AddNode 5: expected HasNode to be true, got false")
	}

	id2 := newGraph.AddNode(6)
	if newGraph.Len() != 2 {
		t.Errorf("AddNode 6: expected Len to be 2, got %d", newGraph.Len())
	}

	ok = newGraph.HasNode(id2)
	if !ok {
		t.Errorf("AddNode 6: expected HasNode to be true, got false")
	}
}

func TestGraph_UpdateNode(t *testing.T) {
	newGraph := graph.New()
	id := newGraph.AddNode(5)
	el := newGraph.Node(id)
	if el == nil || el.Value != 5 {
		t.Errorf("UpdateNode: expected Node to return Node with Value 5, got %v", el)
	}

	newGraph.UpdateNode(id, 6)
	el = newGraph.Node(id)
	if el == nil || el.Value != 6 {
		t.Errorf("UpdateNode: expected Node to return Node with Value 6, got %v", el)
	}
}

func TestGraph_DeleteNode(t *testing.T) {
	newGraph := graph.New()
	id := newGraph.AddNode(5)
	newGraph.DeleteNode(id)
	el := newGraph.Node(id)
	if el != nil {
		t.Errorf("UpdateNode: expected Node to return nil, got %v", el)
	}
}
