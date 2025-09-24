package singleton

import (
	"sync"
	"sync/atomic"
)

type singleton map[string]string

var once sync.Once
var lock = &sync.Mutex{}
var atomicinz uint64

var (
	instance *singleton
)

// GetInstance returns the singleton instance non thread safe_ sync
func GetInstance_example_1() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// GetInstance returns the singleton instance thread safe _async
func GetInstance_example_2() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// GetInstance returns the singleton instance thread safe _async
// Best choice for race condition
func GetInstance_example_3() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func GetInstance_example_4() *singleton {

	if atomic.LoadUint64(&atomicinz) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if atomicinz == 0 {
		instance = &singleton{}
		atomic.StoreUint64(&atomicinz, 1)
	}

	return instance
}

// withtout singleton
func NewMap() map[string]string {
	return make(map[string]string)
}
