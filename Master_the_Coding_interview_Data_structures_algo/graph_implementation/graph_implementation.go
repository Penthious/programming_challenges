package main

import (
	"encoding/json"
	"fmt"
)

type graph struct {
	Vertices map[int]*vertex
}

type vertex struct {
	Adjacent map[int]*vertex
}

func newGraph() *graph {
	return &graph{}
}

func (g *graph) addVertex(data int) {
	if g.Vertices == nil {
		g.Vertices = make(map[int]*vertex)
	}

	g.Vertices[data] = &vertex{}

}
func (g *graph) addEdge(main, link int) bool {
	if g.Vertices[main] == nil || g.Vertices[link] == nil {
		fmt.Println("Missing vertex")
		return false
	}

	if g.Vertices[main].Adjacent == nil {
		g.Vertices[main].Adjacent = map[int]*vertex{}
	}
	// if g.Vertices[link].Adjacent == nil {
	// 	g.Vertices[link].Adjacent = map[int]*vertex{}
	// }

	jsonOut(g)
	g.Vertices[main].Adjacent[link] = g.Vertices[link]
	// g.Vertices[link].Adjacent[main] = g.Vertices[main]
	jsonOut(g)
	return true

}

func main() {

	g := newGraph()

	g.addVertex(1)
	g.addVertex(3)
	g.addVertex(6)
	g.addEdge(1, 3)
	// g.addEdge(1, 6)
	// jsonOut(g)
}

func jsonOut(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
}
