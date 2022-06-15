package ds

import (
	"Gloxinia/util"
	bytes2 "bytes"
	"log"
	"strconv"
)

const setElementSep = 3

// WisteriaSet
// set can't contain any data seq type
// it will cause sep ident conflict

// WisteriaIntSet
// ***************************
type WisteriaIntSet struct {
	data util.Set[int64]
}

func (wis *WisteriaIntSet) ToBytes() []byte {
	arr := wis.data.Export()
	res := make([]byte, len(arr))
	for _, element := range arr {
		res = append(res, NewWisteriaInt(element).ToBytes()...)
		res = append(res, setElementSep)
	}
	return res
}

func NewWisteriaIntSetFromBytes(bytes []byte) *util.Set[int64] {
	res := make([]int64, 0)
	bytesArray := bytes2.Split(bytes, []byte{setElementSep})
	for i := 0; i < len(bytesArray); i++ {
		parseInt, err := strconv.ParseInt(string(bytesArray[i]), 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, parseInt)
	}
	return util.NewSetWithParam[int64](res...)
}

// convert golang data type

func NewWisteriaIntSet(set util.Set[int64]) *WisteriaIntSet {
	return &WisteriaIntSet{
		data: set,
	}
}

func (wis *WisteriaIntSet) Get() util.Set[int64] {
	return wis.data
}
