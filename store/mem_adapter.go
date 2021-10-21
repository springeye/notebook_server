package store

type MemAdapter struct {
	maps map[string]interface{}
}

func (g MemAdapter) Get(key string) interface{} {
	return g.maps[key]
}
func (g MemAdapter) Set(key string, object interface{}) {
	g.maps[key] = object
}

func (g *MemAdapter) Remove(key string) {
	delete(g.maps, key)
}
