package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsRune("woshishui", rune('w')))
	fmt.Println(strings.ContainsRune("wobushishui", rune('h')))
}
