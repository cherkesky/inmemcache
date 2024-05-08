package main

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := newCache()
	if cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestPut(t *testing.T) {
	cache := newCache()
	cache.put("key", "value")
	got := len(cache.items)
	want := 1
	if got != want {
		t.Errorf("Expected cache to have 1 item")
	}
}

func TestGet(t *testing.T) {
	cache := newCache()
	cache.put("key", "value")
	got := cache.get("key")
	want := "value"
	if got != want {
		t.Errorf("Expected value to be 'value'")
	}
}

func TestEvictionPolicy(t *testing.T) {
	cache := newCache()
	cache.put("key", "value")
	time.Sleep(10 * time.Second)
	got := cache.get("key")
	want := ""
	if got != want {
		t.Errorf("Expected value to be '%s', got '%s'", want, got)
	}
}

func TestEvictionPolicy2(t *testing.T) {
	cache := newCache()
	cache.put("key", "value")
	time.Sleep(2 * time.Second)
	cache.put("key", "value")
	time.Sleep(3 * time.Second)
	got := cache.get("key")
	want := "value"
	if got != want {
		t.Errorf("Expected value to be '%s', got '%s'", want, got)
	}
}
