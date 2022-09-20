package main

import (
	"fmt"

	// Notice how the import is <module>/<subpkg>
	"github.com/jolheiser/myapp/fruit"
)

func main() {
	b := fruit.Banana{}

	fmt.Println(b) // banana
}
