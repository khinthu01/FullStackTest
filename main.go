package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/lib/pq"
)


type Order struct {
	OrderId int32 `json:"OrderId"`
	OrderName string `json:"OrderName"`
	CustomerCompany string `json:"CustomerCompany"`
	CustomerName string `json:"CustomerName"`
	OrderDate time.Time `json:"OrderDate"`
	DeliveredAmount float64 `json:"DeliveredAmount"`
	TotalAmount float64 `json:"TotalAmount"`
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "customerorders"
	password = "hBx3uYyRw4"
)


func main() {
	app := fiber.New()
	app.Use(cors.New())

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
	    log.Fatal(err)
	}

	app.Get("/orders", func(c *fiber.Ctx) error {
		var orders []Order

		sql := "SELECT * from public.orderlist"

		if s := c.Query("s"); s != "" {
			sql = fmt.Sprintf("%s WHERE order_name LIKE '%%%s%%'", sql, s)
		}

		rows, err := db.Query(sql)

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var order Order
			rows.Scan(&order.OrderId, &order.OrderName, &order.CustomerCompany, &order.CustomerName, &order.OrderDate, &order.DeliveredAmount, &order.TotalAmount)
	
			orders = append(orders, order)
		}

		return c.JSON(orders)

	})

	app.Listen(":8000")
}




