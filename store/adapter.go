package store

type Adapter interface {
	Get(key string) string
	Set(key string, object string)
	Remove(key string)
}
