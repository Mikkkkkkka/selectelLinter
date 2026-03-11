package a

import "a/notfmt"

func Println(s string) {
	// placeholder
}

func g() {
	Println("hello")
	notfmt.Println("hello")
}
