package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("2.pool.ntp.org")

	if err != nil {
		fmt.Println("Error -> main: ", err)
		os.Exit(1)
	}

	curhour, curmin, cursec := time.Clock()
	fmt.Printf("Current time: [%d:%d:%d]\n", curhour, curmin, cursec)

	os.Exit(0)
}
