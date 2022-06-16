package transform

import bytes2 "bytes"

const (
	listElementSep = 3
	intIdent       = 4
	floatIdent     = 5
	stringIdent    = 6
)

var listElementSepByteArray = []byte{listElementSep}

// ListTransformToBytes
// ensure slice arr only contains
// int64 / float64 / string
func ListTransformToBytes(arr []any) []byte {
	res := make([]byte, 0)
	for i := 0; i < len(arr); i++ {
		res = append(res, judgeElementTypeAndToBytes(arr[i])...)
		res = append(res, listElementSep)
	}
	return res
}

func judgeElementTypeAndToBytes(x any) []byte {
	var ok bool
	_, ok = x.(int64)
	if ok {
		return append([]byte{intIdent}, IntTransformToBytes(x.(int64))...)
	}
	_, ok = x.(float64)
	if ok {
		return append([]byte{floatIdent}, FloatTransformToBytes(x.(float64))...)
	}
	_, ok = x.(string)
	if ok {
		return append([]byte{stringIdent}, StringTransformToBytes(x.(string))...)
	}
	panic("Illegal Type!")
}

func judgeElementTypeAndFromBytes(x []byte) any {
	if x[0] == intIdent {
		return BytesTransformToInt(x[1:])
	}
	if x[0] == floatIdent {
		return BytesTransformToFloat(x[1:])
	}
	if x[0] == stringIdent {
		return BytesTransformToString(x[1:])
	}
	panic("Illegal Ident!")
}

func BytesTransformToList(bytes []byte) []any {
	bytesArray := bytes2.Split(bytes, listElementSepByteArray)
	res := make([]any, 0)
	for i := 0; i < len(bytesArray)-1; i += 1 {
		res = append(res, judgeElementTypeAndFromBytes(bytesArray[i]))
	}
	return res
}
