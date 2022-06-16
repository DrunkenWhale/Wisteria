package test

import (
	"Gloxinia/transform"
	"bytes"
	"testing"
)

func TestFloat_(t *testing.T) {
	n1 := transform.BytesTransformToFloat([]byte("114514.1919810")) + 1919810.114514
	if n1 != 2.034324306495e+06 {
		t.Error("BytesToFloat Method Get Wrong")
	}
	r1 := bytes.Compare(transform.FloatTransformToBytes(114514.191981), []byte("114514.191981"))
	if r1 != 0 {
		t.Error("FloatToBytes Method Get Wrong")
	}
	if 1919810.114514 != transform.BytesTransformToFloat(transform.FloatTransformToBytes(1919810.114514)) {
		t.Error("FloatToBytes or BytesToFloat Get Wrong")
	}
}
