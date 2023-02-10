package container

type MemoryContainer struct {
	conainainer map[int64]bool
}

func NewMemoryContainer() *MemoryContainer {
	return &MemoryContainer{conainainer: make(map[int64]bool)}
}

func (c MemoryContainer) GetBit(containerId int32, val int64) bool {
	return c.conainainer[val]
}

func (c *MemoryContainer) SetBit(containerId int32, val int64) SetStatus {
	if c.GetBit(containerId, val) {
		return SetBitExists
	}
	c.conainainer[val] = true
	return SetBitOK
}

func (c MemoryContainer) GetMaxBitSize() int64 {
	return 2 ^ 31 - 1
}

func (c *MemoryContainer) Reset() bool {
	c.conainainer = make(map[int64]bool)
	return true
}
