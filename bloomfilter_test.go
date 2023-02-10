package bloomfilter_test

import (
	"fmt"
	"testing"

	"github.com/ensteinjun/go-bloomfilter"
)

func TestBloomFilter(t *testing.T) {
	bf := bloomfilter.NewMemoryBloomFilter(100000, 0.001, nil)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	fmt.Printf("Capacity: %d\n", bf.Capacity())
	fmt.Printf("Size: %d\n", bf.Size())
	fmt.Printf("Error: %f\n", bf.Error())
	fmt.Printf("Exists[hello]: %v\n", bf.Contains([]byte("hello")))
	fmt.Printf("Exists[world]: %v\n", bf.Contains([]byte("world")))
	fmt.Printf("Exists[golang]: %v\n", bf.Contains([]byte("golang")))
}
