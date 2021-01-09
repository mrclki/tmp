package graph

import (
	"fmt"
)

var ErrCircularDependency = fmt.Errorf("graph: circular dependency found")

func New() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

// Graph reprecent a directed acyclic graph
type Graph struct {
	Nodes map[string]*Node
}

func (g *Graph) AddNode(nodes ...*Node) {
	for _, node := range nodes {
		g.Nodes[node.Name] = node
	}
}

func (g *Graph) GetNode(name string) (*Node, bool) {
	node, ok := g.Nodes[name]
	return node, ok
}

func Sort(g *Graph) ([]*Node, error) {
	var sorted []*Node
	for len(g.Nodes) > 0 {
		ready := NewSet()
		for _, node := range g.Nodes {
			if len(node.Edges) == 0 {
				ready.Add(node)
			}
		}
		if ready.Len() == 0 {
			return nil, ErrCircularDependency
		}
		for node := range ready.Range() {
			delete(g.Nodes, node.Name)
			sorted = append(sorted, node)
		}

		// go through the nodes with edges
		for _, node := range g.Nodes {
			edgeSet := NewSet()
			for _, edge := range node.Edges {
				edgeSet.Add(edge)
			}
			diffSet := edgeSet.Diff(ready)
			node.Edges = make([]*Node, 0)
			for edge := range diffSet.Range() {
				node.Edges = append(node.Edges, edge)
			}
		}
	}
	return sorted, nil
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Edges: make([]*Node, 0),
	}
}

type Node struct {
	Name  string
	Edges []*Node
}

func (n *Node) AddEdge(edges ...*Node) {
	for _, edge := range edges {
		n.Edges = append(n.Edges, edge)
	}
}
