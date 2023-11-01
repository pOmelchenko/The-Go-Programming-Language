package main

import (
	"fmt"
	"os"
	"strings"
)

// Упражнение 1.2.
//
// Измените программу echo так, чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке.
func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
