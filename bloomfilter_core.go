package bloomfilter

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/Ensteinjun/go-bloomfilter/container"
)

func (c *baseBloomFilter) computeHash(hashId int32, value []byte) int64 {
	var ret int64
	if c.hashFunc != nil {
		ret = c.hashFunc(hashId, value)
	} else {
		bs := sha256.Sum256(append(value, []byte(fmt.Sprintf("%d", hashId))...))
		ret, _ = strconv.ParseInt(fmt.Sprintf("%x", []byte{bs[7], bs[15], bs[23], bs[31]}), 16, 64)
	}
	return ret % int64(c.bloomSize)
}

func (c *baseBloomFilter) Add(value []byte) bool {
	if c.keySize >= c.capacity {
		return false
	}

	var exists bool = true
	for i := int32(0); i < c.hashNum; i++ {
		v := c.computeHash(i, value)

		containerId := int32(v / c.container.GetMaxBitSize())
		containerIndex := v % c.container.GetMaxBitSize()
		status := c.container.SetBit(containerId, containerIndex)
		if status == container.SetBitFailed {
			return false
		} else if status == container.SetBitOK {
			exists = false
		}

	}
	if !exists {
		c.keySize++ // add new key
	}
	return true
}

func (c *baseBloomFilter) Contains(value []byte) bool {
	for i := int32(0); i < c.hashNum; i++ {
		v := c.computeHash(i, value)

		containerId := int32(v / c.container.GetMaxBitSize())
		containerIndex := v % c.container.GetMaxBitSize()
		if !c.container.GetBit(containerId, containerIndex) {
			return false
		}
	}
	return true
}
