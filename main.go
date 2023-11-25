package main

import (
	"fmt"
	"log"

	"github.com/hunderaweke/bahirehasab-go/bahirehasab"
)

func main() {
	bh := bahirehasab.BahireHasab{
		Year:   2015,
		Logger: *log.Default(),
	}
	fmt.Println(bh.Neneweh())
}
