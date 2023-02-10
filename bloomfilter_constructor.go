package bloomfilter

import (
	"math"

	"github.com/Ensteinjun/go-bloomfilter/container"
)

type (
	HashFunction func(hashId int32, value []byte) int64

	baseBloomFilter struct {
		keySize  int64
		hashFunc HashFunction

		capacity     int64
		errorRate    float64
		hashNum      int32
		bloomSize    int32
		containerNum int32
		container    container.BloomFilterContainer
	}
)

func NewBloomFilter(capacity int64, errorRate float64, container container.BloomFilterContainer, hashFunc HashFunction) *baseBloomFilter {
	var (
		n int64   = capacity
		k float64 = -math.Log2(errorRate)
		m float64 = float64(n) * k / math.Log(2)

		hashNum      = int32(math.Ceil(k))
		bloomSize    = int32(math.Ceil(m))
		containerNum = int32(math.Ceil(float64(bloomSize) / float64(container.GetMaxBitSize())))
	)

	return &baseBloomFilter{
		capacity: capacity, errorRate: errorRate, container: container, keySize: 0,
		bloomSize: bloomSize, hashNum: hashNum, containerNum: containerNum, hashFunc: hashFunc,
	}
}

func NewMemoryBloomFilter(capacity int64, errorRate float64, hashFunc HashFunction) *baseBloomFilter {
	return NewBloomFilter(capacity, errorRate, container.NewMemoryContainer(), hashFunc)
}
