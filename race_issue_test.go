package cache

import (
	"sync"
	"testing"
)

func TestConcurrentOnEvictedAndDelete(t *testing.T) {
	c := New(NoExpiration, 0)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			c.Set("key", i, NoExpiration)
			c.Delete("key")
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			c.OnEvicted(func(string, interface{}) {})
			c.OnEvicted(nil)
		}
	}()
	wg.Wait()
}
