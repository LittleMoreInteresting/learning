package singleton

import "sync"

type singleton struct {
	count int
}

var (
	instance *singleton
	mutex sync.Mutex
)

func New() *singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(singleton)
		}
		mutex.Unlock()
	}
	return instance;
}


