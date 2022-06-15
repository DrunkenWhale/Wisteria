package transform

import (
	"log"
	"strconv"
)

func BytesTransformToInt(bytes []byte) int64 {
	parseInt, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return parseInt
}

func IntTransformToBytes(int_ int64) []byte {
	return []byte(strconv.FormatInt(int_, 10))
}
