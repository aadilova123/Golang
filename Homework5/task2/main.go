package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func boringGenerator(msg string) <-chan string {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s - %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(ch)
	}()

	return ch
}

func interestingGenerator(msg string) <-chan string {	
	ch := make(chan string)
	go func() { 
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
	
    
}

func merge(cs ...<-chan string) <-chan string { 
	
	var wg sync.WaitGroup

	merged := make(chan string, 100)
	
	wg.Add(len(cs))
	

	output := func(sc <-chan string) {
		
		for sqr := range sc {
			merged <- sqr
		}
		
		wg.Done()
	}
	
	
	for _, optChan := range cs {
		go output(optChan)
	}
	
	
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged

		
}

func main() {
	fannedInCh := merge(boringGenerator("Joe"), interestingGenerator("Ann"), boringGenerator("Еркебулан"), boringGenerator("Кирилл"),interestingGenerator("Ayana"))

	for v := range fannedInCh {
		fmt.Println(v)
	}

	fmt.Println("GOOD :)")
}






