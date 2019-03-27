package main

import (
	"errors"
	"fmt"
)

func makeThumbnails4(filenames []string) error {
	errorsCh := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errorsCh <- err
		}(f)
	}

	for range filenames {
		if err := <-errorsCh; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
	}
	return nil
}

func ImageFile(name string) (bool, error) {
	if name == "test" {
		return true, errors.New("not test")
	}
	return false, nil
}

func main() {
	filenames := []string{"test1", "test1", "test1", "test1", "test", "test1", "test1", "test1"}

	if err := makeThumbnails4(filenames); err != nil {
		fmt.Println("error:", err.Error())
	}
}
