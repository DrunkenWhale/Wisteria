package main

import (
	"Wisteria/ds"
	"fmt"
	"github.com/Pigeon377/Dandelion/lsm"
)

const sep = 2

func main() {
	fmt.Println(test(ds.NewWisteriaInt(14)))
	lsm.NewLSM().Put(114514, []byte("s"))
	lsm.NewLSM().ClearMemory()
	lsm.NewLSM().ClearFilter()
}

func test(w ds.Wisteria) bool {
	_, ok := w.(*ds.WisteriaInt)
	return ok
}
