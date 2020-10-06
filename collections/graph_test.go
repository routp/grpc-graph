package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructGraph(t *testing.T) {
	var strGraph [][]string
	strGraph = append(strGraph, []string{"A", "B"})
	strGraph = append(strGraph, []string{"B", "C"})
	strGraph = append(strGraph, []string{"B", "D"})
	strGraph = append(strGraph, []string{"C", "E"})
	strGraph = append(strGraph, []string{"C", "F"})
	strGraph = append(strGraph, []string{"E", "G"})
	strGraph = append(strGraph, []string{"F", "H"})
	strGraph = append(strGraph, []string{"F", "I"})
	graph, err := ConstructGraph(strGraph)
	assert.Nil(t, err)
	graph.ToString()
	assert.NotNil(t, graph.GraphId)
	assert.NotNil(t, graph.Nodes)
}
