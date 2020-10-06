package store

import "sync"

var once sync.Once

type graphStore map[string][][]string

var (
	store graphStore
)

func create()  {
	once.Do(func() {
		store = make(graphStore)
	})
}

func Add(id string, graph [][]string) {
	 create()
	 store[id] = graph
}

func Get(id string) [][]string{
	create()
	if graph, ok := store[id]; ok {
		return graph
	}
	return nil
}

func Remove(id string) bool{
	create()
	if _, ok := store[id]; ok {
		delete(store, id)
		return true
	}
	return false
}
