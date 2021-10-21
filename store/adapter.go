package store

type Adapter interface {
	Get(key string) interface{}
	Set(key string, object interface{})
	Remove(key string)
}
