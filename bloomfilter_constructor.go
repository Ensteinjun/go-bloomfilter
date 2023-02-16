package bloomfilter

import (
	"os"

	"github.com/Ensteinjun/go-bloomfilter/container"
)

type (
	HashFunction func(hashId int32, value []byte) int64
	BFOption     func(*baseBloomFilter)

	baseBloomFilter struct {
		hashFunc HashFunction

		capacity     int64
		errorRate    float64
		hashNum      int32
		bloomSize    int32
		containerNum int32
		container    container.BFContainer
	}
)

func NewBloomFilter(capacity int64, errorRate float64, container container.BFContainer, opts ...BFOption) *baseBloomFilter {
	var bf = &baseBloomFilter{container: container}
	bf.initParameters(capacity, errorRate)
	for _, opt := range opts {
		opt(bf)
	}
	return bf
}

func NewMemoryBloomFilter(capacity int64, errorRate float64, opts ...BFOption) *baseBloomFilter {
	return NewBloomFilter(capacity, errorRate, container.NewMemoryContainer(), opts...)
}

func LoadBloomFilter(filename string, container container.BFContainer, opts ...BFOption) (*baseBloomFilter, error) {
	var bf = &baseBloomFilter{container: container}
	for _, opt := range opts {
		opt(bf)
	}
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

func WithHashFunction(hashFunc HashFunction) BFOption {
	return func(bbf *baseBloomFilter) {
		bbf.hashFunc = hashFunc
	}
}
