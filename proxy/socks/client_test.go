package socks_test

import (
	"fmt"
	"testing"

	. "github.com/xtls/xray-core/proxy/socks"
)

type Buffer struct {
	v     []byte
	start int32
	end   int32
}

func (b *Buffer) Bytes() []byte {
	return b.v[b.start:b.end]
}

func TestCompressSocks(t *testing.T) {
	data := []byte("example data 1bcdefghijklmnopqrstuvwxyz~!@#$%^&*()0987654321`")
	fmt.Printf("%#v\n", data)
	CompressSocks(data)
	fmt.Printf("%#v\n", data)

	DeCompressSocks(data)
	fmt.Println(string(data))

	ji := []byte{0x03, 0x01, 0xab}
	CompressSocks(ji)
	fmt.Printf("%#v\n", ji)
	DeCompressSocks(ji)
	fmt.Printf("%#v\n", ji)

	b := &Buffer{
		v:     []byte{0x01, 0x02, 0xa3, 0x04, 0x05, 0xa6, 0x07, 0x08, 0xa9},
		start: 0,
		end:   3,
	}
	fmt.Printf("%#v\n", b.Bytes())
	CompressSocks(b.Bytes())
	fmt.Printf("%#v\n", b.Bytes())
	fmt.Printf("%#v\n", b.v)
	DeCompressSocks(b.Bytes())
	fmt.Printf("%#v\n", b.Bytes())
}
