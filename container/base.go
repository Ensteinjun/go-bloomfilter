package container

type SetStatus int32

const (
	SetBitFailed SetStatus = 0
	SetBitOK     SetStatus = 1
	SetBitExists SetStatus = 2
)

type BloomFilterContainer interface {
	GetBit(conainterId int32, index int64) bool
	SetBit(conainterId int32, index int64) SetStatus
	GetMaxBitSize() int64
	Reset() bool
	Export() map[int32]map[int64]bool
	Import(map[int32]map[int64]bool) bool
}
