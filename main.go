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

func New() *Memory {
	return &Memory{storage: make(map[string]Data)}
}

var ErrKeyNotFound error

// TODO: Store
func (m *Memory) Set(key string, data Data) error {
	m.storage[key] = data
	return nil
}

// TODO: Retrieve
func (m *Memory) Get(key string) (Data, error) {
	
	return Data{}, ErrKeyNotFound
}

func (m *Memory) HasData(key string) bool {
	v := m.storage[key]
	if v != Data{{[] []} []} {
		return true
	}
	return false
}

// TODO: Delete
func (m *Memory) Delete(key string) error {
	delete(m.storage, key)
	return nil
}

// TODO: Serialize data
