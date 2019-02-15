package main

import (
	"fmt"
)

func handle(name string) string {
	if name == "" {
		name = "stranger"
	}
	fmt.Printf("name=%s\n", name)
	return "Hello, " + name + "!"
}
