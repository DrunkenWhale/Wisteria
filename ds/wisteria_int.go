package ds

import (
	"log"
	"strconv"
)

type WisteriaInt struct {
	data int64
}

func (wi *WisteriaInt) ToBytes() []byte {
	return []byte(strconv.FormatInt(wi.data, 10))
}

func NewWisteriaIntFromBytes(bytes []byte) *WisteriaInt {
	parseInt, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return NewWisteriaInt(parseInt)
}

// convert golang data type

func NewWisteriaInt(int_ int64) *WisteriaInt {
	return &WisteriaInt{
		data: int_,
	}
}

func (wi *WisteriaInt) Get() int64 {
	return wi.data
}
