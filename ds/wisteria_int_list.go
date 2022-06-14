package ds

import (
	bytes2 "bytes"
	"log"
	"strconv"
)

const listElementSep = 3

// WisteriaList
// list can't contain any data seq type
// it will cause sep ident conflict

// WisteriaIntList
// ***************************
type WisteriaIntList struct {
	data []*WisteriaInt
}

func (wil *WisteriaIntList) ToBytes() []byte {
	res := make([]byte, len(wil.data))
	for _, element := range wil.data {
		res = append(res, element.ToBytes()...)
		res = append(res, listElementSep)
	}
	return res
}

func NewWisteriaIntListFromBytes(bytes []byte) []int64 {
	res := make([]int64, 0)
	bytesArray := bytes2.Split(bytes, []byte{listElementSep})
	for i := 0; i < len(bytesArray); i++ {
		parseInt, err := strconv.ParseInt(string(bytesArray[i]), 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, parseInt)
	}
	return res
}

// convert golang data type

func NewWisteriaIntFloatList(arr []int64) *WisteriaIntList {
	res := make([]*WisteriaInt, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = NewWisteriaInt(arr[i])
	}
	return &WisteriaIntList{
		data: res,
	}
}

func (wil *WisteriaIntList) Get() []int64 {
	res := make([]int64, len(wil.data))
	for i := 0; i < len(res); i++ {
		res[i] = wil.data[i].Get()
	}
	return res
}
