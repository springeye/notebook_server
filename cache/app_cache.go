package cache

import (
	"context"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
)

type appCache struct {
	Cache *cache.Cache
}

func (a appCache) Get(ctx context.Context, key string) (interface{}, error) {
	return a.Cache.Get(ctx, key)
}

func (a appCache) Set(ctx context.Context, key string, object interface{}, options *store.Options) error {
	return a.Cache.Set(ctx, key, object, options)
}

func (a appCache) Delete(ctx context.Context, key interface{}) error {
	return a.Cache.Delete(ctx, key)
}

func (a appCache) Invalidate(ctx context.Context, options store.InvalidateOptions) error {
	return a.Cache.Invalidate(ctx, options)
}

func (a appCache) Clear(ctx context.Context) error {
	return a.Cache.Clear(ctx)
}

func (a appCache) GetType() string {
	return a.Cache.GetType()
}
