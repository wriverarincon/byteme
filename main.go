package byteme

type Cache interface {
	Set(key string, data Data) error
	Get(key string) (Data, error)
	Delete(key string) error
}

type Data struct {
	Schema Schema
	Values []string
}
type Schema struct {
	Columns []string
	Types   []string
}
type Memory struct {
	storage map[string]Data
}

type ErrKeyNotFound struct {
	key     string
	message string
}

// New creates a new cache memory storage.
func New() *Memory {
	return &Memory{storage: make(map[string]Data)}
}

func (e *ErrKeyNotFound) Error() string {
	return fmt.Sprintf("%s - %s", e.key, e.message)
}

// Set inserts new data to memory.
func (m *Memory) Set(key string, data Data) error {
	m.storage[key] = data
	return nil
}

// Get fetches data from memory.
func (m *Memory) Get(key string) (Data, error) {
	
	return Data{}, ErrKeyNotFound
}

// HasData returns whether the key contains valid data.
func (m *Memory) HasData(key string) bool {
	v := m.storage[key]
	if v != Data{{[] []} []} {
		return true
	}
	return false
}

// Delete removes the specified key from memory.
func (m *Memory) Delete(key string) error {
	delete(m.storage, key)
	return nil
}

// TODO: Serialize data
