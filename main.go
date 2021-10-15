package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

	"database/sql"
	_ "github.com/lib/pq"
)



type ListItem struct {
	OrderName string `json:"Order name"`
	CustomerCompany string `json:"Customer Company"`
	CustomerName string `json:"Customer name"`
	OrderDate time.Time `json:"Order date"`
	DeliveredAmount float64 `json:"Delivered Amount"`
	TotalAmount float64 `json:"Total Amount"`
}

type CustomerCompany struct {
	companyId int32 `json:"Company ID"`
	CompanyName string `json:"Company Name"`
}

type Customer struct {
	UserId string `json:Customer Id`
	Name string `json:"Customer Name"`
	companyId int32 `json:"Company ID"`
}

type Delivery struct {
	Id int32 `json:"id"`
	orderItemId int32 `json:"Order Item ID"`
	DeliveredQuantity int32 `json:"Delivered Quantity"`
}

type OrderItem struct {
	id int32 `json:"ID"`
	OrderId int32 `json:"Order ID"`
	unitPrice float64 `json:"Unit Price"`
	Quantity int32 `json:"Quantity"`
	Product int32 `json:"Product"`
}

type Order struct {
	Id int32 `json:"id"`
	CreatedAt time.Time `json:"Created at"`
	OrderName string `json:"Order Name"`
	customerId string `json:"Customer ID"`
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "customerorders"
	password = "hBx3uYyRw4"
)

func allCompanies() []CustomerCompany {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
    	log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
    	log.Fatal(err)
	}

	companiesCol := client.Database("CustomerOrders").Collection("customer_companies")

	companiesCur, err := companiesCol.Find(context.TODO(), bson.M{})
	if err != nil {
    	log.Fatal(err)
	}

	
	// Close the cursor once finished
	defer companiesCur.Close(context.TODO())

	var companies []CustomerCompany

	for companiesCur.Next(context.TODO()) {
		var companyItem bson.M
		if err = companiesCur.Decode(&companyItem); err != nil {
			log.Fatal(err)
		}

		var company CustomerCompany = CustomerCompany{companyId: companyItem["company_id"].(int32), CompanyName: companyItem["company_name"].(string)}

		companies = append(companies, company)
		
	}
    
	return companies
}

func allCustomers() []Customer {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
    	log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
    	log.Fatal(err)
	}

	customersCol := client.Database("CustomerOrders").Collection("customers")

	customersCur, err := customersCol.Find(context.TODO(), bson.M{})
	if err != nil {
    	log.Fatal(err)
	}

	
	// Close the cursor once finished
	defer customersCur.Close(context.TODO())

	var customers []Customer

	for customersCur.Next(context.TODO()) {
		var customerItem bson.M
		if err = customersCur.Decode(&customerItem); err != nil {
			log.Fatal(err)
		}

		var customer Customer = Customer{UserId: customerItem["user_id"].(string), Name: customerItem["name"].(string), companyId: customerItem["company_id"].(int32)}

		customers = append(customers, customer)
		
	}
    
	return customers
}


func allOrders() ([]Delivery, []OrderItem, []Order) {
		
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
    	log.Fatal(err)
	}

	// retriving deliveries
	rows, err := db.Query("SELECT * FROM public.deliveries")
	if err != nil {
    	log.Fatal(err)
	}

	defer rows.Close()

	var deliveries []Delivery

	for rows.Next() {
		var delivery Delivery
		rows.Scan(&delivery.Id, &delivery.orderItemId, &delivery.DeliveredQuantity)

		deliveries = append(deliveries, delivery)
	}

	// retriving order items
	rows, err = db.Query("SELECT * FROM public.order_items")
	if err != nil {
    	log.Fatal(err)
	}

	defer rows.Close()

	var orderItems []OrderItem

	for rows.Next() {
		var orderItem OrderItem
		rows.Scan(&orderItem.id, &orderItem.OrderId, &orderItem.unitPrice, &orderItem.Quantity, &orderItem.Product)

		orderItems = append(orderItems, orderItem)
	}

	// retriving orders
	rows, err = db.Query("SELECT * FROM public.orders")
	if err != nil {
    	log.Fatal(err)
	}

	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		rows.Scan(&order.Id, &order.CreatedAt, &order.OrderName, &order.customerId)

		orders = append(orders, order)
	}

	return deliveries, orderItems, orders
}

func displayCompanies(w http.ResponseWriter, r *http.Request) {
	companies := allCompanies()
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(companies)

}

func displayCustomers(w http.ResponseWriter, r *http.Request) {
	customers := allCustomers()
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(customers)

}

func displayOrders(w http.ResponseWriter, r *http.Request) {
	deliveries, orderItems, orders := allOrders()
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(deliveries)
	enc.Encode(orderItems)
	enc.Encode(orders)
	
}

func displayOrderList(w http.ResponseWriter, r *http.Request) {
	companies := allCompanies()
	customers := allCustomers()
	deliveries, orderItems, orders := allOrders()

	fmt.Println(deliveries)

	var orderList []ListItem

	for _, value := range orderItems {
		var index int
		for i := range orders {
			if orders[i].Id == value.OrderId {
				index = i
				break
			}
		}
		orderName := orders[index].OrderName
		orderDate := orders[index].CreatedAt
		customerId := orders[index].customerId

		for i := range customers {
			if customers[i].UserId == customerId {
				index = i
				break
			}
		}

		customerName := customers[index].Name
		companyId := customers[index].companyId

		for i := range companies {
			if companies[i].companyId == companyId {
				index = i
				break
			}
		}

		companyName := companies[index].CompanyName
		quantity := float64(value.Quantity)*value.unitPrice

		listItem := ListItem{OrderName: orderName, CustomerCompany: companyName, CustomerName: customerName, OrderDate: orderDate, DeliveredAmount: quantity, TotalAmount: quantity}

		orderList = append(orderList, listItem)
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(orderList)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/companies", displayCompanies)
	myRouter.HandleFunc("/customers", displayCustomers)
	myRouter.HandleFunc("/orders", displayOrders)
	myRouter.HandleFunc("/orderList", displayOrderList).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()	
}