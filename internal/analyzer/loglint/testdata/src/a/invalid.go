package a

import (
	"fmt"
)

func f() {
	fmt.Println("hello") // want "avoid fmt.Println in production code"
}
