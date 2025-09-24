package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
	data  map[string]string
	mutex sync.RWMutex // Protege operações no map
}

var once sync.Once
var lock = &sync.Mutex{}
var atomicinz uint64

var (
	instance *singleton
)

// GetInstance returns the singleton instance (non-thread-safe)
func GetInstance_example_1() *singleton {
	if instance == nil {
		instance = &singleton{
			data: make(map[string]string),
		}
	}
	return instance
}

// GetInstance returns the singleton instance (thread-safe with mutex)
func GetInstance_example_2() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleton{
			data: make(map[string]string),
		}
	}
	return instance
}

// GetInstance returns the singleton instance (thread-safe with sync.Once)
// Best choice for race condition
func GetInstance_example_3() *singleton {
	once.Do(func() {
		instance = &singleton{
			data: make(map[string]string),
		}
	})
	return instance
}

func GetInstance_example_4() *singleton {

	if atomic.LoadUint64(&atomicinz) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if atomic.LoadUint64(&atomicinz) == 0 {
		instance = &singleton{
			data: make(map[string]string),
		}
		atomic.StoreUint64(&atomicinz, 1)
	}

	return instance
}

// Métodos para interagir com o singleton - THREAD SAFE
func (s *singleton) Set(key, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.data[key] = value
}

func (s *singleton) Get(key string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.data == nil {
		return "", false
	}
	value, exists := s.data[key]
	return value, exists
}

func (s *singleton) Delete(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data != nil {
		delete(s.data, key)
	}
}

func (s *singleton) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.data == nil {
		return 0
	}
	return len(s.data)
}

// withtout singleton
func NewMap() map[string]string {
	return make(map[string]string)
}
