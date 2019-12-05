package kn5

import (
	"log"
	"testing"
)

func TestKn5(t *testing.T) {
	f, err := Open("/home/vedran/Go/src/github.com/vedranvuk/goac/data/car.kn5")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(f)
}
