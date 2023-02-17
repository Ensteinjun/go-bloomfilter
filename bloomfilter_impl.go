package bloomfilter

import "math"

func (c *baseBloomFilter) Error() float64 {
	return c.errorRate
}

func (c *baseBloomFilter) Size() int64 {
	return c.container.GetSize()
}

func (c *baseBloomFilter) Capacity() int64 {
	return c.capacity
}

// unit: bytes
func (c baseBloomFilter) ComputeStoreSize() int64 {
	return int64(math.Ceil(float64(c.bloomSize) / 8.0))
}

func (c *baseBloomFilter) initParameters(capacity int64, errorRate float64) {
	var (
		n int64   = capacity
		k float64 = -math.Log2(errorRate)
		m float64 = float64(n) * k / math.Log(2)

		hashNum      = int32(math.Ceil(k))
		bloomSize    = int32(math.Ceil(m))
		containerNum = int32(math.Ceil(float64(bloomSize) / float64(c.container.GetMaxBitSize())))
	)
	c.capacity = capacity
	c.errorRate = errorRate
	c.bloomSize = bloomSize
	c.hashNum = hashNum
	c.containerNum = containerNum
}

func (c *baseBloomFilter) Reset() {
	c.container.Reset()
}
