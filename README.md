# go-bloomfilter
[![GoDoc](https://godoc.org/github.com/Ensteinjun/go-bloomfilter?utm_source=godoc)](https://godoc.org/github.com/Ensteinjun/go-bloomfilter)

bloomfilter for golang

## Containers
```golang
// Implement this interface and make a custom container
bloomfilter.BFContainer

// Support
bloomfilter.NewMemoryBloomFilter
bloomfilter.NewRedisContainer
bloomfilter.NewRedisClusterContainer
```


## Example
```golang
// option 1: memorty bloom filter
bf := bloomfilter.NewMemoryBloomFilter(100000, 0.001)

// option 2: redis bloom filter
rdc := container.NewRedisContainer(&redis.Options{}, "bf_demo", 2)
bf := bloomfilter.NewBloomFilter(100000, 0.001, rdc)

// option 3: redis bloom filter
rdc1 := container.NewRedisClusterContainer(&redis.ClusterOptions{}, "bf_demo", 2)
bf := bloomfilter.NewBloomFilter(100000, 0.001, rdc1)

bf.Add([]byte("hello"))
bf.Add([]byte("world"))
fmt.Printf("Capacity: %d\n", bf.Capacity())
fmt.Printf("Size: %d\n", bf.Size())
fmt.Printf("Error: %f\n", bf.Error())
fmt.Printf("Exists[hello]: %v\n", bf.Contains([]byte("hello")))
fmt.Printf("Exists[world]: %v\n", bf.Contains([]byte("world")))
fmt.Printf("Exists[golang]: %v\n", bf.Contains([]byte("golang")))

// save bloomfilter to file
writer, err := os.OpenFile("./test.bf", os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
    log.Fatalf("write failed failed: %v", err)
}
bf.Save(writer)

/*
Load bloomfilter from File

It is necessary to ensure that the two values ​​of hashFunction and container.GetMaxBitSize() cannot be changed, otherwise the bloomfilter of load cannot work correctly
*/ 
fmt.Println("Load Bloom Filter")
bf2, err := bloomfilter.LoadBloomFilter("./test.bf", container.NewMemoryContainer(), nil)
if err != nil {
    log.Fatalf("load bloom filter failed: %v", err)
}
fmt.Printf("Exists[hello]: %v\n", bf2.Contains([]byte("hello")))
fmt.Printf("Exists[world]: %v\n", bf2.Contains([]byte("world")))
fmt.Printf("Exists[golang]: %v\n", bf2.Contains([]byte("golang")))
fmt.Printf("Capacity: %d\n", bf2.Capacity())
fmt.Printf("Size: %d\n", bf2.Size())
fmt.Printf("Error: %f\n", bf2.Error())
```

## Run Test
```shell
go test
```

## License
[MIT License Copyright (c) 2023 coderjun](http://opensource.org/licenses/MIT)