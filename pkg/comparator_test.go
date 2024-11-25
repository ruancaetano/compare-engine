package pkg

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCompare_EqualUnestedValue(t *testing.T) {
	expect := map[string]any{"a": 1}
	actual := map[string]any{"a": 1}

	config := CompareConfig{
		Path:        "a",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}

func TestCompare_EqualNestedValue(t *testing.T) {
	expect := map[string]any{"a": map[string]any{"b": "2"}}
	actual := map[string]any{"a": map[string]any{"b": "2"}}

	config := CompareConfig{
		Path:        "a.b",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}

func TestCompare_EqualAllArrayValue(t *testing.T) {
	expect := map[string]any{"a": []any{1, 2, 3}}
	actual := map[string]any{"a": []any{1, 2, 3}}

	config := CompareConfig{
		Path:        "a",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}

func TestCompare_EqualArrayIndexValue(t *testing.T) {
	expect := map[string]any{"a": []any{1, 2, 3}}
	actual := map[string]any{"a": []any{1, 2, 3}}

	config := CompareConfig{
		Path:        "a.2",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}

func TestCompare_EqualArrayIndexObjectValue(t *testing.T) {
	expect := map[string]any{"a": []any{map[string]any{"b": map[string]any{"id": 5.5}}}}
	actual := map[string]any{"a": []any{map[string]any{"b": map[string]any{"id": 5.5}}}}

	config := CompareConfig{
		Path:        "a.0.b.id",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}

func TestCompare_EqualAllArrayObjectValue(t *testing.T) {
	expect := map[string]any{"a": []any{map[string]any{"b": map[string]any{"id": 10}}, map[string]any{"b": map[string]any{"id": 20}}}}
	actual := map[string]any{"a": []any{map[string]any{"b": map[string]any{"id": 10}}, map[string]any{"b": map[string]any{"id": 20}}}}

	config := CompareConfig{
		Path:        "a.b.id",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}

func TestCompare_EqualArrayAllObjectValue(t *testing.T) {
	expect := map[string]any{"a": map[string]any{"b": 1, "c": 2, "d": 3}}
	actual := map[string]any{"a": map[string]any{"b": 1, "c": 2, "d": 3}}

	config := CompareConfig{
		Path:        "a",
		CompareType: CompareTypeEq,
	}
	result := Compare(expect, actual, config)
	assert.Equal(t, result, true)
}
