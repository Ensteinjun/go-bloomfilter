package hash

import (
	"encoding/binary"
)

func rotl(x, b uint64) uint64 {
	return (x << b) | (x >> (64 - b))
}

func sipround(v0, v1, v2, v3 uint64) (uint64, uint64, uint64, uint64) {
	v0 += v1
	v1 = rotl(v1, 13)
	v1 ^= v0
	v0 = rotl(v0, 32)
	v2 += v3
	v3 = rotl(v3, 16)
	v3 ^= v2
	v0 += v3
	v3 = rotl(v3, 21)
	v3 ^= v0
	v2 += v1
	v1 = rotl(v1, 17)
	v1 ^= v2
	v2 = rotl(v2, 32)
	return v0, v1, v2, v3
}

func SipHash(in []byte, k0, k1 uint64, cround, dround int) uint64 {
	var (
		v0 uint64 = 0x736f6d6570736575
		v1 uint64 = 0x646f72616e646f6d
		v2 uint64 = 0x6c7967656e657261
		v3 uint64 = 0x7465646279746573
	)
	v3 ^= k1
	v2 ^= k0
	v1 ^= k1
	v0 ^= k0

	buf := in[0:]
	for len(buf) >= 8 {
		m := binary.LittleEndian.Uint64(buf[:8])
		v3 ^= m

		for i := 0; i < cround; i++ {
			v0, v1, v2, v3 = sipround(v0, v1, v2, v3)
		}

		v0 ^= m
		buf = buf[8:]
	}

	var b uint64 = uint64(len(in)) << 56
	switch len(buf) {
	case 7:
		b |= uint64(buf[6]) << 48
		fallthrough
	case 6:
		b |= uint64(buf[5]) << 40
		fallthrough
	case 5:
		b |= uint64(buf[4]) << 32
		fallthrough
	case 4:
		b |= uint64(buf[3]) << 24
		fallthrough
	case 3:
		b |= uint64(buf[2]) << 16
		fallthrough
	case 2:
		b |= uint64(buf[1]) << 8
		fallthrough
	case 1:
		b |= uint64(buf[0])
	}

	v3 ^= b

	for i := 0; i < cround; i++ {
		v0, v1, v2, v3 = sipround(v0, v1, v2, v3)
	}

	v0 ^= b
	v2 ^= 0xff

	for i := 0; i < dround; i++ {
		v0, v1, v2, v3 = sipround(v0, v1, v2, v3)
	}

	return v0 ^ v1 ^ v2 ^ v3
}

func SipHash12(in []byte, k0, k1 uint64) uint64 {
	return SipHash(in, k0, k1, 1, 2)
}

func SipHash24(in []byte, k0, k1 uint64) uint64 {
	return SipHash(in, k0, k1, 2, 4)
}
