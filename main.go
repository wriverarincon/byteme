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
if m.HasData(key) {
		v := m.storage[key]	
	return &v, nil
	}
	return nil, &ErrKeyNotFound{key, "key not found"}
}

// HasData returns whether the key contains valid data.
func (m *Memory) HasData(key string) bool {
	d := m.storage[key]
	// When the key doesn't exist in the map, Go will return an empty `Data` value
	// (i.e., `Data{Schema: {}, Values: []}`), which is not directly useful for checking existence.
	// We check if the `Values` field is non-nil to determine if data exists.
	if d.Values != nil {
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
