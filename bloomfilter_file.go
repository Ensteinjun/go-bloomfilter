package bloomfilter

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"sort"
)

const magicNum uint32 = 0x9587001

func (c *baseBloomFilter) Save(writer io.Writer) error {
	containers, err := c.container.Export()
	if err != nil {
		return err
	}

	containerIds := make([]int32, 0)
	for cid := range containers {
		containerIds = append(containerIds, cid)
	}

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[:4], magicNum)
	writer.Write(buf[:4])
	binary.LittleEndian.PutUint64(buf, uint64(c.capacity))
	writer.Write(buf)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(c.errorRate))
	writer.Write(buf)
	binary.LittleEndian.PutUint64(buf, uint64(c.keySize))
	writer.Write(buf)
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(containerIds)))
	writer.Write(buf[:4])

	sort.Slice(containerIds, func(i, j int) bool { return containerIds[i] < containerIds[j] })
	for _, cid := range containerIds {
		values := make([]int64, 0)
		for vid := range containers[cid] {
			values = append(values, vid)
		}
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

		binary.LittleEndian.PutUint32(buf[:4], uint32(cid))
		writer.Write(buf[:4])
		binary.LittleEndian.PutUint64(buf, uint64(len(values)))
		writer.Write(buf)
		for _, v := range values {
			binary.LittleEndian.PutUint64(buf, uint64(v))
			writer.Write(buf)
		}
	}
	return nil
}

func (c *baseBloomFilter) Load(reader io.Reader) error {
	buf := make([]byte, 8)
	n, err := reader.Read(buf[:4])
	if err != nil || n != 4 {
		return fmt.Errorf("unknown format, can't read file[n=%d]: %v", n, err)
	}
	if binary.LittleEndian.Uint32(buf[:4]) != magicNum {
		return errors.New("unknown format: magic num error")
	}

	n, err = reader.Read(buf)
	if err != nil || n != 8 {
		return fmt.Errorf("unknown format, can't read file[n=%d]: %v", n, err)
	}
	capacity := binary.LittleEndian.Uint64(buf)

	n, err = reader.Read(buf)
	if err != nil || n != 8 {
		return fmt.Errorf("unknown format, can't read file[n=%d]: %v", n, err)
	}
	errorRate := math.Float64frombits(binary.LittleEndian.Uint64(buf))
	c.initParameters(int64(capacity), errorRate)

	n, err = reader.Read(buf)
	if err != nil || n != 8 {
		return fmt.Errorf("unknown format, can't read file[n=%d]: %v", n, err)
	}
	c.keySize = int64(binary.LittleEndian.Uint64(buf))

	n, err = reader.Read(buf[:4])
	if err != nil || n != 4 {
		return fmt.Errorf("unknown format, can't read file[n=%d]: %v", n, err)
	}
	numContainer := int32(binary.LittleEndian.Uint32(buf[:4]))

	containers := make(map[int32]map[int64]bool)
	for i := int32(0); i < numContainer; i++ {
		n, err := reader.Read(buf[:4])
		if err != nil || n != 4 {
			return fmt.Errorf("file broken, can't read container data[n=%d]: %v", n, err)
		}
		cid := int32(binary.LittleEndian.Uint32(buf[:4]))
		containers[cid] = make(map[int64]bool)

		n, err = reader.Read(buf)
		if err != nil || n != 8 {
			return fmt.Errorf("file broken, can't read container data[n=%d]: %v", n, err)
		}
		numValue := int64(binary.LittleEndian.Uint64(buf))
		for j := int64(0); j < numValue; j++ {
			n, err := reader.Read(buf)
			if err != nil || n != 8 {
				return fmt.Errorf("file broken, can't read container data[n=%d]: %v", n, err)
			}
			v := int64(binary.LittleEndian.Uint64(buf))
			containers[cid][v] = true
		}
	}
	return c.container.Import(containers)
}
