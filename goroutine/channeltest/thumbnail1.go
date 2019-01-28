package main

import (
	"errors"
	"fmt"
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func ImageFile(filename string) (string, error) {
	fmt.Println("image file", filename)

	if filename == "2.txt" {
		return "", errors.New("测试的异常")
	}
	return "new filename" + filename, nil
}

func makeThumbnails1(filenames []string) {
	for _, filename := range filenames {
		if _, err := ImageFile(filename); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, filename := range filenames {
		go ImageFile(filename)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})

	for _, filename := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(filename)
	}

	for range filenames {
		<-ch
	}

	fmt.Println("make thubnails end.")
}

func makeThumbnails4(filenames []string) error {
	errorChan := make(chan error)

	for _, filename := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errorChan <- err
		}(filename)
	}

	//这个程序有一个微秒的bug。当它遇到第一个非nil的error时会直接将error返回到调用方，使得没有一
	//个goroutine去排空errors channel。这样剩下的worker goroutine在向这个channel中发送值时，都会
	//永远地阻塞下去，并且永远都不会退出。这种情况叫做goroutine泄露(§8.4.4)，可能会导致整个程序
	//卡住或者跑出out of memory的错误。

	//解决办法
	//最简单的解决办法就是用一个具有合适大小的buffered channel，这样这些worker goroutine向channel
	//中发送测向时就不会被阻塞。(一个可选的解决办法是创建一个另外的goroutine，当main goroutine返回第一个错误的同时去排空channel)

	for range filenames {
		if err := <-errorChan; err != nil {
			return err
		}
	}
	fmt.Println("make thubnails end.")
	return nil
}

func makeThumbnails5(filenames []string) ([]string, error) {
	type item struct {
		filename string
		err      error
	}

	ch := make(chan item, len(filenames))

	for _, filename := range filenames {
		go func(f string) {
			var result item
			result.filename, result.err = ImageFile(f)
			ch <- result
		}(filename)
	}

	var newFilenames []string
	for range filenames {
		result := <-ch
		if result.err != nil {
			return nil, result.err
		}
		newFilenames = append(newFilenames, result.filename)
	}

	fmt.Println("make thubnails end.")

	return newFilenames, nil
}

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines

	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	filenames := []string{"1.txt", "2.txt", "3.txt"}
	_, err := makeThumbnails5(filenames)

	fmt.Println(err)
}
