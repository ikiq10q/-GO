package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	return
}

func main() {
	//zadanie 1
	http.HandleFunc("/movie", MovieHandler)

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}

	//zadanie 2
	jsonData := []byte{

		"name":  "varmilo",
		"price": 250,
	}

	product, err := parseProduct(jsonData)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Некорректный формат JSON")
		return
	}

	fmt.Fprintf(os.Stdout, "Товар: %s, Цена: %d\n", product.Name, product.Price)
}
