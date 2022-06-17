package core

import (
	"Wisteria/transform"
	"github.com/Pigeon377/Dandelion/lsm"
)

const (
	intIdent    = 4
	floatIdent  = 5
	stringIdent = 6
	listIdent   = 7
	mapIdent    = 8
)

func putSingleFloat(key int, intNumber int64, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{intIdent}, transform.IntTransformToBytes(intNumber)...))
}
func putSingleInt(key int, floatNumber float64, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{floatIdent}, transform.FloatTransformToBytes(floatNumber)...))
}
func putSingleString(key int, str string, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{stringIdent}, transform.StringTransformToBytes(str)...))
}
