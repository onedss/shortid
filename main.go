package main

import (
	"fmt"
	shortid "github.com/onedss/shortid/core"
)

func main() {
	id := shortid.MustGenerate()
	fmt.Println("Hello", id)
}
