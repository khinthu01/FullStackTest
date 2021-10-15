package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"os"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

	"database/sql"
	_ "github.com/lib/pq"
)



// type Order struct {
// 	OrderName string `json:"Order name"`
// 	CustomerCompany string `json:"Customer Company"`
// 	CustomerName string `json:"Customer name"`
// 	OrderDate time.Time `json:"Order date"`
// 	DeliveredAmount float64 `json:"Delivered Amount"`
// 	TotalAmount float64 `json:"Total Amount"`
// }

type CustomerCompany struct {
	companyID int32 `json:"Company ID"`
	CompanyName string `json:"Company Name"`
}

type Customer struct {
	Name string `json:"Customer Name"`
	companyID int32 `json:"Company ID"`
}

type Order struct {
	id int32 `json:"id"`
	CreatedAt time.Time `json:"Created at"`
	OrderName string `json:"Order Name"`
	customerID string `json:"Customer ID"`
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "customerorders"
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

		var company CustomerCompany = CustomerCompany{companyID: companyItem["company_id"].(int32), CompanyName: companyItem["company_name"].(string)}

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

		var customer Customer = Customer{Name: customerItem["name"].(string), companyID: customerItem["company_id"].(int32)}

		customers = append(customers, customer)
		
	}
    
	return customers
}


func allOrders() []Order {
	password := os.Getenv("PASSWORD")

	fmt.Println(password)
		
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
    	log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM public.orders")
	if err != nil {
    	log.Fatal(err)
	}

	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		rows.Scan(&order.id, &order.CreatedAt, &order.OrderName, &order.customerID)

		orders = append(orders, order)
	}

	return orders
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
	orders := allOrders()
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(orders)
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
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {

	

	handleRequests()

	
}