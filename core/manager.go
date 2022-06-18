package core

import (
	"Wisteria/transform"
	"errors"
	"fmt"
	"github.com/Pigeon377/Dandelion/lsm"
)

const (
	intIdent    = 4
	floatIdent  = 5
	stringIdent = 6
	listIdent   = 7
	mapIdent    = 8
)

func query(key int, lsm *lsm.LSM) (byte, any, error) {
	bytes, ok := lsm.Get(key)
	if !ok {
		return 0, nil, errors.New(fmt.Sprintf("key [%v] not exist", key))
	}
	operator := bytes[0]
	switch operator {
	case intIdent:
		return operator, transform.BytesTransformToInt(bytes[1:]), nil
	case floatIdent:
		return operator, transform.BytesTransformToFloat(bytes[1:]), nil
	case stringIdent:
		return operator, transform.BytesTransformToString(bytes[1:]), nil
	case listIdent:
		return operator, transform.BytesTransformToList(bytes[1:]), nil
	case mapIdent:
		return operator, transform.BytesTransformToMap(bytes[1:]), nil
	default:
		return 0, nil, errors.New("unknown type")
	}
}

func putSingleFloat(key int, intNumber int64, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{intIdent}, transform.IntTransformToBytes(intNumber)...))
}

func putSingleInt(key int, floatNumber float64, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{floatIdent}, transform.FloatTransformToBytes(floatNumber)...))
}

func putSingleString(key int, str string, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{stringIdent}, transform.StringTransformToBytes(str)...))
}

func putNewList(key int, arr []any, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{listIdent}, transform.ListTransformToBytes(arr)...))
}

func putNewMap(key int, m map[string]string, lsm *lsm.LSM) error {
	return lsm.Put(key, append([]byte{mapIdent}, transform.MapTransformToBytes(m)...))
}

func appendElementToList(key int, element any, lsm *lsm.LSM) error {
	bytes, ok := lsm.Get(key)
	if !ok {
		return errors.New(fmt.Sprintf("List Unexist In Key : [%v]", key))
	}
	transform.ListAppendInBytes(bytes, element)
	return nil
}

func appendElementToMap(key int, mapKey string, value string, lsm *lsm.LSM) error {
	bytes, ok := lsm.Get(key)
	if !ok {
		return errors.New(fmt.Sprintf("Map Unexist In Key : [%v]", key))
	}
	transform.MapAppendInBytes(bytes, mapKey, value)
	return nil
}
