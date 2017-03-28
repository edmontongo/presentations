
import "sync"

// * Data encapsulation

// _Simple_concurrent_hashmap_:

type ConcurrentHash struct {
	m          map[string]string // precious data
	sync.Mutex                   // maps are not goroutine safe
}

func (c ConcurrentHash) Add(key, val string) {
	c.Lock()
	c.m[key] = val
	c.Unlock()
}

func (c ConcurrentHash) Get(key string) (val string, ok bool) {
	c.Lock()
	val, ok = c.m[key]
	c.Unlock()
	return val, ok
}
