package main

import (
	"fmt"

	"github.com/hunderaweke/bahirehasab-go/bahirehasab"
)

func main() {
	bh := bahirehasab.BahireHasab{
		Year: 2016,
	}
	fmt.Println(bh.Tnsae())
}
