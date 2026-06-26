package main

import (
	"fmt"
	"sync"
)

// zadanie 1
func sendGreeting(ch chan string) {
	ch <- "Привет из горутины!"
}

// zadanie 2
func squareWorker(num int, ch_2 chan int) {

	fmt.Println("num:=", num)
	ch_2 <- num * num
}

// zadanie 3
func emitReader(ch_3 chan int) {
	for i := 1; i <= 5; i++ {
		ch_3 <- i
	}
	close(ch_3)
}

// zadanie 4
func sumReader(ch_4 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for n := range ch_4 {
		sum += n
	}
	fmt.Println("Сумма:", sum)
}

// zadanie 5
func filterEven(input chan int, output chan int) {
	for n := range input {
		if n%2 == 0 {
			output <- n
		}
	}
	close(output)
}

// zadanie 6
func checkChannel(ch_6 chan string) {
	val, ok := <-ch_6
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Канал закрыт!")
	}
}

func main() {

	//zadanie 1
	ch_1 := make(chan string)

	go sendGreeting(ch_1)

	message := <-ch_1
	fmt.Println(message)

	//zadanie2
	ch_2 := make(chan int)

	go squareWorker(9, ch_2)
	result := <-ch_2
	fmt.Printf("Квадрат числа %d = %d\n", 9, result)

	//zadanie 3
	ch_3 := make(chan int)

	go emitReader(ch_3)
	for num := range ch_3 {
		fmt.Println(num)
	}

	//zadanie 4
	//Cоздайте функцию sumReader, которая принимает канал chan int.
	// Внутри функции в цикле for range читайте числа из канала и складывайте их в переменную sum. Когда канал закроется,
	// функция должна вывести итоговую сумму на экран. В main создайте канал, запустите эту функцию как горутину,
	// отправьте в канал 3 любых числа, после чего закройте канал.

	ch_4 := make(chan int)

	var wg_4 sync.WaitGroup
	wg_4.Add(1)
	go sumReader(ch_4, &wg_4)

	ch_4 <- 52
	ch_4 <- 20
	ch_4 <- 10
	close(ch_4)

	wg_4.Wait()

	//zadanie 5
	// Создайте функцию-фильтр filterEven, которая принимает два канала: input chan int и output chan int.
	// Функция должна в цикле читать числа из input и, если число чётное, отправлять его в output.
	// В main запустите эту горутину, отправьте в канал input числа от 1 до 10, закройте input, а затем прочитайте и выведите все отфильтрованные числа из канала output.

	input := make(chan int, 10)
	output := make(chan int, 10)

	go filterEven(input, output)
	for i := 1; i <= 10; i++ {
		input <- i
	}
	close(input)

	for n := range output {
		fmt.Println(n)
	}

	//zadanie 6
	ch_6 := make(chan string, 1)

	ch_6 <- "GOOOOAL!"
	checkChannel(ch_6)

	close(ch_6)
	checkChannel(ch_6)

	//zadanie 7
	ch_7 := make(chan string, 3)

	ch_7 <- "El_matadora"
	ch_7 <- "JG"
	ch_7 <- "HOK"

	// ch_7 <- "PPP" //deadlock!

	// go func() {
	fmt.Println(<-ch_7)
	fmt.Println(<-ch_7)
	fmt.Println(<-ch_7)
	// 	close(ch_7)
	// }()

}
