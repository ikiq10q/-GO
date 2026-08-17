package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderItem struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var orderItems map[string]OrderItem
var orders map[string]*Order

type Order struct {
	ID           string      `json:"id"`
	CustomerName string      `json:"customer_name"`
	Status       string      `json:"status"`
	Items        []OrderItem `json:"items"`
	TotalPrice   float64     `json:"total_price"`
}

// 1 эндпоинт:
// Получение списка всех заказов. Возвращает список всех существующих заказов.
func GetOrders(c *gin.Context) {
	resl := make([]*Order, 0, len(orders))
	for _, order := range orders {
		resl = append(resl, order)
	}
	c.JSON(http.StatusOK, resl)
}

// 2 эндпоинт:
// Получение заказа по URL-параметру
func OrderID(c *gin.Context) {
	id := c.Param("id")

	//проверка на наличии заказа
	order, ok := orders[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// 3 эндпоинт:
// Поиск заказов по Query-параметру
func SearOrders(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр status обязателен"})

		return
	}

	//массив для найденых заказов
	resultQuery := make([]*Order, 0)

	for _, order := range orders {
		if order.Status == status {
			resultQuery = append(resultQuery, order)
		}
	}
	c.JSON(http.StatusOK, resultQuery)

}

// 4 эндпоинт:
// Создание нового заказа (JSON в теле запроса)
func CreateOrder(c *gin.Context) {
	var newOrder Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неправельный формат JSON",
		})
		return
	}
}

// 5 Эндпоинт: и валидация
// Изменение статуса через динамические параметры

func OrderStatusCooked(c *gin.Context) {
	id := c.Param("id")
	newStatus := c.Param("new_status")

	//проверка на наличии заказа
	order, ok := orders[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	//статсус заказа
	validStatuses := map[string]bool{
		"cooking":   true,
		"delivered": true,
		"cancelled": true,
	}
	// проверка на валидацию
	if !validStatuses[newStatus] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недопустимый статус"})
		return
	}
	if order.Status == "delivered" || order.Status == "cancelled" {
		c.JSON(http.StatusConflict, gin.H{"error": "Нельзя изменить статус уже завершенного заказа"})
		return
	}

	order.Status = newStatus
	c.JSON(http.StatusOK, order)

}

func main() {
	orders = map[string]*Order{
		"ORD-1": {
			ID:           "опв-1",
			CustomerName: "Матвей",
			Status:       "created",
			Items: []OrderItem{{
				Title:    "Кукси",
				Price:    1690,
				Quantity: 2,
			}, {
				Title:    "Лагман",
				Price:    1690,
				Quantity: 1,
			}},
			TotalPrice: 1690*2 + 1690*1,
		},
		"ORD-2": {
			ID:           "опв-2",
			CustomerName: "Айбат",
			Status:       "created",
			Items: []OrderItem{{
				Title:    "Ташкенский",
				Price:    890,
				Quantity: 1,
			}, {
				Title:    "Хинкали",
				Price:    490,
				Quantity: 5,
			}},
			TotalPrice: 890*1 + 490*5,
		},
	}

	r := gin.Default()
	r.GET("/orders", GetOrders)                                  //1
	r.GET("/orders/:id", OrderID)                                //2
	r.GET("/orders/search", SearOrders)                          //3
	r.GET("/orders/newOrder", CreateOrder)                       //4
	r.PATCH("/orders/:id/status/:new_status", OrderStatusCooked) //5

	fmt.Println("Сервер запущен на 8080 порту")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
