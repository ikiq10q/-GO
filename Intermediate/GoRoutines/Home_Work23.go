package main

import (
	"fmt"
	"sync"
	"time"
)

// zadanie 1
func printInfo() {
	fmt.Println("Горутина запущена")
}

// zadanie 2
func sayHello(name string) {
	fmt.Printf("Привет! %s\n", name)
}

// zadanie 3
func heavyTask(wg *sync.WaitGroup) {
	fmt.Println("Выполняю задачу...")
	wg.Done()
}

// zadanie 4
func count(id int, wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Горутина %d: число %d\n", id, i)
	}
	wg.Done()
}

// zadanie 5
func calculateSquare(num int, wg *sync.WaitGroup) {
	square := num * num
	fmt.Printf("Квадрат числа %d равен: %d\n", num, square)
	wg.Done()
}

// zadanie 6
func checkStatus(url string, wg *sync.WaitGroup) {

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Сайт %s доступен\n", url)
	wg.Done()
}

// zadanie 7
func processData(id int, category string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Printf("Запуск процесса %d в категории %s\n", id, category)
	time.Sleep(delay)
	fmt.Printf("Процесс %d завершен\n", id)
	wg.Done()
}

func main() {
	//zadanie 1
	go printInfo()
	time.Sleep(time.Second)

	//zadanie 2
	go sayHello("Argyn")
	go sayHello("Ardak")
	go sayHello("Daniil")

	//zadanie 3
	var wg3 sync.WaitGroup
	wg3.Add(1)
	go heavyTask(&wg3)
	wg3.Wait()
	fmt.Println("Горутина заверершила выполнение")

	//zadanie 4
	var wg4 sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg4.Add(1)
		go count(i, &wg4)
	}
	wg4.Wait()
	fmt.Println("КОнец")

	//zadanie 5
	NumSqr := [4]int{2, 3, 4, 5}
	var wg5 sync.WaitGroup
	for _, n := range NumSqr {
		wg5.Add(1)
		go calculateSquare(n, &wg5)
	}
	wg5.Wait()
	fmt.Println("Конец")

	//zadanie 6
	sites := []string{
		"google.com",
		"yandex.ru",
		"go.dev",
	}
	var wg6 sync.WaitGroup

	for _, site := range sites {
		wg6.Add(1)

		go checkStatus(site, &wg6)
	}

	wg6.Wait()
	fmt.Println("Конец")

	//zadaie 7
	//Создайте функцию processData, которая принимает три параметра: id (int), category (string) и delay (time.Duration).
	//Внутри функции выведите сообщение: "Запуск процесса [ID] в категории [Category]".
	//Затем используйте time.Sleep(delay), чтобы имитировать длительную работу.
	//После паузы выведите: "Процесс [ID] завершен".
	//Обязательно вызовите wg.Done() в конце работы функции.
	//В main создайте sync.WaitGroup и запустите в цикле 3 горутины с разными ID и разным временем задержки (например, 500мс, 1с и 200мс). Категорию и время задержки запросите у пользователя.
	// Используйте wg.Wait(), чтобы программа не закрылась, пока самый "медленный" процесс не завершит свою работу.

	var wg7 sync.WaitGroup
	var category string
	fmt.Print("Введите категорию:")
	fmt.Scan(&category)

	delays := [3]time.Duration{
		500 * time.Millisecond,
		1 * time.Second,
		200 * time.Millisecond,
	}

	for i := 0; i < 3; i++ {
		wg7.Add(1)
		go processData(i+1, category, delays[i], &wg7)
	}

	wg7.Wait()
	fmt.Println("Конец")

}
