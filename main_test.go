package main

import (
	"memoryStorage/lib/storage"
	"sync"
	"testing"
)

type Config struct {
	sync.RWMutex
	endpoint string
}

func BenchmarkPMutexSet(b *testing.B) {
	var localStorage storage.SyncStorage
	store := localStorage.Create()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Set("key", "value")
		}
	})
}

func BenchmarkPMutexGet(b *testing.B) {
	var localStorage storage.SyncStorage
	store := localStorage.Create()
	store.Set("key", "value")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = store.Get("key")
		}
	})
}

func BenchmarkPAtomicSet(b *testing.B) {
	var localStorage storage.AtomicStorage
	store := localStorage.Create()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Set("key", "value")
		}
	})
}

func BenchmarkPAtomicGet(b *testing.B) {
	var localStorage storage.AtomicStorage
	store := localStorage.Create()
	store.Set("key", "value")
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = store.Get("key")
		}
	})
}

func BenchmarkMemoryGet(b *testing.B) {
	var localStorage = make(storage.Storage)
	localStorage["key"] = "value"
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = localStorage["key"]
		}
	})
}
