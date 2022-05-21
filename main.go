package main

import (
	"fmt"
	shortid "githud.com/onedss/shortid/core"
)

func main() {
	id := shortid.MustGenerate()
	fmt.Println("Hello", id)
}
