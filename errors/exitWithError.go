package errors

import (
	"fmt"
	"os"
)

// Завершение программы с выводом об ошибке
func ExitWithError(err error) {
	fmt.Fprintln(os.Stderr, "Ошибка:", err)
	os.Exit(2)
}
