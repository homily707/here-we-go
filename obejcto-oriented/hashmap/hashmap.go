package hashmap

import "sync"

type HashMap interface {
	Set(string, interface{})
	Get(string) interface{}

	Size() int
}

var _ HashMap = (*lockHashMap)(nil)

type lockHashMap struct {
	mutex sync.Mutex
	data  map[string]interface{}
}

// Get implements HashMap.
func (m *lockHashMap) Get(k string) interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.data[k]
}

// Size implements HashMap.
func (m *lockHashMap) Size() int {
	return len(m.data)
}

func (m *lockHashMap) Set(k string, v interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[k] = v
}
