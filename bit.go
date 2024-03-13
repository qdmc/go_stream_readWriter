package go_stream_readWriter

import "fmt"

/*
BitsInterface   bit位操作接口
  - ParseByte    解析一个byte
  - SetBit       设置位数据,index:0~~7;val:0~~1
  - GetBit       返回位上的值(0,1);index:0~~7有效,超范围返回-1
  - ToByte       返回Byte
  - ToString     返回8位二进制的string
*/
type BitsInterface interface {
	ParseByte(u uint8)
	SetBit(index uint8, val int)
	GetBit(index uint8) int
	ToByte() uint8
	ToString() string
}

// NewBits     生成一个 BitsInterface
func NewBits() BitsInterface {
	return &defaultBits{bits: [8]int{0, 0, 0, 0, 0, 0, 0, 0}}
}

type defaultBits struct {
	bits [8]int
}

func (b *defaultBits) ParseByte(u uint8) {
	byteStr := fmt.Sprintf("%08b", u)
	for index, val := range byteStr {
		b.bits[index] = int(val)
	}
}

func (b *defaultBits) SetBit(index uint8, val int) {
	if index <= 7 && (val == 0 || val == 1) {
		b.bits[index] = val
	}
}

func (b *defaultBits) GetBit(index uint8) int {
	if index <= 7 {
		return b.bits[index]
	}
	return -1
}

func (b *defaultBits) ToByte() uint8 {
	res := uint8(0)
	for index, val := range b.bits {
		if val == 1 {
			if index == 7 {
				res += 1
			} else {
				moveBite := 7 - index
				res += uint8(1) << moveBite
			}
		}
	}
	return res
}

func (b *defaultBits) ToString() string {
	return fmt.Sprintf(
		"%d%d%d%d%d%d%d%d",
		b.bits[0],
		b.bits[1],
		b.bits[2],
		b.bits[3],
		b.bits[4],
		b.bits[5],
		b.bits[6],
		b.bits[7],
	)
}
