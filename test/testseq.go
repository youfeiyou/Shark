package main

import (
	"fmt"
	"shark/pkg/util"
)

func main() {
	for i := 0; i < 100; i++ {
		seq, _ := util.GenSeq(util.UINSEQ)
		fmt.Println(seq)
	}
}
