package ds

import (
	bytes2 "bytes"
)

const mapElementSep = 3

type WisteriaMap struct {
	data map[string]Wisteria
}

func (wm *WisteriaMap) ToBytes() []byte {
	res := make([]byte, 0)
	for k, v := range wm.data {
		res = append(res, k...)
		res = append(res, mapElementSep)
		res = append(res, switchWisteriaBasicTypeAndToBytes(v)...)
		res = append(res, mapElementSep)
	}
	return res
}

func NewWisteriaMapFromBytes(bytes []byte) *WisteriaMap {
	m := WisteriaMap{
		data: map[string]Wisteria{},
	}
	bytesArray := bytes2.Split(bytes, []byte{mapElementSep})
	for i := 0; i < len(bytesArray); i += 2 {
		k := string(bytesArray[i])
		v := switchWisteriaBasicTypeAndFromBytes(bytesArray[i+1])
		m.data[k] = v
	}
	return &m
}

// ident code
// 15 int
// 16 float
// 17 string

const (
	mapIntElementIdent    = 15
	mapFloatElementIdent  = 16
	mapStringElementIdent = 17
)

func switchWisteriaBasicTypeAndToBytes(wb Wisteria) []byte {
	var ok bool
	_, ok = wb.(*WisteriaInt)
	if ok {
		return append([]byte{mapIntElementIdent}, wb.ToBytes()...)
	}
	_, ok = wb.(*WisteriaFloat)
	if ok {
		return append([]byte{mapFloatElementIdent}, wb.ToBytes()...)
	}
	_, ok = wb.(*WisteriaString)
	if ok {
		return append([]byte{mapStringElementIdent}, wb.ToBytes()...)
	}
	panic("Illegal Ident!")
}

func switchWisteriaBasicTypeAndFromBytes(bytes []byte) Wisteria {
	switch bytes[0] {
	case mapIntElementIdent:
		return NewWisteriaIntFromBytes(bytes[1:])
	case mapFloatElementIdent:
		return NewWisteriaFloatFromBytes(bytes[1:])
	case mapStringElementIdent:
		return NewWisteriaStringFromBytes(bytes[1:])
	}
	panic("Illegal Ident!")
}
