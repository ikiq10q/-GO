package main

import (
	"fmt"
	"strings"
)

func main() {
	menu := map[string]float64{
		"эспрессо":       800,
		"латте":          1200,
		"фрапучино":      1100,
		"раф":            1300,
		"хачапури":       1500,
		"моти":           900,
		"напиток недели": 1190,
		"сироп":          200,
	}
	order := []string{}
	summa := 0.0
	var zapros string

	fmt.Println("SMART-MENU") //МЕНЮ

	for name, price := range menu {
		fmt.Printf("%s — %.0f тг\n", name, price)
	}

	// цикл выбора из меню и добавление в корзину
	for {
		fmt.Print("Выберете из меню: ")
		fmt.Scan(&zapros)
		zapros = strings.ToLower(zapros)
		if zapros == "exit" {
			break
		}
		price, key := menu[zapros]

		if !key {
			fmt.Println("К сожалению, этого блюда нет в меню")
			continue
		}

		order = append(order, zapros)
		summa += price
		fmt.Println("Добавлено:", zapros)
	}
	fmt.Println(order)
	GroupMap := make(map[string]int)
	for _, coffe := range order {
		GroupMap[coffe]++
	}

	fmt.Println(" CHEK") //ЧЕК

	for coffe, count := range GroupMap {
		total := menu[coffe] * float64(count)
		fmt.Printf("%s x%d — %.0f тг\n", coffe, count, total)
	}

	fmt.Printf("\nСумма без скидки: %.0f тг\n", summa)

	//СКИДКА
	Skidka := 0.0
	if summa > 5000 {
		Skidka = summa * 0.10
	}

	lastSkidka := summa - Skidka

	nds := lastSkidka * 0.12     //НДС
	FullChek := lastSkidka + nds //ФИНАЛЬНЫЙ ЧЕК

	fmt.Println("Скидка:", Skidka)
	fmt.Println("НДС ", nds)
	fmt.Println("Итого к оплате: ", FullChek)

}
