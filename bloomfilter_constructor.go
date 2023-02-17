package bloomfilter

import (
	"os"

	"github.com/Ensteinjun/go-bloomfilter/container"
	"github.com/Ensteinjun/go-bloomfilter/hash"
)

var DefaultHash = hash.BFSipHash

type (
	Option func(*baseBloomFilter)

	baseBloomFilter struct {
		capacity  int64
		errorRate float64
		hashNum   int32
		bloomSize int32

		containerNum int32
		container    container.BFContainer
		computeHash  hash.BFHash
	}
)

func NewBloomFilter(capacity int64, errorRate float64, container container.BFContainer, opts ...Option) *baseBloomFilter {
	var bf = &baseBloomFilter{container: container, computeHash: DefaultHash}
	bf.initParameters(capacity, errorRate)
	for _, opt := range opts {
		opt(bf)
	}
	return bf
}

func NewMemoryBloomFilter(capacity int64, errorRate float64, opts ...Option) *baseBloomFilter {
	return NewBloomFilter(capacity, errorRate, container.NewMemoryContainer(), opts...)
}

func LoadBloomFilter(filename string, container container.BFContainer, opts ...Option) (*baseBloomFilter, error) {
	var bf = &baseBloomFilter{container: container, computeHash: DefaultHash}
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

func WithHashFunction(hashFunc hash.BFHash) Option {
	return func(bbf *baseBloomFilter) {
		bbf.computeHash = hashFunc
	}
}
