package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	bh "github.com/hunderaweke/bahirehasab-go/bahirehasab"
)

func main() {
	year := flag.Int("year", time.Now().Year()-8, "the year for calculations")
	// logger := flag.Int("log", log.New(), "the log level for debugging")
	bh := bh.BahireHasab{
		Year:   *year,
		Logger: *log.Default(),
	}
	beal, date := bh.AbiyTsom()
	fmt.Println(beal + " : " + date)
}
