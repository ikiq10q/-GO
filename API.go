package main

import (
	"fmt"
	// "io"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func HandleUserName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func HandleSum(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	fmt.Fprintf(w, "The sum of %d and %d is: %d", a, b, a+b)
}

func HandleAdder(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(r.PathValue("id"))
	fmt.Fprintf(w, "User ID: %d", userId+10)
}

func main3() {
	// resp, err := http.Get("https://dogapi.dog/api/v2/breeds")
	// defer resp.Body.Close()
	// if err != nil {
	// 	fmt.Println("Error fetching data:", err)
	// 	return
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println(string(body))
	http.HandleFunc("/MyEnd", handler)
	http.HandleFunc("/UserName", HandleUserName)
	http.HandleFunc("/Sum", HandleSum)
	http.HandleFunc("/Adder/{id}", HandleAdder)
	http.ListenAndServe(":8080", nil)

}
