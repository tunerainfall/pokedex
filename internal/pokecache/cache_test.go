package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache, err := NewCache(5 * time.Second)
	if err != nil {
		t.Fatalf("NewCache failed: %v", err)
	}
	if cache.interval != 5*time.Second {
		t.Errorf("expected interval 5s, got %v", cache.interval)
	}
	if cache.cache == nil {
		t.Error("cache map not initialized")
	}
}

func TestAddGet(t *testing.T) {
	cache, _ := NewCache(5 * time.Second)
	key := "testKey"
	val := []byte("testValue")

	err := cache.Add(key, val)
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	retrieved, ok := cache.Get(key)
	if !ok {
		t.Fatal("Get failed: key not found")
	}
	if string(retrieved) != string(val) {
		t.Errorf("expected %s, got %s", val, retrieved)
	}
}

func TestReapLoop(t *testing.T) {
	interval := 1 * time.Second
	cache, _ := NewCache(interval)
	key := "testKey"
	val := []byte("testValue")

	cache.Add(key, val)

	// Check it's there
	_, ok := cache.Get(key)
	if !ok {
		t.Fatal("key should be present initially")
	}

	// Wait for expiration (interval + a bit)
	time.Sleep(interval + 100*time.Millisecond)

	// Now it should be gone
	_, ok = cache.Get(key)
	if ok {
		t.Error("key should have been reaped")
	}
}
