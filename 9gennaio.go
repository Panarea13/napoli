package ints

import (
	"fmt"
	"os"
)

func Napoli() {
	fmt.Fprintln(os.Stdout, []any{"Napoli"}...)
}
