package container

type SetStatus int32

const (
	SetBitFailed SetStatus = 0
	SetBitOK     SetStatus = 1
	SetBitExists SetStatus = 2
)

type BFContainer interface {
	Export() (map[int32]map[int64]bool, error)
	Import(map[int32]map[int64]bool) error
	GetBit(conainterId int32, index int64) (bool, error)
	SetBit(conainterId int32, index int64) SetStatus

	GetSize() int64
	SetSize(int64)
	IncreaseSize()

	Reset() bool
	GetMaxBitSize() int64
}
