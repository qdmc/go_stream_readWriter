package go_stream_readWriter

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func ExampleNewBits() {
	bits := NewBits()
	bits.ParseByte(uint8(128))
	fmt.Println(bits.ToString())
	bits.SetBit(7, 1)
	fmt.Println(bits.ToByte())
	fmt.Println(bits.ToString())
	// output: 10000000
	// 129
	// 10000001
}

func ExampleBytesDivision() {
	bs := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	arr := BytesDivision(bs, 3)
	fmt.Println(len(bs))
	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))
	fmt.Println(len(arr[1]))
	fmt.Println(len(arr[2]))
	// output: 9
	// 3
	// 3
	// 3
	// 3

}
func ExampleNewReader() {
	// 测试用报文
	type testPacket struct {
		U8     byte
		U16    uint16
		U32    uint32
		U64    uint64
		StrLen uint16
		Str    string
	}
	p1 := testPacket{
		U8:  8,
		U16: 16,
		U32: 32,
		U64: 64,
		Str: "this is a test packet",
	}
	p1.StrLen = uint16(len([]byte(p1.Str)))
	var err error
	w := NewWriter()
	_, err = w.WriteUint8(p1.U8)
	_, err = w.WriteUint16(p1.U16)
	_, err = w.WriteUint32(p1.U32)
	_, err = w.WriteUint64(p1.U64)
	_, err = w.WriteUint16(p1.StrLen)
	_, err = w.WriteBytes([]byte(p1.Str))
	if err != nil {
		fmt.Println("writer_err")
	}
	buf := bytes.NewReader(w.BufBytes())
	p2 := new(testPacket)
	r := NewReader()
	p2.U8, err = r.ReadUint8(buf)
	p2.U16, err = r.ReadUint16(buf)
	p2.U32, err = r.ReadUint32(buf)
	p2.U64, err = r.ReadUint64(buf)
	strLen, err := r.ReadUint16(buf)
	if err != nil {
		fmt.Println("reader_err")
	}
	strBs, err := r.ReadBytes(buf, uint(strLen))
	if err != nil {
		fmt.Println("reader_err")
	}
	p2.Str = string(strBs)
	fmt.Println(p1.U8 == p2.U8)
	fmt.Println(p1.U16 == p2.U16)
	fmt.Println(p1.U32 == p2.U32)
	fmt.Println(p1.U64 == p2.U64)
	fmt.Println(p1.Str == p2.Str)
	// output: true
	// true
	// true
	// true
	// true
}

func ExampleNewWriter() {
	var err error
	order := binary.BigEndian
	w := NewWriter()
	w.SetOrder(order)
	_, err = w.WriteUint8(8)
	if err != nil {
		fmt.Println("writer_err")
	}
	// 每一次调用BufBytes(),会有buf.Reset()
	fmt.Println(w.BufBytes()[0])
	_, err = w.WriteUint16(16)
	if err != nil {
		fmt.Println("writer_err")
	}
	fmt.Println(order.Uint16(w.BufBytes()))
	_, err = w.WriteUint32(32)
	if err != nil {
		fmt.Println("writer_err")
	}
	fmt.Println(order.Uint32(w.BufBytes()))
	_, err = w.WriteUint64(64)
	if err != nil {
		fmt.Println("writer_err")
	}
	fmt.Println(order.Uint64(w.BufBytes()))
	_, err = w.WriteBytes([]byte("hello"))
	if err != nil {
		fmt.Println("writer_err")
	}
	fmt.Println(string(w.BufBytes()))
	// output: 8
	// 16
	// 32
	// 64
	// hello
}

func Test_bits(t *testing.T) {
	bits := NewBits()
	bits.ParseByte(uint8(128))
	if bits.ToString() != "10000000" {
		fmt.Println(bits.ToString())
		t.Fatal("to string error")
	}
	bits.SetBit(7, 1)
	if bits.ToByte() != 129 {
		fmt.Println(bits.ToString())
		t.Fatal("to string error")
	}
	if bits.ToString() != "10000001" {
		fmt.Println(bits.ToString())
		t.Fatal("to string error")
	}

}
