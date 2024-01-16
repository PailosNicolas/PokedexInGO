package cache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestCacheGetAndAddEntry(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	_, ok := cache.GetEntry("newEntry")

	if ok {
		t.Error("ok should have been false")
	}

	cache.AddEntry("newEntry", []byte("testing value"))

	time.Sleep(interval + time.Microsecond)

	entry, ok := cache.GetEntry("newEntry")

	if !ok {
		t.Error("error getting entry")
	}
	if entry == nil {
		t.Error("entry should have not been nil")
	}
}

func TestCacheGetAndAddEntryFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	_, ok := cache.GetEntry("newEntry")

	if ok {
		t.Error("ok should have been false")
	}

	cache.AddEntry("newEntry", []byte("testing value"))

	time.Sleep(interval / 2)

	entry, ok := cache.GetEntry("newEntry")

	if !ok {
		t.Error("error getting entry")
	}
	if entry == nil {
		t.Error("entry should have not been nil")
	}
}
