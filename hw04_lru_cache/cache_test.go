package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(5)
		data := map[Key]int{"A": 100, "B": 200, "C": 300, "D": 400, "E": 500}
		for i, v := range data {
			c.Set(i, v)
		}
		val, ok := c.Get("A")
		require.True(t, ok)
		require.Equal(t, 100, val)

		c.Clear()

		val, ok = c.Get("D")
		require.False(t, ok)
		require.Nil(t, val)

		cta, ok := c.(*lruCache)
		if ok {
			require.Empty(t, cta.items)
			l, ok := cta.queue.(*list)
			if ok {
				require.Nil(t, l.first)
				require.Nil(t, l.last)
				require.Equal(t, 0, l.length)
			}
		}
	})

	t.Run("push out", func(t *testing.T) {
		c := NewCache(5)
		data := []cacheItem{{"A", 100}, {"B", 200}, {"C", 300}, {"D", 400}, {"E", 500}}
		for _, v := range data {
			c.Set(v.key, v.value)
		}

		val, ok := c.Get("D")
		require.True(t, ok)
		require.Equal(t, 400, val)

		wasInCache := c.Set("A", 150)
		require.True(t, wasInCache)

		val, ok = c.Get("A")
		require.True(t, ok)
		require.Equal(t, 150, val)

		result := []int{150, 400, 500, 300, 200}
		cta, _ := c.(*lruCache)
		k, _ := cta.queue.(*list)
		next := k.first
		for _, v := range result {
			e, _ := next.Value.(cacheItem)
			r, _ := e.value.(int)
			require.Equal(t, v, r)
			next = next.Next
		}

		wasInCache = c.Set("F", 50)
		require.False(t, wasInCache)

		result = []int{50, 150, 400, 500, 300}
		cta, _ = c.(*lruCache)
		k, _ = cta.queue.(*list)
		next = k.first
		for _, v := range result {
			e, _ := next.Value.(cacheItem)
			r, _ := e.value.(int)
			require.Equal(t, v, r)
			next = next.Next
		}

		c = NewCache(4)
		data = []cacheItem{{"A", 100}, {"B", 200}, {"C", 300}, {"D", 400}, {"E", 500}}
		for _, v := range data {
			c.Set(v.key, v.value)
		}

		result = []int{500, 400, 300, 200}
		cta, _ = c.(*lruCache)
		k, _ = cta.queue.(*list)
		next = k.first
		for _, v := range result {
			e, _ := next.Value.(cacheItem)
			r, _ := e.value.(int)
			require.Equal(t, v, r)
			next = next.Next
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
