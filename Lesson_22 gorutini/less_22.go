package main

import (
	
)

func count_sum(a, b int, sum *int) {
	*sum = 0
	for i := a; i < b; i++ {
		*sum += i
	}

	fmt.Println(*sum)
}

func main() {
	var sum1, sum2, sum3 int = 0, 0, 0

	go count_sum(0, 1000000/3, &sum1)
	go count_sum(1000000/3, 2*1000000/3, &sum2)
	go count_sum(2*1000000/3, 1000000, &sum3)

	time.Sleep(20 * time.Second)
	fmt.Println(sum1 + sum2 + sum3)


func wait(minutes int, wg *sync.WaitGroup) {
	now := time.Now()
	target := time.Date(now.Year(), now.Month(), now.Day(), 20, minutes, 0, 0, now.Location())
	for {
		now := time.Now()
		if now.After(target) {
			fmt.Printf("It's now %d\n", minutes)
			break
		}
		fmt.Println("waite")
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

// 	fmt.Println(sum)
// }

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("We started")

	go wait(58, &wg)
	fmt.Println("added 1")
	go wait(59, &wg)
	fmt.Println("added 2")

	wg.Wait()
	fmt.Println("We ended")
}

// func count_sum(a, b int) {
// 	sum := 0
// 	for i := a; i < b; i++ {
// 		sum += i
// 	}

// 	fmt.Println(sum)
// }

func main() {
	for i := 1; i < 7; i++ {
		go func(n int) {
			result := 0
			for j := 1; j <= n; j++ {
				result += j
			}
			fmt.Println(n, "-", result)
		}(i)
	}

	time.Sleep(time.Second)
}
