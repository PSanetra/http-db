package store

type Store interface {
	Get(string) (*Item, error)
	Put(string, *Item) (created bool, err error)
	Delete(string) error
}
