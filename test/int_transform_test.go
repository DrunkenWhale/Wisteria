package test

import (
	"Gloxinia/transform"
	"bytes"
	"testing"
)

func TestInt_(t *testing.T) {
	n1 := transform.BytesTransformToInt([]byte("114514")) + 1919810
	if n1 != 114514+1919810 {
		t.Error("BytesToInt Method Get Wrong")
	}
	r1 := bytes.Compare(transform.IntTransformToBytes(114514), []byte("114514"))
	if r1 != 0 {
		t.Error("IntToBytes Method Get Wrong")
	}
	if 1919114514 != transform.BytesTransformToInt(transform.IntTransformToBytes(1919114514)) {
		t.Error("IntToBytes or BytesToInt Get Wrong")
	}
}
