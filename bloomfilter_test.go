package bloomfilter_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Ensteinjun/go-bloomfilter"
	"github.com/Ensteinjun/go-bloomfilter/container"
)

func TestBloomFilter(t *testing.T) {
	bf := bloomfilter.NewMemoryBloomFilter(100000, 0.001, nil)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	fmt.Printf("Exists[hello]: %v\n", bf.Contains([]byte("hello")))
	fmt.Printf("Exists[world]: %v\n", bf.Contains([]byte("world")))
	fmt.Printf("Exists[golang]: %v\n", bf.Contains([]byte("golang")))
	fmt.Printf("Capacity: %d\n", bf.Capacity())
	fmt.Printf("Size: %d\n", bf.Size())
	fmt.Printf("Error: %f\n", bf.Error())

	writer, err := os.OpenFile("./test.bf", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("write failed failed: %v", err)
	}

	bf.Save(writer)

	fmt.Println("Load Bloom Filter")
	bf, err = bloomfilter.LoadBloomFilter("./test.bf", container.NewMemoryContainer(), nil)
	if err != nil {
		log.Fatalf("load bloom filter failed: %v", err)
	}
	fmt.Printf("Exists[hello]: %v\n", bf.Contains([]byte("hello")))
	fmt.Printf("Exists[world]: %v\n", bf.Contains([]byte("world")))
	fmt.Printf("Exists[golang]: %v\n", bf.Contains([]byte("golang")))
	fmt.Printf("Capacity: %d\n", bf.Capacity())
	fmt.Printf("Size: %d\n", bf.Size())
	fmt.Printf("Error: %f\n", bf.Error())
}
