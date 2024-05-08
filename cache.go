package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Cache struct {
	mutex sync.Mutex
	items map[string]*Item
}

type Item struct {
	key      string
	data     string
	ttl      time.Duration
	expireAt time.Time
}

func newCache() *Cache {
	fmt.Println("Starting Cache...")

	cache := &Cache{
		items: make(map[string]*Item),
	}
	go func() {
		fmt.Println("Starting Eviction Policy...")
		for {
			time.Sleep(1 * time.Second)
			cache.mutex.Lock()
			for key, item := range cache.items {
				if item.ttl > 0 && item.expireAt.Before(time.Now()) {
					fmt.Println("Evicting item with key: ", key, item.expireAt)
					delete(cache.items, key)
					wg.Done()
					fmt.Println("len(m):", len(cache.items))
				}
			}
			cache.mutex.Unlock()
		}
	}()
	return cache
}

func (cache *Cache) put(key string, data string) {
	item := newItem(key, data)
	cache.mutex.Lock()
	cache.items[key] = item
	fmt.Println("len(m):", len(cache.items))
	cache.mutex.Unlock()
}

func newItem(key string, data string) *Item {
	item := &Item{
		data: data,
		ttl:  3 * time.Second,
		key:  key,
	}
	item.init()
	wg.Add(1)
	return item
}

func (item *Item) init() {
	if item.ttl > 0 {
		item.expireAt = time.Now().Add(item.ttl)
	}
}

func (cache *Cache) get(key string) string {
	cache.mutex.Lock()
	var data string

	if item, exists := cache.items[key]; exists {
		data = item.data
		item.init()
	} else {
		data = ""
	}
	cache.mutex.Unlock()
	return data
}
