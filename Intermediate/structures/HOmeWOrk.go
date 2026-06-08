package main

import (
	"fmt"
	"math"
)

// Zadanie 1
type Person struct {
	Name string
	Age  int
}

// zadanie 2
type Book struct {
	Title, Author string
	Pager         int
}

// zadanie 3
type Car struct {
	Brand string
	Year  int
}

// zadanie 4
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Square() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//zadanie 5

type Student struct {
	Name  string
	Grade int
}

func (s Student) Tabel() string {
	return fmt.Sprintf("Студент %s получил %d баллов", s.Name, s.Grade)
}

// zadanie 6
type Circle struct {
	Radius float64
}

// zadanie 7
type Car7 struct {
	Brand string
	Speed int
}

func (с Car7) CheckSpeed() string {
	if с.Speed > 100 {
		fmt.Println("Слишком быстро")
		return "Слишком быстро"
	} else if с.Speed > 60 {
		fmt.Println("Нормальная скорость")
		return "Нормальная скорость"
	} else {
		fmt.Println("Медленно")
		return "Медленно"
	}
}

// zadanie 10
type Book_1 struct {
	Title  string
	Pages  int
	IsRead bool
}

func getUnreadLongBooks(books []Book_1) []Book_1 {
	var result []Book_1
	for _, v := range books {
		if v.IsRead == false && v.Pages > 300 {
			result = append(result, v)
		}
	}
	return result
}

// zadanie 11
type Company struct {
	CompanyName string
}
type Department struct {
	Name string
	Org  Company
}

func updateCompanyName(up *Department, upName string) {
	up.Org.CompanyName = upName

}

// zadanie 12
type Item struct {
	ID     int
	Name   string
	Weight float64
}

func findHeaviestItem(Items []Item) *Item {

	if len(Items) == 0 {
		return nil
	}
	result := 0
	for i, v := range Items {
		if v.Weight > Items[result].Weight {
			result = i
		}

	}
	return &Items[result]

}

// zadanie 13
type Account struct {
	Username string
	Age      int
	Status   string
}

func createAccount(Username string, Age int) Account {
	a := Account{ //  объект  в переданными данными
		Username: Username,
		Age:      Age,
	}
	if a.Age < 18 {
		a.Status = "restricted"
	} else if a.Age >= 18 {
		a.Status = "active"
	}
	return a
}

