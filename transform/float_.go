package transform

import (
	"fmt"
	"log"
	"strconv"
)

func BytesTransformToFloat(bytes []byte) float64 {
	parseFloat, err := strconv.ParseFloat(string(bytes), 64)
	if err != nil {
		log.Fatalln(err)
	}
	return parseFloat
}

func FloatTransformToBytes(float_ float64) []byte {
	return []byte(fmt.Sprintf("%f", float_))
}
