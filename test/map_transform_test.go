package test

import (
	"Gloxinia/transform"
	"testing"
)

func TestMap_(t *testing.T) {
	m := map[string]string{}
	m["114514"] = "1919810"
	m["1919810"] = "114514"
	if string(transform.MapTransformToBytes(m)) != "114514\u00031919810\u00031919810\u0003114514\u0003" {
		t.Error("MapTransformToBytes Method Has Error")
	}
	m = transform.BytesTransformToMap(transform.MapTransformToBytes(m))
	if t1, ok := m["114514"]; !ok || t1 != "1919810" {
		t.Error("BytesTransformToMap Method has error")
	}
	if t2, ok := m["1919810"]; !ok || t2 != "114514" {
		t.Error("BytesTransformToMap Method has error")
	}
}
