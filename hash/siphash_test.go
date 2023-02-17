package hash_test

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/Ensteinjun/go-bloomfilter/hash"
)

func TestSIPHash(t *testing.T) {
	k := []byte("12345678")
	ki := binary.LittleEndian.Uint64(k)

	v := hash.SipHash24([]byte("hello world"), ki, ki)
	fmt.Printf("V: %x\n", v)
	v1 := hash.SipHash12([]byte("hello world"), ki, ki)
	fmt.Printf("V: %d\n", v1)
}
