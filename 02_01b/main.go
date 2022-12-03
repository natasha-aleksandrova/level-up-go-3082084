package main

import (
	"flag"
	"log"
	"sync"
)

var wg = sync.WaitGroup{}

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			log.Printf("[G%d]:%s\n", i, message)
			wg.Done()
		}(i)
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()

	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}

	wg.Wait()
}
