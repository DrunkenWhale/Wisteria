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

type Wisteria struct {
	l *lsm.LSM
}

func NewWisteria() *Wisteria {
	return &Wisteria{
		l: lsm.NewLSM(),
	}
}

func (w *Wisteria) Query(key int) (byte, any, error) {
	bytes, ok := w.l.Get(key)
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

func (w *Wisteria) PutSingleFloat(key int, intNumber int64) error {
	return w.l.Put(key, append([]byte{intIdent}, transform.IntTransformToBytes(intNumber)...))
}

func (w *Wisteria) PutSingleInt(key int, floatNumber float64) error {
	return w.l.Put(key, append([]byte{floatIdent}, transform.FloatTransformToBytes(floatNumber)...))
}

func (w *Wisteria) PutSingleString(key int, str string) error {
	return w.l.Put(key, append([]byte{stringIdent}, transform.StringTransformToBytes(str)...))
}

func (w *Wisteria) PutNewList(key int, arr []any) error {
	return w.l.Put(key, append([]byte{listIdent}, transform.ListTransformToBytes(arr)...))
}

func (w *Wisteria) PutNewMap(key int, m map[string]string) error {
	return w.l.Put(key, append([]byte{mapIdent}, transform.MapTransformToBytes(m)...))
}

func (w *Wisteria) AppendElementToList(key int, element any) error {
	bytes, ok := w.l.Get(key)
	if !ok {
		return errors.New(fmt.Sprintf("List Unexist In Key : [%v]", key))
	}
	return w.l.Put(key, transform.ListAppendInBytes(bytes, element))
}

func (w *Wisteria) AppendElementToMap(key int, mapKey string, value string) error {
	bytes, ok := w.l.Get(key)
	if !ok {
		return errors.New(fmt.Sprintf("Map Unexist In Key : [%v]", key))
	}
	return w.l.Put(key, transform.MapAppendInBytes(bytes, mapKey, value))
}
