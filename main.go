package main

import (
	"fmt"
	"time"
)

func main() {
	cache := newCache()

	for i := 0; i < 1000; i++ {
		k, v := fmt.Sprint("key", i), fmt.Sprint("value", i)
		fmt.Println("Putting key: ", k, " value: ", v, "expireAt: ", time.Now().Add(15*time.Second))
		cache.put(k, v)
	}

	// get() the same key 10 times will increment the ttl time
	// key3 will be evicted the last
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		key3 := cache.get("key3")
		fmt.Println("Getting key: ", "key3", " value: ", key3)
	}

	fmt.Println("len(m):", len(cache.items))
	wg.Wait()
}
