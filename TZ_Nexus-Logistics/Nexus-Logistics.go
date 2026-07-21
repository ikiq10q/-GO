package main

import (
	"fmt"
	"strings"
	"time"
)

// 1.1. Константы и глобальные переменные
const BaseDeliveryRate = 10 / 2 //базовая стоимость/коэффициент энергозатрат.
const MaxBattareLevel = 100     //максимальный заряд батареи
const EmergencyThreshold = 20   // критический уровень энергии (20.0), ниже которого робот не может отправляться на задание.
const BufferCapacity = 5        //фиксированная емкость буфера заказов

// 1.2. Базовые типы данных
type RobotType string

const (
	Drone       RobotType = "Дрон"
	Wheeled     RobotType = "Колесная платформа"
	HeavyLifter RobotType = "Тяжелый погрузчик"
)

var AllRobotTypes = []RobotType{
	Drone,
	Wheeled,
	HeavyLifter,
}

type Product struct {
	ID     int
	Name   string
	Weight float64
}
type Robot struct {
	ID          int
	Model       string
	Type        RobotType //кастомный тип робота
	Battery     float64
	IsAvailable bool //статус занятости
}

// 2.1. Интерфейс Mover
type Mover interface {
	Move(distance float64) error
}

// метод moVe
func (r *Robot) Move(product *Product, distance float64) error {
	var CF float64 //CF = коэффициент

	switch r.Type {
	case Drone:
		CF = 1.5
	case Wheeled:
		CF = 0.8
	case HeavyLifter:
		CF = 2.5

	}
	energiya := distance * CF // необходимый заряд для преоделения

	if r.Battery < energiya {
		fmt.Printf("недостаточно заряда %s : нужно %.1f, есть %.1f\n", r.Model, energiya, r.Battery)
		return fmt.Errorf("недостаточно заряда %s: нужно %.1f, есть %.1f", r.Model, energiya, r.Battery)
	}

	r.Battery -= energiya
	fmt.Printf("%s переместился на %.1f м, потрачено %.1f заряда, осталось: %.1f\n",
		r.Model, distance, energiya, r.Battery)

	if r.Battery < EmergencyThreshold {
		err := string(fmt.Sprintf("Низкий уровень заряда: %.1f", r.Battery))
		fmt.Println(DeliveryReport(*r, product, err))
	}
	return nil
}

// 2.2. Функция Recharge
func Recharge(robot *Robot) {
	robot.Battery = MaxBattareLevel

}

// 2.3 Валидация и обработок строк
// грузоподъемность роботов
const (
	DroneMaxWeight       = 2.0 //kg
	WheeledMaxWeight     = 10.0
	HeavyLifterMaxWeight = 50.0
)

func (r *Robot) MaxLoadCarrying() float64 {
	switch r.Type {
	case Drone:
		return DroneMaxWeight
	case Wheeled:
		return WheeledMaxWeight
	case HeavyLifter:
		return HeavyLifterMaxWeight
	default:
		return 0
	}
}

// фунцкия проверки грузоподъемности
func (r *Robot) CanCarryMrRobot(product *Product) error {
	maxWeight := r.MaxLoadCarrying()

	if product.Weight > maxWeight {
		return fmt.Errorf(
			"%s не может вести торав %s весом %.1f kg: максгрузоподъемность %.1f kg",
			r.Model,
			product.Name,
			product.Weight,
			product.Weight,
		)
	}
	return nil

}

func DeliveryReport(robot Robot, product *Product, status string) string {
	finalStatus := strings.ToUpper(status)

	report := fmt.Sprintf(
		" Робот: %s\n  Тип: %s\n Товар: %s\n  Вес: %.1f\n kg  Статус: %s\n ",
		robot.Model,
		robot.Type,
		product.Name,
		product.Weight,
		finalStatus,
	)

	return strings.TrimSpace(report)
}

// zadanie 3 Concurrent
// 3.1. Функция доставки (DeliveryTask)

func DeliveryTask(robot *Robot, product *Product, distance float64, reportCh chan<- string) {

	time.Sleep(time.Duration(distance) * time.Millisecond * 10)

	fmt.Printf("%-15s готовится к доставке товара: %s\n", robot.Model, product.Name)

	// апроветка на грузоподъемность
	if err := robot.CanCarryMrRobot(product); err != nil {
		reportCh <- DeliveryReport(*robot, product, "ошибка: "+err.Error())
		return
	}

	//проверка на критический заряд батареии
	if robot.Battery < EmergencyThreshold {
		reportCh <- DeliveryReport(*robot, product, fmt.Sprintf("критический заряд: %.1f", robot.Battery))
		return
	}

	robot.IsAvailable = false
	// time.Sleep(time.Second)

	// выполнение, если прошел проверки выше
	if err := robot.Move(product, distance); err != nil {
		reportCh <- DeliveryReport(*robot, product, "ошибка: "+err.Error())
		robot.IsAvailable = true
		return
	}

	reportCh <- DeliveryReport(*robot, product, "Выполненно")

	robot.IsAvailable = true
}

