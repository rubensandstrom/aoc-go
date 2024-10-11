package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	vertices []*Vertex
	visited map[string]bool
}

type Vertex struct {
	key string
	adjacent []Edge
}

type Edge struct {
	next *Vertex
	weight int
}

// Adds only if vertex does not already exist, otherwise does nothing.
func (g *Graph) AddVertex(k string) {
	if !g.hasVertex(k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// Adds the vertices if it does not exist.
func (g *Graph) AddEdge(from, to string, weight int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to) 

	if !fromVertex.hasEdge(to) {
		fromVertex.adjacent = append(fromVertex.adjacent, Edge{toVertex, weight})
	}
	// TODO: if hasEdge update weight
}

func (g *Graph) getVertex(k string) *Vertex {
	g.AddVertex(k)
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

// Prints all vertices and edges as JSON format
func (g *Graph) Print() {
	fmt.Printf("{")
	for i, v := range g.vertices {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Printf("\n  \"%v\": {", v.key)
		for i, v := range v.adjacent {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Printf("\n	\"%v\": %d", v.next.key, v.weight)
		}
		fmt.Print("\n  }")
	}
	fmt.Println("\n}")
}

func (g *Graph) hasVertex(k string) bool {
	for _, v := range g.vertices {
		if k == v.key {
			return true
		}
	}
	return false
}

func (v *Vertex) hasEdge(k string) bool {
	for _, e := range v.adjacent {
		if e.next.key == k {
			return true
		}
	}
	return false
}

// TODO: Implement depth first search
func (g *Graph) DFS(from, to string) (steps, cost int, err error) {
	return 0, 0, errors.New(fmt.Sprintf("No path from %s to %s", from, to))
}

// TODO: Implement bredth first search
func (g *Graph) BFS(from, to string) (steps, cost int, err error) {
	return 0, 0, errors.New(fmt.Sprintf("No path from %s to %s", from, to))
}

// TODO: Implement A* search
func (g *Graph) AStar(from, to string) (steps, cost int, err error) {
	return 0, 0, errors.New(fmt.Sprintf("No path from %s to %s", from, to))
}

// TODO: Implement minimum spanning tree
func (g *Graph) MST() (cost int, err error) {
	return 0, errors.New("Could not find a minimum spanning tree")
}

// TODO: Implement optimal traveling salesman problem solution
func (g *Graph) TSP(at string) int {
	if g.visited[at] { 
		return 0 
	}
	g.visited[at] = true

	neighbors := g.getVertex(at).adjacent
	for _, n := range neighbors {
		g.TSP(n.next.key)
	}


	return 0
}


