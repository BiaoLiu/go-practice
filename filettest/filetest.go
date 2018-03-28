package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main() {
	fmt.Println(filepath.Abs("."))

	fmt.Println(filepath.Abs(os.Args[0]))

}
