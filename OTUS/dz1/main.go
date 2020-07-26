package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "dz1 - ", log.Ltime | log.Ldate)
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err)
	} else {
		fmt.Println(t)
	}
}
