package core

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
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

func StartREPL() {
	r := NewRepl(NewWisteria())
	r.startRepl()
}

func (r *repl) startRepl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("wisteria>  ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
			break
		}
		r.analyze(data)
	}
}

// There exist and only exist some operation
// PUT/GET/DELETE/UPDATE a STRING/INT/FLOAT/MAP/LIST
// specially, we must analyze LIST Element to know its type
// and make sure whether MAP's kvs' type are string

func (r *repl) analyze(sentence string) {

	str := strings.TrimSpace(sentence)
	strArray := strings.SplitN(str, " ", 3)
	var after = ""
	var keyString = ""
	op := strArray[0]
	if len(strArray) > 1 {
		keyString = strArray[1]
	}
	if len(strArray) > 2 {
		after = strArray[2]
	}
	key, err := strconv.ParseInt(keyString, 10, 64)
	re := regexp.MustCompile(`[^\s"]+|"([^"]*)"`)
	arr := re.FindAllString(after, -1)
	if err != nil && op != "flush" {
		log.Println(err)
	}
	switch op {
	case "put":
		err := r.putOperation(key, arr)
		if err != nil {
			log.Println(err)
		}
	case "get":
		_, res, err := r.getOperation(key)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%v\n", res)
		}
		return
	case "update":
		err := r.updateOperation(key, arr)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("succeed")
		}
		return
	case "delete":
		err := r.deleteOperation(key)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("succeed")
		}
		return
	case "append":
		err := r.appendOperation(key, arr)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("succeed")
		}
		return
	case "flush":
		err := r.flushOperation()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("succeed")
		}
	default:
		log.Println("can't recognize syntax")
	}
}

func (r *repl) putOperation(key int64, arr []string) error {
	if len(arr) != 1 {
		return errors.New("too much param~")
	} else {
		if arr[0] == "LIST" {
			err := r.wisteria.PutNewList(int(key), []any{})
			if err != nil {
				log.Println(err)
			}
		}
		if arr[0] == "MAP" {
			err := r.wisteria.PutNewMap(int(key), map[string]string{})
			if err != nil {
				log.Println(err)
			}
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
		if len(arr[0]) >= 2 && arr[0][0] == '"' && arr[0][len(arr[0])] == '"' {
			string_ := arr[0][1 : len(arr[0])-2]
			return r.wisteria.AppendElementToList(int(key), string_)
		}
		float_, err := strconv.ParseFloat(arr[0], 64)
		if err == nil {
			return r.wisteria.AppendElementToList(int(key), float_)
		}
		int_, err := strconv.ParseInt(arr[0], 10, 64)
		if err == nil {
			return r.wisteria.AppendElementToList(int(key), int_)
		}
		return errors.New(fmt.Sprintf("unknown type %v", arr[0]))
	}
	if len(arr) == 2 {
		return r.wisteria.AppendElementToMap(int(key), arr[0], arr[1])
	}
	return errors.New("unknown operator type")
}

func (r *repl) flushOperation() error {
	return r.wisteria.Flush()
}
