package ds

import (
	bytes2 "bytes"
)

// WisteriaList
// list can't contain any data seq type
// it will cause sep ident conflict

// WisteriaStringList
// ***************************
type WisteriaStringList struct {
	data []*WisteriaString
}

func (wil *WisteriaStringList) ToBytes() []byte {
	res := make([]byte, len(wil.data))
	for _, element := range wil.data {
		res = append(res, element.ToBytes()...)
		res = append(res, listElementSep)
	}
	return res
}

func NewWisteriaStringListFromBytes(bytes []byte) []string {
	res := make([]string, 0)
	bytesArray := bytes2.Split(bytes, []byte{listElementSep})
	for i := 0; i < len(bytesArray); i++ {
		parseString := string(bytesArray[i])
		res = append(res, parseString)
	}
	return res
}

// convert golang data type

func NewWisteriaStringList(arr []string) *WisteriaStringList {
	res := make([]*WisteriaString, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = NewWisteriaString(arr[i])
	}
	return &WisteriaStringList{
		data: res,
	}
}

func (wil *WisteriaStringList) Get() []string {
	res := make([]string, len(wil.data))
	for i := 0; i < len(res); i++ {
		res[i] = wil.data[i].Get()
	}
	return res
}
