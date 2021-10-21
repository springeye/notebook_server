package store

import "github.com/gin-gonic/gin"

type IStore interface {
	Get(key string) string
	Set(key string, object string)
	Remove(key string)
}
type appStore struct {
	store Adapter
}

func (a appStore) Get(key string) string {
	return a.store.Get(key)
}

func (a appStore) Set(key string, object string) {
	a.store.Set(key, object)
}

func (a appStore) Remove(key string) {
	a.store.Remove(key)
}

func NewStore(store Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := &appStore{store: store}
		c.Set("cache_store", s)
		c.Next()
	}
}
func Default(c *gin.Context) IStore {
	return c.MustGet("cache_store").(IStore)
}
