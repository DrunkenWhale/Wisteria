package test

import (
	"Wisteria/transform"
	"testing"
)

func TestList_(t *testing.T) {
	var arr []any
	arr = append(arr, int64(114514))
	arr = append(arr, "This is a string")
	arr = append(arr, 1919810.114514)
	if "\u0004114514\u0003\u0006This is a string\u0003\u00051919810.114514\u0003" != string(transform.ListTransformToBytes(arr)) {
		t.Error("ListTransformToBytes Method Has Error")
	}
	array := transform.BytesTransformToList(transform.ListTransformToBytes(arr))
	if array[0] != int64(114514) && array[1] != "This is a string" && array[2] != 1919810.114514 {
		t.Error("BytesTransformToList Method Has Error")
	}

}
