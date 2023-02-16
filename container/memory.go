package container

type MemoryContainer struct {
	conainainer map[int32]map[int64]bool
}

func NewMemoryContainer() *MemoryContainer {
	return &MemoryContainer{conainainer: make(map[int32]map[int64]bool)}
}

func (c MemoryContainer) GetBit(containerId int32, val int64) bool {
	if _, ok := c.conainainer[containerId]; !ok {
		return false
	}
	if _, ok := c.conainainer[containerId][val]; !ok {
		return false
	}
	return true
}

func (c *MemoryContainer) SetBit(containerId int32, val int64) SetStatus {
	if c.GetBit(containerId, val) {
		return SetBitExists
	}
	if _, ok := c.conainainer[containerId]; !ok {
		c.conainainer[containerId] = make(map[int64]bool)
	}
	c.conainainer[containerId][val] = true
	return SetBitOK
}

func (c MemoryContainer) GetMaxBitSize() int64 {
	return 2 ^ 31 - 1
}

func (c *MemoryContainer) Reset() bool {
	c.conainainer = make(map[int32]map[int64]bool)
	return true
}

func (c *MemoryContainer) Export() map[int32]map[int64]bool {
	return c.conainainer
}

func (c *MemoryContainer) Import(data map[int32]map[int64]bool) bool {
	c.conainainer = data
	return true
}
