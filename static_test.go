package golazy

import (
	"testing"
)

func TestStaticValue(t *testing.T) {
	value := "static_value"
	s := static[string]{value: value}

	v, err := s.Value("ctx1")
	if err != nil {
		t.Error("static value should not return an error")
	}
	if v != value {
		t.Errorf("expected %q, got %q", value, v)
	}
}

func TestStaticValue_DifferentContexts(t *testing.T) {
	value := "static_value"
	s := static[string]{value: value}

	v1, _ := s.Value("ctx1")
	v2, _ := s.Value("ctx2")

	if v1 != v2 || v1 != value {
		t.Error("static value should return the same value for different contexts")
	}
}

func TestStaticClear(t *testing.T) {
	s := static[string]{value: "static_value"}
	s.Clear("ctx1")
	v, _ := s.Value("ctx1")
	if v != "static_value" {
		t.Error("clear should be a no-op for static values")
	}
}

func TestStaticClearAll(t *testing.T) {
	s := static[string]{value: "static_value"}
	s.ClearAll()
	v, _ := s.Value("ctx1")
	if v != "static_value" {
		t.Error("clear all should be a no-op for static values")
	}
}

func TestStaticValue_IntType(t *testing.T) {
	s := static[int]{value: 42}
	v, err := s.Value("any_ctx")
	if err != nil || v != 42 {
		t.Error("static value should work with different types")
	}
}
