package go_stream_readWriter

import (
	"fmt"
	"testing"
)

func ExampleNewBits() {

}

func ExampleBytesDivision() {

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
