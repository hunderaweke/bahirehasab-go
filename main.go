package main

import (
	"fmt"

	"github.com/hunderaweke/bahirehasabgo/bahirehasab"
)

func main() {
	bh := bahirehasab.BahireHasab{
		Year: 2006,
	}
	fmt.Println(bh.Abekte())
}

