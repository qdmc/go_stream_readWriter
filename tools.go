// Package go_stream_readWriter  一个读写报文的小工具
package go_stream_readWriter

/*
BytesDivision           bytes定长分割
  - bs                  []byte
  - buffLen             分割长度
*/
func BytesDivision(bs []byte, buffLen int) [][]byte {
	var list [][]byte
	if len(bs) < 1 {
		return list
	}
	if buffLen < 1 || len(bs) < buffLen {
		list = append(list, bs)
		return list
	}
	start := 0
	for start < len(bs) {
		end := start + buffLen
		if end >= (len(bs) - 1) {
			list = append(list, bs[start:])
		} else {
			list = append(list, bs[start:end])

		}
		start += buffLen
	}
	return list
}
