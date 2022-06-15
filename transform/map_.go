package transform

import bytes2 "bytes"

const mapElementSep = 3

var mapElementSepByteArray = []byte{mapElementSep}

func MapTransformToBytes(m map[string]string) []byte {
	res := make([]byte, 0)
	for k, v := range m {
		res = append(res, []byte(k)...)
		res = append(res, mapElementSep)
		res = append(res, []byte(v)...)
		res = append(res, mapElementSep)
	}
	return res
}

func BytesTransformToMap(bytes []byte) map[string]string {
	bytesArray := bytes2.Split(bytes, mapElementSepByteArray)
	m := make(map[string]string, 0)
	for i := 0; i < len(bytesArray); i += 2 {
		m[string(bytesArray[i])] = string(bytesArray[i+1])
	}
	return m
}
