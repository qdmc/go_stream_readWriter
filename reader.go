package go_stream_readWriter

import (
	"encoding/binary"
	"io"
)

/*
StreamReaderInterface         字节流阻塞读取接口
  - SetOrder        配置字节序,大小端,默认是binary.BigEndian(大端)
  - ReadUint8       读取一个Byte
  - ReadUint16      读取一个Uin16
  - ReadUint32      读取一个Uin32
  - ReadUint64      读取一个Uin64
  - ReadBytes       读取定长的bytes,
*/
type StreamReaderInterface interface {
	SetOrder(order binary.ByteOrder)
	ReadUint8(r io.Reader) (byte, error)
	ReadUint16(r io.Reader) (uint16, error)
	ReadUint32(r io.Reader) (uint32, error)
	ReadUint64(r io.Reader) (uint64, error)
	ReadBytes(r io.Reader, length uint) ([]byte, error)
}

// NewReader    生成一个 StreamReaderInterface
func NewReader() StreamReaderInterface {
	return &defaultReader{order: binary.BigEndian}
}

type defaultReader struct {
	order binary.ByteOrder
}

func (reader *defaultReader) SetOrder(order binary.ByteOrder) {
	reader.order = order
}

func (reader *defaultReader) ReadUint8(r io.Reader) (byte, error) {
	bs := make([]byte, 1)
	_, err := io.ReadFull(r, bs)
	if err != nil {
		return 0, err
	}
	return bs[0], nil
}

func (reader *defaultReader) ReadUint16(r io.Reader) (uint16, error) {
	bs := make([]byte, 2)
	_, err := io.ReadFull(r, bs)
	if err != nil {
		return 0, err
	}
	return reader.order.Uint16(bs), nil
}

func (reader *defaultReader) ReadUint32(r io.Reader) (uint32, error) {
	bs := make([]byte, 4)
	_, err := io.ReadFull(r, bs)
	if err != nil {
		return 0, err
	}
	return reader.order.Uint32(bs), nil
}

func (reader *defaultReader) ReadUint64(r io.Reader) (uint64, error) {
	bs := make([]byte, 8)
	_, err := io.ReadFull(r, bs)
	if err != nil {
		return 0, err
	}
	return reader.order.Uint64(bs), nil
}

func (reader *defaultReader) ReadBytes(r io.Reader, length uint) ([]byte, error) {
	if length == 0 {
		return nil, nil
	}
	bs := make([]byte, length)
	_, err := io.ReadFull(r, bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
