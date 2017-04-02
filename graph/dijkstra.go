package graph

import (
	"fmt"
	"log"
)

type Edge struct {
	v, len int
}

type Node struct {
	id, shortestDistance int
	visited              bool
	edges                []*Edge
}

type Graph struct {
	nodes map[int]*Node
}

func (g *Graph) String() string {
	var s string
	for k, v := range g.nodes {
		s = s + fmt.Sprintf("Node %d (id: %d, shortest distance: %d):\n", k, v.id, v.shortestDistance)
		s = s + v.edgesToString()
	}
	return s
}

func (n *Node) edgesToString() string {
	var s string
	for _, v := range n.edges {
		s += fmt.Sprintf("\tHead ID: %d\tLength: %d\n", v.v, v.len)
	}
	return s + "\n"
}

func (g *Graph) ShowShortestDistance(id int) {
	n, ok := g.nodes[id]
	if ok != true {
		log.Fatal("Attempting to print shortest distance; could not find node %d", id)
	}
	fmt.Printf("The shortest distance for Node %d is: %d\n", id, n.shortestDistance)
}

func New() *Graph {
	var g Graph
	g.nodes = make(map[int]*Node)
	return &g
}

func newNode(id int) *Node {
	var n Node
	n.id = id
	n.edges = make([]*Edge, 0)
	n.shortestDistance = -1
	return &n
}

func (n *Node) AddEdge(v, len int) {
	e := &Edge{v: v, len: len}
	n.edges = append(n.edges, e)
}

func (g *Graph) AddNode(id int) *Node {
	n := newNode(id)
	g.nodes[n.id] = n
	return n
}

const maxInt = int(^uint(0) >> 1)

func (g *Graph) ComputeShortestDistances(s int) {
	sN, ok := g.nodes[s]
	if ok == false {
		log.Fatal("Bad node ID: %d", s)
	}
	sN.shortestDistance = 0
	sN.visited = true // node is in X

	var nID int
	var minScore int

	for found := true; found == true; {
		found = false
		nID = -1
		minScore = maxInt

		// look over all the edges reaching out of X to find node
		// with min Dijkstra's greedy criterion score
		for _, x := range g.nodes {
			if x.visited == true {
				for _, e := range x.edges {
					if g.nodes[e.v].visited == false {
						if p := x.shortestDistance + e.len; p < minScore {
							found = true
							nID = e.v
							minScore = p
						}
					}
				}
			}
		}
		if found == true {
			// assign the minScore to the choice node and bring it into X
			g.nodes[nID].visited = true
			g.nodes[nID].shortestDistance = minScore
		}
	}
}

func (g *Graph) resetVisited() {
	for _, n := range g.nodes {
		n.visited = false
	}
}
