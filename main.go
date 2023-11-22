package main

import (
	"fmt"

	"github.com/hunderaweke/bahirehasabgo/bahirehasab"
)

func main() {
	bh := bahirehasab.BahireHasab{
		Year: 2009,
	}
	fmt.Println(bh.Wenber())
}
