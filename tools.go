// Package go_stream_readWriter  一个读写报文的小工具
package go_stream_readWriter

/*
BytesDivision           bytes定长分割
  - bs                  []byte
  - buffLen             分割长度
*/
func BytesDivision(bs []byte, buffLen int) [][]byte {
	var list [][]byte
	if bs == nil || len(bs) < 1 {
		return list
	}
	if buffLen < 1 || len(bs) <= buffLen {
		list = append(list, bs)
		return list
	}
	for len(bs) > buffLen {
		list = append(list, bs[0:buffLen])
		bs = bs[buffLen:]
	}
	if bs != nil && len(bs) > 0 {
		list = append(list, bs)
	}
	return list
}
