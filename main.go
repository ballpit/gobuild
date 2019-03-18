package main

import (
	"log"

	"github.com/pkg/errors"
)

func main() {
	myError := errors.Wrap(nil, "just an error")
	log.Fatal(myError)
}