func main() {
	//Zadanie 1
	Persona1 := Person{
		Name: "Kamik",
		Age:  22,
	}
	fmt.Println(Persona1)

	//zadanie 2
	Book1 := Book{
		Title:  "Gragas_supp",
		Author: "Kamik",
		Pager:  99,
	}

	fmt.Println(Book1)

	//Zadanie 3

	// Car1 := Car{
	// 	Brand: "Mercedes",
	// 	Year:  2022,
	// }
	// var C_Car *Car = &Car1
	// fmt.Println(C_Car)

	var Car1 *Car = &Car{
		Brand: "Mercedes",
		Year:  2022,
	}
	fmt.Println(Car1)

	//zadanie 4
	Rectangle1 := Rectangle{
		Width:  5,
		Height: 6,
	}

	fmt.Println("Периметр:", Rectangle1.Perimeter())
	fmt.Println("Площадь:", Rectangle1.Square())

	fmt.Println("Периметр:", 2*(Rectangle1.Width+Rectangle1.Height))
	fmt.Println("Площадь:", Rectangle1.Width*Rectangle1.Height)

	Rectangle1.Height = 7
	fmt.Println("Периметр:", Rectangle1.Perimeter())
	fmt.Println("Площадь:", Rectangle1.Square())

	//zadanie 5
	student1 := Student{
		Name:  "Aibek",
		Grade: 95,
	}

	student2 := Student{
		Name:  "Alice",
		Grade: 88,
	}
	fmt.Printf("Студент %s получил %d баллов \n", student1.Name, student1.Grade)
	fmt.Printf("Студент %s получил %d баллов \n", student2.Name, student2.Grade)

	//через фунцию
	fmt.Println(student1.Tabel())
	fmt.Println(student2.Tabel())

	//zadanie 6
	Radius120 := Circle{
		Radius: 120,
	}
	fmt.Println("Периметр:", 2*math.Pi*Radius120.Radius)
	fmt.Println("Площадь:", math.Pi*Radius120.Radius*Radius120.Radius)

	Radius120.Radius = 4
	fmt.Println("Периметр:", 2*math.Pi*Radius120.Radius)
	fmt.Println("Площадь:", math.Pi*Radius120.Radius*Radius120.Radius)

	//zadanie 7
	balide1 := Car7{
		Brand: "Mercedes",
		Speed: 120,
	}
	balide2 := Car7{
		Brand: "BMW",
		Speed: 200,
	}
	balide3 := Car7{
		Brand: "Lada",
		Speed: 12,
	}
	balide1.CheckSpeed()
	balide2.CheckSpeed()
	balide3.CheckSpeed()

	fmt.Printf("Машина: %s, Скорость: %d км/ч -> Статус: %s\n", balide1.Brand, balide1.Speed, balide1.CheckSpeed())

	//zadanie 9
	// В main создайте анонимную структуру для конфигурации сервера.
	// Она должна содержать поля Host (строка), Port (int) и Environment (строка).
	// Сразу же инициализируйте её значениями ("localhost", 8080, "development") и выведите в консоль фразу:
	// "Сервер запущен на localhost:8080", используя данные из созданной структуры.

	//через функию
	//func (s struct { Host string; Port int; Environment string }) Address() string {
	//return fmt.Sprintf("%s:%d", s.Host, s.Port)

	Serv := struct {
		Host        string
		Port        int
		Environment string
	}{
		Host:        "localhost",
		Port:        8080,
		Environment: "development",
	}
	fmt.Printf("Сервер запущен на %s: %d\n", Serv.Host, Serv.Port)

	//zadanie 10
	// Создайте структуру Book с полями Title (строка), Pages (int) и IsRead (bool).
	// Напишите функцию getUnreadLongBooks, которая принимает слайс []Book и возвращает новый слайс,
	// содержащий только те книги, которые еще не прочитаны (IsRead == false) и в которых больше 300 страниц.
	// В main протестируйте функцию на слайсе из 4-5 разных книг.

	library := []Book_1{
		{"Мастер и Маргарита", 480, false},
		{"Karmalogic", 840, true},
		{"«Маленький принц", 98, false},
		{"Sapiens", 520, false},
	}
	fmt.Println(getUnreadLongBooks(library))

	//zadanie 11
	// Создайте структуру Company с полем CompanyName (строка).
	// Создайте структуру Department с полями Name (строка) и Org (вложенная структура Company).
	// Напишите функцию updateCompanyName, которая принимает указатель на структуру департамента *Department и строку с новым названием компании.
	// Функция должна изменить название компании внутри департамента.
	// В main проверьте, что имя компании действительно изменилось.
	Company1 := Department{
		Name: "Apple",
		Org: Company{
			CompanyName: "Nvidia",
		},
	}
	fmt.Println(Company1.Name)

	updateCompanyName(&Company1, "AMD")
	fmt.Println(Company1)

	//zadanie 12
	// Создайте структуру Item с полями ID (int), Name (строка) и Weight (float64).
	// Напишите функцию findHeaviestItem, которая принимает слайс []Item и возвращает указатель на самую тяжелую вещь (*Item).
	// Если слайс пустой, функция должна возвращать nil.
	// В main создайте слайс из 3 предметов, найдите самый тяжелый и через полученный указатель измените его имя на "Самый тяжелый предмет".
	// Выведите исходный слайс, чтобы убедиться, что имя изменилось.

	Items_3 := []Item{
		{1, "Dumbbell", 12},
		{2, "Plates", 8},
		{3, "Barbell", 20},
	}
	heaviest := findHeaviestItem(Items_3)
	heaviest.Name = "Самый тяжелый предмет: " + heaviest.Name
	fmt.Println(Items_3)

	//zadanie 13
	// Создайте структуру Account с полями Username (строка), Age (int) и Status (строка).
	// Напишите функцию createAccount, которая принимает username и age, а возвращает готовую структуру Account.
	// Внутри функции добавьте логику: если age меньше 18, то поле Status должно автоматически заполняться значением "restricted",
	// А если 18 и старше — "active". В main вызовите функцию дважды (для возраста 15 и 23 лет) и выведите результаты.

	user1 := createAccount("Ikiq", 17)
	user2 := createAccount("kam1kjke", 23)
	user3 := createAccount("Hashira", 28)

	fmt.Printf("Пользователь 1:%+v\n", user1)
	fmt.Printf("Пользователь 2:%+v\n", user2)
	fmt.Printf("Пользователь 3:%+v\n", user3)

}
