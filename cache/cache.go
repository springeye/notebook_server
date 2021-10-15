package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	conf "notebook/config"
	"notebook/database"
)

var Cache ICache

type ICache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, object interface{}, options *store.Options) error
	Delete(ctx context.Context, key interface{}) error
	Invalidate(ctx context.Context, options store.InvalidateOptions) error
	Clear(ctx context.Context) error
	GetType() string
}

func init() {
	var c *cache.Cache
	ctx := context.Background()
	cacheConf := conf.Conf.Cache
	if cacheConf.Type == conf.Memory {
		bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(cacheConf.Expiration))
		cacheStore := store.NewBigcache(bigcacheClient, nil) // No options provided (as second argument)
		c = cache.New(cacheStore)

	} else if cacheConf.Type == conf.Redis {

		cacheStore := store.NewRedis(database.Redis, &store.Options{
			Expiration: cacheConf.Expiration,
		})
		c = cache.New(cacheStore)
	} else {
		panic("config cache.type must \"memory\" or \"redis\"")
	}
	err := c.Set(ctx, "test-key", []byte("test"), nil)
	if err != nil {
		panic(err)
	}
	c.Delete(ctx, "test-key")
	Cache = appCache{Cache: c}

}
