package utils

import (
	"reflect"
	"testing"
)

// ---------- Tests StrinToLower ----------
func TestStrinToLower(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"World", "world"},
		{"áÉÍÓÚ", "áéíóú"}, // prueba con acentos
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := StrinToLower(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// ---------- Tests StringSliceToLower ----------
func TestStringSliceToLower(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"A", "B", "C"}, []string{"a", "b", "c"}},
		{[]string{"Hello", "WORLD"}, []string{"hello", "world"}},
		{[]string{}, []string{}}, // slice vacío
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := StringSliceToLower(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// ---------- Tests MapKeysToLower ----------
func TestMapKeysToLower(t *testing.T) {
	tests := []struct {
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			map[string]interface{}{"Key1": 1, "KEY2": "value"},
			map[string]interface{}{"key1": 1, "key2": "value"},
		},
		{
			nil,
			nil,
		},
		{
			map[string]interface{}{},
			map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := MapKeysToLower(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// ---------- Tests Int64ToString ----------
func TestInt64ToString(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{123, "123"},
		{-456, "-456"},
		{0, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := Int64ToString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
