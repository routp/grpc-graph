package traversal

import (
	"container/list"
	"graph/collections"
	"strings"
)

const PathNotFound  = "Path not found"
func ShortestPath(g *collections.Graph, source string, dest string) *list.List{
	// Map: CurrentNode -> Parent Node
	visited := make(map[string]*collections.Node)

	// If given source node exists in graph
	srcNode := g.GetNode(source)
	if srcNode == nil {
		return nil
	}

	// BFS: Start from source with parent as empty
	queue := list.New()
	queue.PushBack(srcNode)
	visited[srcNode.ID] = nil

	for queue.Len() > 0 {
		e := queue.Front()

		current:= e.Value.(*collections.Node)
		for _, edge := range current.Edges {
			if _, ok := visited[edge.ID]; !ok {
				queue.PushBack(edge)
				visited[edge.ID] = current
				if edge.ID == dest {
					paths := getPath(visited, edge);
					return paths
				}
			}
		}
		queue.Remove(e)
	}
	return nil
}

func getPath(parents map[string]*collections.Node, destNode *collections.Node) *list.List {
	list := list.New()
	node := destNode
	for node != nil {
		list.PushFront(node)
		parent := parents[node.ID]
		node = parent
	}
	return list
}

func StrPath(paths *list.List) string {
	s := ""
	delim := " -> "
	if paths != nil {
		for e := paths.Front(); e != nil; e = e.Next() {
				s += e.Value.(*collections.Node).ID + delim
			}
		}
	if strings.Contains(s, delim){
		s = s[0:strings.LastIndex(s, delim)]
	}

	if len(s) == 0{
		s = PathNotFound
	}
	return s
}