package container

import (
	"math"
	"sync"
	"sync/atomic"
)

type MemoryContainer struct {
	conainainer map[int32]map[int64]bool
	size        int64
	mutex       sync.Mutex
}

func NewMemoryContainer() *MemoryContainer {
	return &MemoryContainer{conainainer: make(map[int32]map[int64]bool)}
}

func (c *MemoryContainer) GetBit(containerId int32, val int64) (bool, error) {
	if _, ok := c.conainainer[containerId]; !ok {
		return false, nil
	}
	if _, ok := c.conainainer[containerId][val]; !ok {
		return false, nil
	}
	return true, nil
}

func (c *MemoryContainer) SetBit(containerId int32, val int64) SetStatus {
	exists, err := c.GetBit(containerId, val)
	if err != nil {
		return SetBitFailed
	} else if exists {
		return SetBitExists
	}

	if _, ok := c.conainainer[containerId]; !ok {
		c.mutex.Lock()
		c.conainainer[containerId] = make(map[int64]bool)
		c.mutex.Unlock()
	}

	c.mutex.Lock()
	c.conainainer[containerId][val] = true
	c.mutex.Unlock()

	return SetBitOK
}

func (c *MemoryContainer) GetMaxBitSize() int64 {
	return math.MaxInt32
}

func (c *MemoryContainer) Reset() bool {
	c.mutex.Lock()
	c.conainainer = make(map[int32]map[int64]bool)
	c.size = 0
	c.mutex.Unlock()
	return true
}

func (c *MemoryContainer) Export() (map[int32]map[int64]bool, error) {
	return c.conainainer, nil
}

func (c *MemoryContainer) Import(data map[int32]map[int64]bool) error {
	c.mutex.Lock()
	c.conainainer = data
	c.mutex.Unlock()
	return nil
}

func (c *MemoryContainer) IncreaseSize() {
	atomic.SwapInt64(&c.size, c.size+1)
}

func (c *MemoryContainer) GetSize() int64 {
	return c.size
}

func (c *MemoryContainer) SetSize(size int64) {
	atomic.SwapInt64(&c.size, size)
}
