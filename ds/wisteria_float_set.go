package ds

import (
	bytes2 "bytes"
	"log"
	"strconv"
)

// WisteriaSet
// set can't contain any data seq type
// it will cause sep ident conflict

// WisteriaFloatSet
// ***************************
type WisteriaFloatSet struct {
	data Set[float64]
}

func (wis *WisteriaFloatSet) ToBytes() []byte {
	arr := wis.data.Export()
	res := make([]byte, len(arr))
	for _, element := range arr {
		res = append(res, NewWisteriaFloat(element).ToBytes()...)
		res = append(res, setElementSep)
	}
	return res
}

func NewFloatSetFromBytes(bytes []byte) *Set[float64] {
	res := make([]float64, 0)
	bytesArray := bytes2.Split(bytes, []byte{setElementSep})
	for i := 0; i < len(bytesArray); i++ {
		parseFloat, err := strconv.ParseFloat(string(bytesArray[i]), 64)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, parseFloat)
	}
	return NewSetWithParam[float64](res...)
}

// convert golang data type

func NewWisteriaFloatSet(set Set[float64]) *WisteriaFloatSet {
	return &WisteriaFloatSet{
		data: set,
	}
}

func (wis *WisteriaFloatSet) Get() Set[float64] {
	return wis.data
}
