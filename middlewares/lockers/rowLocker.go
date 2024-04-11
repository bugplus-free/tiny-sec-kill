package lockers

import (
	"sync"
	"time"
)

func PremissionLock(m *sync.RWMutex) {
	m.Lock()
}

func PremissionUnLock(m *sync.RWMutex) {
	time.Sleep(time.Duration(10) * time.Millisecond)
	m.Unlock()
}
