package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	conf "notebook/config"
	"notebook/database"
)

var Cache *cache.Cache

func init() {
	ctx := context.Background()
	cacheConf := conf.Config.Cache
	if cacheConf.Type == conf.Memory {
		bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(cacheConf.Expiration))
		cacheStore := store.NewBigcache(bigcacheClient, nil) // No options provided (as second argument)
		Cache = cache.New(cacheStore)

	} else if cacheConf.Type == conf.Redis {
		cacheStore := store.NewRedis(database.Redis, &store.Options{
			Expiration: cacheConf.Expiration,
		})
		Cache = cache.New(cacheStore)
	} else {
		panic("config cache.type must \"memory\" or \"redis\"")
	}
	err := Cache.Set(ctx, "test-key", []byte("test"), nil)
	if err != nil {
		panic(err)
	}
	Cache.Delete(ctx, "test-key")
}
