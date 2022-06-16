package test

import (
	"Gloxinia/transform"
	"bytes"
	"testing"
)

func TestString_(t *testing.T) {
	n1 := transform.BytesTransformToString([]byte("1145141919810")) + "1919810"
	if n1 != "1145141919810"+"1919810" {
		t.Error("BytesToString Method Get Wrong")
	}
	r1 := bytes.Compare(transform.StringTransformToBytes("114514.191981"), []byte("114514.191981"))
	if r1 != 0 {
		t.Error("StringToBytes Method Get Wrong")
	}
	if "1919810.114514" != transform.BytesTransformToString(transform.StringTransformToBytes("1919810.114514")) {
		t.Error("StringToBytes or BytesToString Get Wrong")
	}
}
