package bloomfilter

import (
	"math"
)

func (c baseBloomFilter) Error() float64 {
	return c.errorRate
}

func (c baseBloomFilter) Size() int64 {
	return c.keySize
}

func (c baseBloomFilter) Capacity() int64 {
	return c.capacity
}

func (c baseBloomFilter) ComputeStoreSize() int64 {
	// unit: bytes
	return int64(math.Ceil(float64(c.bloomSize) / 8.0))
}

func (c *baseBloomFilter) Reset() {
	c.keySize = 0
	c.container.Reset()
}
