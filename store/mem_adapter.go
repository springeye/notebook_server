package store

type MemAdapter struct {
	maps map[string]string
}

func (g MemAdapter) Get(key string) string {
	return g.maps[key]
}
func (g MemAdapter) Set(key string, object string) {
	g.maps[key] = object
}

func (g *MemAdapter) Remove(key string) {
	delete(g.maps, key)
}
