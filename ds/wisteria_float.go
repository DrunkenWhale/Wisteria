package ds

import (
	"fmt"
	"log"
	"strconv"
)

type WisteriaFloat struct {
	data float64
}

func (wf *WisteriaFloat) ToBytes() []byte {
	return []byte(fmt.Sprintf("%f", wf.data))
}

func NewWisteriaFloatFromBytes(bytes []byte) *WisteriaFloat {
	parseFloat, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		log.Fatalln(err)
	}
	return NewWisteriaFloat(parseFloat)
}

// convert golang data type

func NewWisteriaFloat(float_ float64) *WisteriaFloat {
	return &WisteriaFloat{
		data: float_,
	}
}

func (wf *WisteriaFloat) Get() float64 {
	return wf.data
}
