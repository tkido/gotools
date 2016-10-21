package main

import (
	"fmt"

	"./bitset"
	"./log"
)

func main() {
	defer log.Close()

	log.D("Hello Logger!!")
	log.D(1)
	log.D(2, 3, 4)
	log.D("test")

	var b bitset.BitSet
	fmt.Println(b)
}
