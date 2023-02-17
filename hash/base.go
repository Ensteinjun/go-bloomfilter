package hash

type BFHash func(hashId int32, value []byte) int64

func BFSipHash(hashId int32, value []byte) int64 {
	return int64(SipHash12(value, uint64(hashId), uint64(hashId)))
}
