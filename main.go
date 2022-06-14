package main

import (
	"github.com/Pigeon377/Dandelion/lsm"
)

const sep = 2

func main() {
	l := lsm.NewLSM()
	l.Put(11751, []byte("s"))
	println("Hello World!")
}
