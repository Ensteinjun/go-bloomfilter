package bloomfilter

import (
	"os"

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
	var bf = &baseBloomFilter{container: container, keySize: 0, hashFunc: hashFunc}
	bf.initParameters(capacity, errorRate)
	return bf
}

func NewMemoryBloomFilter(capacity int64, errorRate float64, hashFunc HashFunction) *baseBloomFilter {
	return NewBloomFilter(capacity, errorRate, container.NewMemoryContainer(), hashFunc)
}

func LoadBloomFilter(filename string, container container.BloomFilterContainer, hashFunc HashFunction) (*baseBloomFilter, error) {
	var bf = &baseBloomFilter{container: container, hashFunc: hashFunc}
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	err = bf.Load(reader)
	if err != nil {
		return nil, err
	}
	return bf, nil
}
