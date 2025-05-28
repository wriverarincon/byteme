package byteme

import "testing"

func TestCache(t *testing.T) {
	m := New()

	schema := Schema{
		Columns: []string{"name", "age"},
		Types:   []string{"string", "int"},
	}

	data := Data{
		Schema: schema,
		Values: []string{"john", "25"},
	}

	m.Set("john", data)

	result, err := m.Get("john")
	if err != nil {
		t.Fatalf("getting cache memory: %v", err)
	} else if result.Values[1] != "25" {
		t.Fatalf("expected '25', got: %v", result.Values[1])
	}
	if err := m.Delete("john"); err != nil {
		t.Fatalf("deleting cache memory: %v", err)
	} else if v, err := m.Get("john"); err == nil {
		t.Fatalf("expected nothing, got: %v", v)
	}
}
