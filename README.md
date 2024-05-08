<img src="https://raw.githubusercontent.com/cherkesky/inmemcache/master/logo.png" height="250" width="250">


### by Guy Cherkesky | [LinkedIn](http://linkedin.com/in/cherkesky) | [Website](http://cherkesky.com)

### In-Memory Cache with TTL and Eviction Policy
#### Go implementation

This repository contains an implementation of an in-memory cache written in GoLang. The cache supports Time-To-Live (TTL) for cached items and an eviction policy to remove unused items from the cache. Additionally, it extends the TTL of an accessed item upon retrieval.

## Features

- **Time-To-Live (TTL):** Cached items have an associated TTL, after which they are automatically removed from the cache.
- **Eviction Policy:** An eviction policy ensures that the cache does not exceed a specified capacity by removing the least recently used (LRU) items when necessary.
- **TTL Extension:** When an item is accessed (read or updated), its TTL is extended to ensure it remains in the cache longer.

## Makefile top commands:

Run:
Run an implementation example  
```bash
$ make run
```

Test:
Run unit & integration tests
```bash
$ make test
```

Build:
Build the source file 
```bash
$ make build
```

