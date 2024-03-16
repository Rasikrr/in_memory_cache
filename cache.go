package main

type Cache struct {
	m map[string]any
}

func New() *Cache {
	cache := &Cache{m: make(map[string]any)}
	return cache
}

func (c *Cache) Get(key string) any {
	return c.m[key]
}

func (c *Cache) Set(key string, value any) {
	c.m[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.m, key)
}
