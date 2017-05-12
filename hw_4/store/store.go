// Package store implements a simple, object oriented key/value store.
package store

// Data is the key value store's memory type.
type Data map[string]string

// Repository is an abstract representation of a data repository for keeping
// the actual key/value store data.
type Repository interface {
	ReadDb() Data
	WriteDb(Data)
}

// Store contains a filename where the actual data is stored and a map of the
// data.
type Store struct {
	Repo Repository
	data Data
}

// NewStore returns the pointer to a new Store struct.
func NewStore(r Repository) *Store {
	return &Store{
		Repo: r,
		data: r.ReadDb(),
	}
}

// Get returns the value for the given key.
// Returns empty string if key not found in store.
func (s *Store) Get(k string) string {
	v, ok := s.data[k]
	if ok {
		return v
	}
	return ""
}

// Set sets the given value for the given key in the store.
// Silently overwrites existing key/value pair.
func (s *Store) Set(k, v string) {
	s.data[k] = v
}

// GetAll returns all key/value pairs in the store.
func (s *Store) GetAll() Data {
	return s.data
}

// Flush writes all changes made in memory to the store's repository.
func (s *Store) Flush() {
	s.Repo.WriteDb(s.data)
}
