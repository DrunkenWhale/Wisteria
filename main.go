package main

import (
	"github.com/Pigeon377/Dandelion/lsm"
)

const sep = 2

func main() {
	lsm.NewLSM().Put(114514, []byte("s"))
	lsm.NewLSM().ClearMemory()
	lsm.NewLSM().ClearFilter()
}
