package main

import (
	"fmt"
	"sync"
	"time"
)

func CapturarDados(s string, wg *sync.WaitGroup, threads *chan bool, ac chan *string) {
	defer func() {
		wg.Done()
		<-*threads
	}()

	*threads <- true

	time.Sleep(4 * time.Second)
	ac <- &s
	fmt.Println("Capturando...")
}

func ExecutarDados(ac chan *string, acReturn chan *[]string) {
	var ac2 []string
	cc := 0
	for {
		a, open := <-ac
		if open {
			ac2 = append(ac2, *a)
			cc++
			fmt.Printf("channel %d \n", cc)
		} else {
			acReturn <- &ac2
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	qtThreads := 2
	threads := make(chan bool, qtThreads)

	a := []string{"valor1", "valor2", "valor3", "valor4"}
	ac := make(chan *string, len(a))
	acReturn := make(chan *[]string, 1)

	for _, str := range a {
		wg.Add(1)
		go CapturarDados(str, &wg, &threads, ac)
	}

	go ExecutarDados(ac, acReturn)

	wg.Wait()

	close(ac)
	close(threads)

	p := <-acReturn
	close(acReturn)

	fmt.Println(p)

	fmt.Println("FINALIZOU")
}
