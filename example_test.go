package go_stream_readWriter

import (
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

}

func ExampleNewWriter() {

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
