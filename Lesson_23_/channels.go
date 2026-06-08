package main

import (
	"fmt"
	"sync"
)

func main() {
	// 	ch := make(chan int, 3)
	// 	// go square(ch)

	// 	// fmt.Println("main func")
	// 	ch <- -5
	// 	// ch <- 52
	// 	// ch <- 56

	// 	fmt.Println(<-ch)
	// 	fmt.Println("main func ended ")

	// }

	// func square(ch chan int) {
	// 	fmt.Println("go func")
	// 	x := <-ch
	// 	ch <- x * x
	// 	fmt.Println("func ended ")
	//}
	// ch := make(chan int, 3)

	// // go square(ch)

	// fmt.Println("main func")
	// // ch <- 42
	// ch <- 45
	// ch <- 15
	// ch <- 45
	// ch <- 15 //

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// fmt.Println("main func ended")

	counter := 0
	ch := make(chan int, 1)
	ch <- counter
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(ch, &wg)

	}
	wg.Wait()
	fmt.Println(<-ch)

}

func increment(ch chan int, wg *sync.WaitGroup) {
	x := <-ch
	x++
	ch <- x
	defer func() {
		fmt.Println(x)
		wg.Done()
	}()
}

// func square(ch chan int) {
// 	fmt.Println("go func")
// 	x := <-ch
// 	ch <- x * x
// 	fmt.Println("go func ended")
// }
