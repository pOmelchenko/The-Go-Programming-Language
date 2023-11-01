package main

import (
	"fmt"
	"os"
)

// Упражнение 1.2.
//
// Измените программу echo так, чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке.
func main() {
	for inx, value := range os.Args[:] {
		fmt.Println(inx, "-", value)
	}
}
