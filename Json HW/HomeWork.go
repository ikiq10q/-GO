package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// zadanie 1
type Movie struct {
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

func MovieHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	movie := Movie{
		Title:  "Karmalogic",
		Year:   2017,
		Rating: 8.2,
	}

	jsonData, err := json.Marshal(movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ошибка маршала: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(jsonData))
}

// zadanie 2
type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func parseProduct(data []byte) (*Product, error) {

	var product Product

	err := json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &product, nil

}

// zadanie 3
type UserRegistration struct {
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func registerUSer(regData []byte) (*UserRegistration, error) {
	var parseRegs UserRegistration

	err := json.Unmarshal(regData, &parseRegs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if parseRegs.Email == "" {
		return nil, errors.New("ЗАполните поле Email")
	}
	if parseRegs.Age < 18 {
		return nil, errors.New("Вам меньше 18")
	}
	return &parseRegs, nil
}

// zadanie 4
type WeatherReport struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	var report WeatherReport

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&report); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Некорректный формат JSON")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Город: %s, Температура: %.1f°C", report.City, report.Temperature)
}

// zadanie 5
type ServerStatus struct {
	Status        string `json:"status"`
	UptimeSeconds int    `json:"uptime_seconds"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	var nowStatus ServerStatus

	fmt.Print("Введите статус:")
	fmt.Scan(&nowStatus.Status)
	nowStatus.Status = strings.ToLower(nowStatus.Status)

	fmt.Print("Введите время работы: ")
	fmt.Scan(&nowStatus.UptimeSeconds)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(nowStatus)
	if err != nil {
		http.Error(w, "Ошибка отправки JSON", http.StatusInternalServerError)
		return
	}
}
func main() {
	//zadanie 1
	http.HandleFunc("/movie", MovieHandler)

	fmt.Println("Server is running on port 8080...")

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("Ошибка сервера:", err)
	// }

	//zadanie 2
	jsonData2 := []byte(`{"name":  "varmilo","price": 250}`)

	product, err := parseProduct(jsonData2)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Некорректный формат JSON")
		return
	}

	fmt.Fprintf(os.Stdout, "Товар: %s, Цена: %d\n", product.Name, product.Price)

	//zadanie 3
	// jsonData3 := []byte(`{"email": "AArdak@ikiq.com","age": 29}`)

	// parseRegs, err := registerUSer(jsonData3)
	// if err != nil {
	// 	fmt.Printf("Некорректный формат JSON", err)
	// 	return
	// }
	// fmt.Printf("Email: %s, Возраст: %d\n", parseRegs.Email, parseRegs.Age)

	// jsonData4 := []byte(`{"email": ","age": 2}`)

	// parseRegs, err = registerUSer(jsonData4)
	// if err != nil {
	// 	fmt.Printf("Некорректный формат JSON", err)
	// 	return
	// }
	testMAil := []struct {
		name string
		json []byte
	}{
		{
			name: "test1",
			json: []byte(`{"email": "AArdak@ikiq.com", "age": 29}`),
		},
		{
			name: "test2",
			json: []byte(`{"email": ","age": 2}`),
		},
	}
	for _, tM := range testMAil {
		fmt.Printf("Тест: %s \n", tM.name)

		parseRegs, err := registerUSer(tM.json)
		if err != nil {
			fmt.Printf("Некорректный формат JSON", err)
			continue
		}
		fmt.Printf("Email: %s, Возраст: %d\n", parseRegs.Email, parseRegs.Age)
	}

	//zadanie 4
	http.HandleFunc("/weather", weatherHandler)

	// err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}

	//zadanie 5
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Запуск сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
