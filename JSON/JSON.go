package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	IsDone      bool   `json:"is_done"`
}

var tasks []Task

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Не удалось закодировать задачи", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	task.ID = len(tasks) + 1
	task.CreatedAt = time.Now().Format(time.RFC3339)

	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	otvet := map[string]string{
		"message": "Вы добавили задачу",
	}

	data, err := json.Marshal(otvet)
	if err != nil {
		http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasksHandler(w, r)
	case http.MethodPost:
		createTaskHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}

// func main() {
// 	jsonData := []byte(`{

// 	"title":   "Karm",
//     "author":  "ads",
//     "pages":   323,
//     "reading": false
// 	}`)

// 	var book Book

// err := json.Unmarshal(jsonData, &Task)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(book.Title)
// 	fmt.Println(book.Author)
// 	fmt.Println(book.Pages)
// 	fmt.Println(book.Reading)

// }