func main() {
	//4. Спецификация цикла управления (CLI в main)
	//Сценарий работы Консольного меню:
	// Пункт 1

	tovar := map[string]int{
		"яблоки":        200, //quantity/qty - количество
		"профлисты":     200,
		"коробка":       300,
		"ред-булл":      150,
		"комплектующие": 300,
	}

	flotRobotov := []*Robot{{
		ID: 1, Model: "DRN-1", Type: Drone, Battery: 100, IsAvailable: true},
		{ID: 2, Model: "WLH-1", Type: Wheeled, Battery: 100, IsAvailable: true},
		{ID: 3, Model: "HVLift", Type: HeavyLifter, Battery: 100, IsAvailable: true},
		{ID: 4, Model: "DRN-2", Type: Drone, Battery: 70, IsAvailable: true},
	}

	//BufferCapacity = 5 буфер заказов, временное хранилище товаров перед отправкой
	orderBuffer := [BufferCapacity]Product{} //массив на 5 товаро по коснстанте

	orderCount := 0 // счетчик товаров в буфЕРЕ
	productID := 1  // счетчик для айди нов. товара. Увел. на 1

	for {
		fmt.Println(" МЕНЮ")
		fmt.Println("1: добавить товар")
		fmt.Println("2: склад и роботы")
		fmt.Println("3: дотставка")
		fmt.Println("4: зарядка роботов")
		fmt.Println("0: выход")

		var menu int
		fmt.Scan(&menu)

		if menu == 0 {
			fmt.Println("Завершение работы")
			break

		}

		switch menu {
		case 1: //чекин буфера
			for {
				if orderCount >= BufferCapacity {
					fmt.Println("Буфер заказов полон. Выполните доставку.")
					continue
				}

				fmt.Print("Введите товар:")
				var name string
				fmt.Scan(&name)

				if name == "стоп" || name == "0" {
					fmt.Println("Возврат в главное меню.")
					break // Выходим из внутреннего цикла
				}

				name = strings.ToLower(name)

				// чекин склада
				qty, ok := tovar[name]
				if !ok || qty == 0 {
					fmt.Printf("Товар %s отсутствует на складе\n", name)
					continue
				}

				tovar[name]-- //умен. колл товара на складе

				//генерирует продукт и записывает в буфр
				orderBuffer[orderCount] = Product{
					ID:   productID,
					Name: name,
				}

				productID++
				orderCount++

				fmt.Printf("Товар %s добавлен в заказ. В буфере: %d/%d\n", name, orderCount, BufferCapacity)
			}
		case 2:
			fmt.Println("Склад")
			for name, qty := range tovar {
				fmt.Printf("%-15s: %d шт\n", name, qty) // "-" выравнивает по левому краю. 15задает мин ширну в 15 симв
			}
			fmt.Println("Роботы ")
			for _, r := range flotRobotov {
				status := "занят"
				if r.IsAvailable {
					status = "свободен"
				}
				fmt.Printf(" %-10s %-20s Заряд: %.1f %s\n", r.Model, r.Type, r.Battery, status)
			}

			fmt.Println("Буфер заказов")
			if orderCount == 0 {
				fmt.Println("Буфер пуст")
			} else {
				for _, p := range orderBuffer[:orderCount] {
					fmt.Printf("[ID:%d] %s\n", p.ID, p.Name)
				}
			}

		case 3:
			if orderCount == 0 {
				fmt.Println("заказов нет")
				continue
			}

			reportCh := make(chan string, len(flotRobotov))
			activeJobs := 0

			// цикл по заказам в буфере
			for i := 0; i < orderCount; i++ {
				product := orderBuffer[i]

				//поиск робота
				var fRobot *Robot
				for _, r := range flotRobotov {
					if r.IsAvailable && r.Battery > EmergencyThreshold && r.CanCarryMrRobot(&product) == nil {
						fRobot = r
						break
					}
				}

				if fRobot == nil {
					fmt.Printf("Нет доступного робота %s \n", product.Name)
					continue
				}

				fRobot.IsAvailable = false
				r := fRobot // доступный робот
				p := product

				go func() {
					DeliveryTask(r, &p, 20, reportCh)
					r.IsAvailable = true
					if r.Battery < EmergencyThreshold {
						Recharge(r)
						fmt.Printf("%s перезарядка\n", r.Model)
					}
				}()
				activeJobs++
			}

			// сбор результатов. подсчет горутин
			for i := 0; i < activeJobs; i++ {
				msg := <-reportCh
				fmt.Println(msg)
			}

			orderCount = 0
			fmt.Println("доставки завершены")

		case 4:
			for _, r := range flotRobotov {
				Recharge(r)
				r.IsAvailable = true
			}
			fmt.Println("Все роботы заряжены")

		default:
			fmt.Println("Введие число из меню")
		}

		// err := HeavyLifter_1.Move(200)
		// if err != nil {
		// 	fmt.Println("Ошибка:", err)
		// }
		// fmt.Println("")
		// var wg sync.WaitGroup

		// wg.Wait()
		fmt.Println("Все роботы завершили работу")
	}

}
