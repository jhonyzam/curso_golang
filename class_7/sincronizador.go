package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	inicio := time.Now()

	rand.Seed(inicio.UnixNano())

	var controle sync.WaitGroup

	for i := 0; i < 5; i++ {
		controle.Add(1)
		go executarTimeSleep(&controle)
	}

	controle.Wait()

	fmt.Printf("Finalizando em %s.\n", time.Since(inicio))
}

func executarTimeSleep(controle *sync.WaitGroup) {
	defer controle.Done()

	duracao := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Dormindo por %s... \n", duracao)
	time.Sleep(duracao)

}
