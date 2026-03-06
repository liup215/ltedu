package cache

import (
	"testing"
	"time"
)

func TestLRU_SetGet(t *testing.T) {
	c := New[string, int](3, 0)
	c.Set("a", 1)
	c.Set("b", 2)

	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatalf("expected Get(a)=1, got %v %v", v, ok)
	}
	if v, ok := c.Get("b"); !ok || v != 2 {
		t.Fatalf("expected Get(b)=2, got %v %v", v, ok)
	}
}

func TestLRU_Eviction(t *testing.T) {
	c := New[int, int](2, 0)
	c.Set(1, 10)
	c.Set(2, 20)
	// Access 1 to make it recently used
	c.Get(1)
	// Adding 3 should evict 2 (least recently used)
	c.Set(3, 30)

	if _, ok := c.Get(2); ok {
		t.Fatal("expected key 2 to be evicted")
	}
	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatalf("expected key 1 to survive eviction, got %v %v", v, ok)
	}
	if v, ok := c.Get(3); !ok || v != 30 {
		t.Fatalf("expected key 3 to be present, got %v %v", v, ok)
	}
}

func TestLRU_TTLExpiry(t *testing.T) {
	c := New[string, string](10, 50*time.Millisecond)
	c.Set("x", "hello")

	if v, ok := c.Get("x"); !ok || v != "hello" {
		t.Fatalf("expected Get(x)=hello before expiry, got %v %v", v, ok)
	}

	time.Sleep(100 * time.Millisecond)

	if _, ok := c.Get("x"); ok {
		t.Fatal("expected key x to be expired after TTL")
	}
}

func TestLRU_Delete(t *testing.T) {
	c := New[string, int](5, 0)
	c.Set("del", 99)
	c.Delete("del")

	if _, ok := c.Get("del"); ok {
		t.Fatal("expected key del to be absent after Delete")
	}
}

func TestLRU_Purge(t *testing.T) {
	c := New[string, int](5, 0)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Purge()

	if c.Len() != 0 {
		t.Fatalf("expected empty cache after Purge, got Len=%d", c.Len())
	}
}

func TestLRU_Update(t *testing.T) {
	c := New[string, int](5, 0)
	c.Set("k", 1)
	c.Set("k", 2)

	if v, ok := c.Get("k"); !ok || v != 2 {
		t.Fatalf("expected updated value 2, got %v %v", v, ok)
	}
	if c.Len() != 1 {
		t.Fatalf("expected Len=1 after update, got %d", c.Len())
	}
}
