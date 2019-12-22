package kn5

import (
	"testing"
)

func TestKn5(t *testing.T) {
	f, err := Load("/home/vedran/Go/src/github.com/vedranvuk/goac/data/car.kn5")
	if err != nil {
		t.Fatal(err)
	}

	if err := Save("/home/vedran/Go/src/github.com/vedranvuk/goac/data/car.out.kn5", f); err != nil {
		t.Fatal(err)
	}
}
