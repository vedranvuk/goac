package main

import (
	"fmt"
	"os"
	"time"

	"github.com/vedranvuk/goac/kn5"
)

func main() {
	start := time.Now()
	f, err := kn5.Open("/home/vedran/Go/src/github.com/vedranvuk/goac/data/car.kn5")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ellapsed := time.Since(start)
	fmt.Println(f)
	fmt.Printf("Loaded file in %v.\n", ellapsed)
}
