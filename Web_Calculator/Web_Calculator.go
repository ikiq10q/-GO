package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем наш обработчик
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Server is listening on port 8080...")
	// Запускаем сервер
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Готовим переменные, как в обычном калькуляторе
	var a, b float64
	var op string

	// МАГИЯ ТУТ:
	// fmt.Fscan читает данные из входящего запроса (r.Body) точно так же,
	// как fmt.Scan читал их с клавиатуры.
	// Ожидается формат: "Число Число Операция" (например: 10 20 +)
	_, err := fmt.Fscan(r.Body, &a, &b, &op)

	if err != nil {
		fmt.Fprintf(w, "Ошибка! Отправьте данные в формате: 10 20 +")
		return
	}

	// Логика вычислений
	var result float64

	switch op {
	// TODO: Реализуйте вычитание (+)
	case "+":
		result = a + b

	// TODO: Реализуйте вычитание (-)
	case "-":
		result = a - b
	// TODO: Реализуйте умножение (*)
	case "*":
		result = a * b
		// TODO: Реализуйте деление (/) с проверкой на 0 (if b == 0)
	case "/":
		if b == 0 {
			fmt.Fprintf(w, "Не определенно")
			return
		}
		result = a / b

	default:
		fmt.Fprintf(w, "Неизвестная операция: %s", op)
		return
	}

	// Отправляем результат обратно в Postman
	fmt.Fprintf(w, "Результат: %f", result)
}
