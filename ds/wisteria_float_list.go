package ds

import (
	bytes2 "bytes"
	"log"
	"strconv"
)

// WisteriaList
// list can't contain any data seq type
// it will cause sep ident conflict

// WisteriaFloatList
// ***************************
type WisteriaFloatList struct {
	data []*WisteriaFloat
}

func (wfl *WisteriaFloatList) ToBytes() []byte {
	res := make([]byte, len(wfl.data))
	for _, element := range wfl.data {
		res = append(res, element.ToBytes()...)
		res = append(res, listElementSep)
	}
	return res
}

func NewWisteriaFloatListFromBytes(bytes []byte) []float64 {
	res := make([]float64, 0)
	bytesArray := bytes2.Split(bytes, []byte{listElementSep})
	for i := 0; i < len(bytesArray); i++ {
		parseFloat, err := strconv.ParseFloat(string(bytesArray[i]), 64)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, parseFloat)
	}
	return res
}

// convert golang data type

func NewWisteriaFloatList(arr []float64) *WisteriaFloatList {
	res := make([]*WisteriaFloat, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = NewWisteriaFloat(arr[i])
	}
	return &WisteriaFloatList{
		data: res,
	}
}

func (wfl *WisteriaFloatList) Get() []float64 {
	res := make([]float64, len(wfl.data))
	for i := 0; i < len(res); i++ {
		res[i] = wfl.data[i].Get()
	}
	return res
}
