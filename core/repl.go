package core

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

const (
	PUT    int8 = 1
	GET    int8 = 2
	UPDATE int8 = 3
	DELETE int8 = 4
)

const (
	LIST int8 = 5
	MAP  int8 = 6
)

type repl struct {
	wisteria *Wisteria
}

func NewRepl(wisteria *Wisteria) *repl {
	return &repl{
		wisteria: wisteria,
	}
}

// There exist and only exist some operation
// PUT/GET/DELETE/UPDATE a STRING/INT/FLOAT/MAP/LIST
// specially, we must analyze LIST Element to know its type
// and make sure whether MAP's kvs' type are string

func (r *repl) analyze(sentence string) error {

	str := strings.TrimSpace(sentence)
	strArray := strings.SplitN(str, " ", 3)
	op, keyString, after := strArray[0], strArray[1], strArray[2]
	key, err := strconv.ParseInt(keyString, 10, 64)
	if err != nil {
		return err
	}
	switch op {
	case "put":
	case "get":
	case "update":
	case "delete":
	case "append":
	case "flush":
	}
	log.Println(strArray)
	return nil
}

func (r *repl) putOperation(key int64, arr []string) error {
	if len(arr) != 1 {
		return errors.New("too much param~")
	}
	if len(arr[0]) >= 2 && arr[0][0] == '"' && arr[0][len(arr[0])] == '"' {
		string_ := arr[0][1 : len(arr[0])-2]
		return r.wisteria.PutSingleString(int(key), string_)
	}
	float_, err := strconv.ParseFloat(arr[0], 64)
	if err == nil {
		return r.wisteria.PutSingleFloat(int(key), float_)
	}
	int_, err := strconv.ParseInt(arr[0], 10, 64)
	if err == nil {
		return r.wisteria.PutSingleInt(int(key), int_)
	}
	return errors.New("illegal param")
}

func (r *repl) getOperation(key int64) (byte, any, error) {
	return r.wisteria.Query(int(key))
}

func (r *repl) updateOperation(key int64, arr []string) error {
	return r.putOperation(key, arr)
}

func (r *repl) deleteOperation(key int64) error {
	return r.wisteria.Delete(int(key))
}

func (r *repl) appendOperation(key int64, arr []string) error {
	if len(arr) == 1 {
		float_, err := strconv.ParseFloat(arr[0], 64)
		if err == nil {
			return r.wisteria.AppendElementToList(int(key), float_)
		}
		int_, err := strconv.ParseInt(arr[0], 10, 64)
		if err == nil {
			return r.wisteria.AppendElementToList(int(key), int_)
		}
		string_ := arr[0]
		return r.wisteria.AppendElementToList(int(key), string_)
	}
	if len(arr) == 2 {
		return r.wisteria.AppendElementToMap(int(key), arr[0], arr[1])
	}
	return errors.New("unknown operator type")
}

func (r *repl) flushOperation() error {
	return r.wisteria.Flush()
}
