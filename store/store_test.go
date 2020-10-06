package store

import (
	"github.com/stretchr/testify/assert"
	"graph/collections"
	"testing"
)

func TestStore(t *testing.T) {
	var strGraph [][]string
	strGraph = append(strGraph, []string{"Node1", "Node2"})
	strGraph = append(strGraph, []string{"Node2", "Node3"})
	strGraph = append(strGraph, []string{"Node2", "Node2"})
	strGraph = append(strGraph, []string{"Node3", "Node1"})
	strGraph = append(strGraph, []string{"Node3", "Node4"})
	strGraph = append(strGraph, []string{"Node4", "Node5"})
	strGraph = append(strGraph, []string{"Node5", "Node1"})

	graph, err := collections.ConstructGraph(strGraph)
	assert.Nil(t, err)
	assert.NotNil(t, graph.GraphId)

	// Save
	Add(graph.GraphId, strGraph)

	// Get
	strGraph2 := Get(graph.GraphId)
	assert.Equal(t, strGraph, strGraph2)

	// Reconstruct
	graph2, err2 := collections.ConstructGraphWithId(graph.GraphId, strGraph2)
	assert.Nil(t, err2)
	assert.Equal(t, graph.GraphId, graph2.GraphId)
	assert.Equal(t, strGraph, strGraph2)

	// Delete
	isRemoved := Remove(graph.GraphId)
	assert.True(t, isRemoved)

	// Delete Again
	isRemoved = Remove(graph.GraphId)
	assert.False(t, isRemoved)

	// Get again
	strGraph3 := Get(graph.GraphId)
	assert.Nil(t, strGraph3)

}

