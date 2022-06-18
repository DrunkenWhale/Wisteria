package test

import (
	"Wisteria/core"
	"testing"
)

func TestPutNewList(t *testing.T) {
	w := core.NewWisteria()
	err := w.PutNewList(114, []any{int64(114514), "1919810", 1919810.114514})
	if err != nil {
		t.Error(err)
	}
	code, r1_, err := w.Query(114)
	r1 := r1_.([]any)
	if !(code == 7 && r1[0] == int64(114514) && r1[1] == "1919810" && r1[2] == 1919810.114514 && err == nil) {
		t.Error("Query Method In List Type Get Wrong")
	}
}

func TestAppendList(t *testing.T) {
	w := core.NewWisteria()
	err := w.PutNewList(114, []any{})
	if err != nil {
		t.Error(err)
	}
	err = w.AppendElementToList(114, int64(114514))
	if err != nil {
		t.Error(err)
	}
	err = w.AppendElementToList(114, "1919810")
	if err != nil {
		t.Error(err)
	}
	err = w.AppendElementToList(114, 1919810.114514)
	if err != nil {
		t.Error(err)
	}
	code, r1_, err := w.Query(114)
	r1 := r1_.([]any)
	if !(code == 7 && r1[0] == int64(114514) && r1[1] == "1919810" && r1[2] == 1919810.114514 && err == nil) {
		t.Error("Query Method In List Type Get Wrong")
	}
}

func TestPutNewMap(t *testing.T) {
	w := core.NewWisteria()
	err := w.PutNewMap(114514, map[string]string{
		"114514":  "1919810",
		"1919810": "114514",
	})
	if err != nil {
		t.Error(err)
	}
	code, r1_, err := w.Query(114514)
	r1 := r1_.(map[string]string)
	if err != nil {
		t.Error(err)
	}
	if code != 8 {
		t.Error("Unknown Ident Type")
	}
	if res, ok := r1["114514"]; !(ok && res == "1919810") {
		t.Error("Map Put Method Get Wrong")
	}
	if res, ok := r1["1919810"]; !(ok && res == "114514") {
		t.Error("Map Put Method Get Wrong")
	}
}

func TestAppendMap(t *testing.T) {
	w := core.NewWisteria()
	err := w.PutNewMap(114514, map[string]string{})
	if err != nil {
		t.Error(err)
	}
	err = w.AppendElementToMap(114514, "114514", "1919810")
	if err != nil {
		t.Error(err)
	}
	err = w.AppendElementToMap(114514, "1919810", "114514")
	if err != nil {
		t.Error(err)
	}
	code, r1_, err := w.Query(114514)
	r1 := r1_.(map[string]string)
	if err != nil {
		t.Error(err)
	}
	if code != 8 {
		t.Error("Unknown Ident Type")
	}
	if res, ok := r1["114514"]; !(ok && res == "1919810") {
		t.Error("Map Put Method Get Wrong")
	}
	if res, ok := r1["1919810"]; !(ok && res == "114514") {
		t.Error("Map Put Method Get Wrong")
	}
}
