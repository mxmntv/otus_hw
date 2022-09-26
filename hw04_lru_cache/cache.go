package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, ok := c.items[key]
	if ok {
		item.Value = cacheItem{key, value}
		c.queue.MoveToFront(item)
		return true
	} else if c.capacity == c.queue.Len()-1 {
		removedItem := c.queue.Back()
		val, ok := removedItem.Value.(cacheItem)
		if ok {
			delete(c.items, val.key)
		}
		c.queue.Remove(removedItem)
		setItem := c.queue.PushFront(cacheItem{
			key,
			value,
		})
		c.items[key] = setItem
		return false
	}
	setItem := c.queue.PushFront(cacheItem{key, value})
	c.items[key] = setItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(item)
		val, ok := item.Value.(cacheItem)
		if ok {
			return val.value, true
		}
	}
	return nil, false
}

func (c *lruCache) Clear() {
	q, ok := c.queue.(*list)
	if ok {
		q.length = 0
		q.first = nil
		q.last = nil
	}
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
