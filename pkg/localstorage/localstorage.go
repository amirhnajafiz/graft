package localstorage

// LocalStorage is a map of arrays. Each key
// represents a table or section. In each section there are
// entities that are presented as interfaces.
type LocalStorage interface {
	Put(key string, value interface{})
	Fetch(key string) interface{}
}
