package bloomfilter

import "github.com/Ensteinjun/go-bloomfilter/container"

func (c *baseBloomFilter) Add(value []byte) bool {
	if c.container.GetSize() >= c.capacity {
		return false
	}

	var exists bool = true
	for i := int32(0); i < c.hashNum; i++ {
		v := c.computeHash(i, value) % int64(c.bloomSize)
		containerId := int32(v / c.container.GetMaxBitSize())
		containerIndex := v % c.container.GetMaxBitSize()

		status := c.container.SetBit(containerId, containerIndex)
		if status == container.SetBitFailed {
			return false
		} else if status == container.SetBitOK {
			exists = false
		}

	}
	if !exists {
		c.container.IncreaseSize()
	}
	return true
}

func (c *baseBloomFilter) Contains(value []byte) (bool, error) {
	for i := int32(0); i < c.hashNum; i++ {
		v := c.computeHash(i, value) % int64(c.bloomSize)
		containerId := int32(v / c.container.GetMaxBitSize())
		containerIndex := v % c.container.GetMaxBitSize()

		exists, err := c.container.GetBit(containerId, containerIndex)
		if err != nil {
			return false, err
		} else if !exists {
			return false, nil
		}
	}
	return true, nil
}
