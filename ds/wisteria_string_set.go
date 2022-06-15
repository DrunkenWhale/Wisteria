package ds

import (
	"Gloxinia/util"
	bytes2 "bytes"
)

// WisteriaSet
// set can't contain any data seq type
// it will cause sep ident conflict

// WisteriaStringSet
// ***************************
type WisteriaStringSet struct {
	data util.Set[string]
}

func (wis *WisteriaStringSet) ToBytes() []byte {
	arr := wis.data.Export()
	res := make([]byte, len(arr))
	for _, element := range arr {
		res = append(res, NewWisteriaString(element).ToBytes()...)
		res = append(res, setElementSep)
	}
	return res
}

func NewStringSetFromBytes(bytes []byte) *util.Set[string] {
	res := make([]string, 0)
	bytesArray := bytes2.Split(bytes, []byte{setElementSep})
	for i := 0; i < len(bytesArray); i++ {
		res = append(res, string(bytesArray[i]))
	}
	return util.NewSetWithParam[string](res...)
}

// convert golang data type

func NewWisteriaStringSet(set util.Set[string]) *WisteriaStringSet {
	return &WisteriaStringSet{
		data: set,
	}
}

func (wis *WisteriaStringSet) Get() util.Set[string] {
	return wis.data
}
