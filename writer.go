package go_stream_readWriter

import (
	"bytes"
	"encoding/binary"
)

/*
StreamWriterInterface     字节流写接口
  - SetOrder        配置字节序,大小端,默认是binary.BigEndian(大端)
  - BufLen()        返回buf的长度
  - BufBytes()      返回buf的字节流并对buf做Reset
  - WriteUint8()    写入一个Byte
  - WriteUint16()   写入一个uint16
  - WriteUint32()   写入一个uint32
  - WriteUint64()   写入一个uint64
  - WriteBytes()    写入[]byte
*/
type StreamWriterInterface interface {
	SetOrder(order binary.ByteOrder)
	BufLen() int
	BufBytes() []byte
	WriteUint8(u uint8) (int, error)
	WriteUint16(u uint16) (int, error)
	WriteUint32(u uint32) (int, error)
	WriteUint64(u uint64) (int, error)
	WriteBytes(bs []byte) (int, error)
}

// NewWriter     生成一个 StreamWriterInterface
func NewWriter() StreamWriterInterface {
	return &defaultWriter{
		order: binary.BigEndian,
		buf:   new(bytes.Buffer),
	}
}

type defaultWriter struct {
	order binary.ByteOrder
	buf   *bytes.Buffer
}

func (w *defaultWriter) SetOrder(order binary.ByteOrder) {
	w.order = order
}

func (w *defaultWriter) BufLen() int {
	return w.buf.Len()
}

func (w *defaultWriter) BufBytes() []byte {
	bs := w.buf.Bytes()
	w.buf.Reset()
	return bs
}

func (w *defaultWriter) WriteUint8(u uint8) (int, error) {
	return 1, w.buf.WriteByte(u)
}

func (w *defaultWriter) WriteUint16(u uint16) (int, error) {
	bs := make([]byte, 2)
	w.order.PutUint16(bs, u)
	return w.buf.Write(bs)
}

func (w *defaultWriter) WriteUint32(u uint32) (int, error) {
	bs := make([]byte, 4)
	w.order.PutUint32(bs, u)
	return w.buf.Write(bs)
}

func (w *defaultWriter) WriteUint64(u uint64) (int, error) {
	bs := make([]byte, 8)
	w.order.PutUint64(bs, u)
	return w.buf.Write(bs)
}

func (w *defaultWriter) WriteBytes(bs []byte) (int, error) {
	if bs == nil || len(bs) == 0 {
		return 0, nil
	}
	return w.buf.Write(bs)
}
