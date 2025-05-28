package byteme

import "testing"

func TestCache(t *testing.T) {
	m := New()
	name := "john"

	schema := Schema{
		Columns: []string{"name", "age"},
		Types:   []string{"string", "int"},
	}

	data := Data{
		Schema: schema,
		Values: []string{name, "25"},
	}

	m.Set(name, data)

	result, err := m.Get(name)
	if err != nil {
		t.Fatalf("getting cache memory: %v", err)
	}
	if result.Values[1] != "25" {
		t.Fatalf("expected '25', got: %v", result.Values[1])
	}
	if err := m.Delete(name); err != nil {
		t.Fatalf("deleting cache memory: %v", err)
	}
	if v, err := m.Get(name); err == nil {
		t.Fatalf("expected nothing, got: %v %v", v, m.storage[name])
	}
}
