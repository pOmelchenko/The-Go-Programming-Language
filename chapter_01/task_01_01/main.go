package main

import (
	"fmt"
	"os"
	"strings"
)

// Упражнение 1.1.
//
// Измените программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды.
func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
