package main

import (
	"Gloxinia/ds"
	"fmt"
)

const sep = 2

func main() {
	fmt.Println(test(ds.NewWisteriaInt(14)))
}

func test(w ds.Wisteria) bool {
	_, ok := w.(*ds.WisteriaInt)
	return ok
}
