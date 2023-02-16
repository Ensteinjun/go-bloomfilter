package container

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/redis/go-redis/v9"
)

type RedisContainer struct {
	rdb            redis.Cmdable
	containerIdMap map[int32]int32
	keyPrefix      string
	maxKeyNum      int32
	size           int64
}

func NewRedisContainer(opt *redis.Options, keyPrefix string, maxKeyNum int32) *RedisContainer {
	rdb := redis.NewClient(opt)
	return &RedisContainer{
		rdb: rdb, containerIdMap: make(map[int32]int32),
		keyPrefix: keyPrefix, maxKeyNum: maxKeyNum,
	}
}

func NewRedisClusterContainer(opt *redis.ClusterOptions, keyPrefix string, maxKeyNum int32) *RedisContainer {
	rdb := redis.NewClusterClient(opt)
	return &RedisContainer{
		rdb: rdb, containerIdMap: make(map[int32]int32),
		keyPrefix: keyPrefix, maxKeyNum: maxKeyNum,
	}
}

func (c *RedisContainer) getKey(containerId int32) string {
	if _, ok := c.containerIdMap[containerId]; !ok {
		c.containerIdMap[containerId] = containerId % c.maxKeyNum
		// metaKey := fmt.Sprintf("%s:meta", c.keyPrefix)
		// c.rdb.HSet(context.Background(), metaKey, containerId, c.containerIdMap[containerId])
	}
	return fmt.Sprintf("%s:%d", c.keyPrefix, c.containerIdMap[containerId])
}

func (c *RedisContainer) GetBit(containerId int32, val int64) (bool, error) {
	res, err := c.rdb.GetBit(context.Background(), c.getKey(containerId), val).Result()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (c *RedisContainer) SetBit(containerId int32, val int64) SetStatus {
	exists, err := c.GetBit(containerId, val)
	if err != nil {
		return SetBitFailed
	} else if exists {
		return SetBitExists
	}
	_, err = c.rdb.SetBit(context.Background(), c.getKey(containerId), val, 1).Result()
	if err != nil {
		return SetBitFailed
	} else {
		return SetBitOK
	}
}

func (c *RedisContainer) Export() (map[int32]map[int64]bool, error) {
	data := make(map[int32]map[int64]bool)
	for cid := range c.containerIdMap {
		val, err := c.rdb.Get(context.Background(), c.getKey(cid)).Result()
		if err != nil {
			return nil, err
		}
		data[cid] = make(map[int64]bool)
		for _, bit := range []byte(val) {
			data[cid][int64(bit)] = true
		}
	}
	return data, nil
}

func (c *RedisContainer) Import(data map[int32]map[int64]bool) error {
	for cid, containerData := range data {
		values := make([]int64, 0)
		for v := range containerData {
			values = append(values, v)
		}
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

		bitSize := int64(math.Ceil(float64(values[len(values)-1]+1)/8.0) * 8.0)
		byteData := make([]byte, bitSize)
		for _, v := range values {
			byteData[v] = 1
		}
		_, err := c.rdb.Set(context.Background(), c.getKey(cid), byteData, 0).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *RedisContainer) Reset() bool {
	for cid := range c.containerIdMap {
		c.rdb.Del(context.Background(), c.getKey(cid)).Result()
	}
	c.size = 0
	return true
}

func (c *RedisContainer) GetMaxBitSize() int64 {
	return 2 ^ 20 // 128 KB
}

func (c *RedisContainer) IncreaseSize() {
	c.size++
}

func (c *RedisContainer) GetSize() int64 {
	return c.size
}

func (c *RedisContainer) SetSize(size int64) {
	c.size = size
}
