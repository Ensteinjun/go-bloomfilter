package bloomfilter_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Ensteinjun/go-bloomfilter"
	"github.com/Ensteinjun/go-bloomfilter/container"
	"github.com/redis/go-redis/v9"
)

func TestBloomFilter(t *testing.T) {
	bf := bloomfilter.NewMemoryBloomFilter(100000, 0.001)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))

	fmt.Printf("Capacity: %d\n", bf.Capacity())
	fmt.Printf("Size: %d\n", bf.Size())
	fmt.Printf("Error: %f\n", bf.Error())

	exists, err := bf.Contains([]byte("hello"))
	fmt.Printf("Exists[hello]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("world"))
	fmt.Printf("Exists[world]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("golang"))
	fmt.Printf("Exists[golang]: %v, Err: %v\n", exists, err)

	writer, err := os.OpenFile("./test.bf", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("write failed failed: %v", err)
	}

	bf.Save(writer)

	fmt.Println("\nLoad Bloom Filter")
	bf2, err := bloomfilter.LoadBloomFilter("./test.bf", container.NewMemoryContainer())
	if err != nil {
		log.Fatalf("load bloom filter failed: %v", err)
	}

	fmt.Printf("Capacity: %d\n", bf2.Capacity())
	fmt.Printf("Size: %d\n", bf2.Size())
	fmt.Printf("Error: %f\n", bf2.Error())
	exists, err = bf.Contains([]byte("hello"))
	fmt.Printf("Exists[hello]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("world"))
	fmt.Printf("Exists[world]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("golang"))
	fmt.Printf("Exists[golang]: %v, Err: %v\n", exists, err)
}

func TestRedisBloomFilter(t *testing.T) {
	rdc := container.NewRedisContainer(&redis.Options{}, "bf_demo5", 2)
	bf := bloomfilter.NewBloomFilter(100000, 0.001, rdc)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))

	fmt.Printf("Capacity: %d\n", bf.Capacity())
	fmt.Printf("Size: %d\n", bf.Size())
	fmt.Printf("Error: %f\n", bf.Error())
	exists, err := bf.Contains([]byte("hello"))
	fmt.Printf("Exists[hello]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("world"))
	fmt.Printf("Exists[world]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("golang"))
	fmt.Printf("Exists[golang]: %v, Err: %v\n", exists, err)

	writer, err := os.OpenFile("./test.bf", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("write failed failed: %v", err)
	}
	bf.Save(writer)

	fmt.Println("\nLoad Bloom Filter")
	bf2, err := bloomfilter.LoadBloomFilter("./test.bf", container.NewRedisContainer(&redis.Options{}, "bf_demo2", 2))
	if err != nil {
		log.Fatalf("load bloom filter failed: %v", err)
	}

	fmt.Printf("Capacity: %d\n", bf2.Capacity())
	fmt.Printf("Size: %d\n", bf2.Size())
	fmt.Printf("Error: %f\n", bf2.Error())
	exists, err = bf.Contains([]byte("hello"))
	fmt.Printf("Exists[hello]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("world"))
	fmt.Printf("Exists[world]: %v, Err: %v\n", exists, err)
	exists, err = bf.Contains([]byte("golang"))
	fmt.Printf("Exists[golang]: %v, Err: %v\n", exists, err)
}
