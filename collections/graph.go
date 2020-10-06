package collections

import (
	"fmt"
	"github.com/google/uuid"
	"graph/gerror"
	"reflect"
)

// Graph
type Graph struct {
	GraphId string
	Nodes 	map[string]*Node
}

// Create a new Graph
func NewGraph() Graph {
	return Graph{
		GraphId: uuid.New().String(),
		Nodes: make(map[string]*Node),
	}
}

// Create a new Graph
func NewGraphWithId(graphId string) Graph {
	return Graph{
		GraphId: graphId,
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) GetNode(sn string) *Node {
	for _, node := range g.Nodes {
		if node.ID == sn {
			return node
		}
	}
	return nil
}

// Add edge to graph
func (g *Graph) addEdges(edge1, edge2 string) {
	isLoop := edge1==edge2

	node1 := g.GetNode(edge1)
	if node1 == nil {
		node1 = NewNode(edge1)
	}

	if !isLoop {
		node2 := g.GetNode(edge2)
		if node2 == nil {
			node2 = NewNode(edge2)
		}
		node1.Edges = append(node1.Edges, node2)
		node2.Edges = append(node2.Edges, node1)

		g.Nodes[edge1] = node1
		g.Nodes[edge2] = node2
	}else{
		node1.Edges = append(node1.Edges, node1)
		g.Nodes[edge1] = node1
	}
}

// Construct a graph
func ConstructGraph(edges [][]string) (g *Graph, err *gerror.GraphError) {
	return ConstructGraphWithId("", edges)
}

func ConstructGraphWithId(graphId string, edges [][]string) (g *Graph, err *gerror.GraphError) {
	var graph Graph
	if len(graphId) == 0 {
		graph = NewGraph()
	}else{
		graph = NewGraphWithId(graphId)
	}

	if edges == nil || len(edges) == 0{
		return &graph, nil
	}

	for  i := 0; i < len(edges); i++ {
		if len(edges[i])== 2 {
			graph.addEdges(edges[i][0], edges[i][1])
		}else{
			return nil, gerror.NewGraphError(gerror.EdgeCountCode, fmt.Sprintf(gerror.EdgeCountError, i))
		}
	}
	fmt.Println(graph.ToString())
	return &graph, nil
}


// Print a graph
func (g *Graph) ToString() string{
	s := "graphID: " + g.GraphId +"\n"
	keys := reflect.ValueOf(g.Nodes).MapKeys()
	for i := 0; i < len(keys); i++ {
		key := keys[i].String()
		s += key + " -> "
		edges := g.Nodes[key].Edges
		for j := 0; j < len(edges); j++ {
			s += edges[j].ID + " "
		}
		s += "\n"
	}
	return s
}