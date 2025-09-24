package main

import (
	"sync"

	singleton_ "github.com/HGalassi/patterns/cmd/singleton"
)

func main() {
	//call sync singleton
	s1 := singleton_.GetInstance_example_1()
	s2 := singleton_.GetInstance_example_2()
	s3 := singleton_.GetInstance_example_3()
	var s4, s5 map[string]string
	//call without singleton
	for i := 0; i < 2; i++ {
		s4 = singleton_.NewMap()
		s5 = singleton_.NewMap()

	}
	println(s1, s2, s3, s4, s5)

	//async call
	var wg sync.WaitGroup
	for x := 0; x < 100; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			tempvar := singleton_.NewMap()
			println(tempvar, x)
		}(x)
	}
	wg.Wait() // Aguarda todas as goroutines terminarem

}
