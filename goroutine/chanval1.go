package main

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				count := elem["count"]
				count++
			} else {
				break
			}
		}

		syncChan <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			count := make(map[string]int)
			count["count"] = i
			mapChan <- count
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}
