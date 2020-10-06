package traversal

import (
	"github.com/stretchr/testify/assert"
	"graph/collections"
	"testing"
)

func TestShortestPath(t *testing.T) {
	var strGraph [][]string
	strGraph = append(strGraph, []string{"A", "B"})
	strGraph = append(strGraph, []string{"A", "C"})
	strGraph = append(strGraph, []string{"B", "D"})
	strGraph = append(strGraph, []string{"C", "E"})
	strGraph = append(strGraph, []string{"C", "F"})
	strGraph = append(strGraph, []string{"E", "G"})
	strGraph = append(strGraph, []string{"F", "H"})
	strGraph = append(strGraph, []string{"F", "I"})

	graph, err := collections.ConstructGraph(strGraph)
	assert.Nil(t, err)
	assert.NotNil(t, graph.GraphId)

	strPath := StrPath(ShortestPath(graph, "D", "I"))
	assert.NotEqual(t, strPath, PathNotFound)
	t.Log(strPath)
	assert.Equal(t, "D -> B -> A -> C -> F -> I", strPath)
}
