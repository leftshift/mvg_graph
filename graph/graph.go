package network_graph
import (
    "gonum.org/v1/gonum/graph"
    "gonum.org/v1/gonum/graph/simple"
)

type Node struct {
    graph.Node
    Name        string
}

type Graph struct {
    simple.WeightedDirectedGraph
}

func NewGraph() *Graph {
    return &Graph{*simple.NewWeightedDirectedGraph(0, -1)}
}

func (g *Graph) NewNode() Node {
    n := g.WeightedDirectedGraph.NewNode()
    return Node{n, ""}
}
