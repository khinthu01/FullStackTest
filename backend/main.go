package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)



type Order struct {
	OrderName string `json:"Order name"`
	CustomerCompany string `json:"Customer Company"`
	CustomerName string `json:"Customer name"`
	OrderDate time.Time `json:"Order date"`
	DeliveredAmount float64 `json:"Delivered Amount"`
	TotalAmount float64 `json:"Total Amount"`
}

type Orders []Orders

func connectToDB() {
	// Set client options
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

	fmt.Println("Connected to MongoDB!")
}

func allOrders(w http.ResponseWriter, r *http.Request) {
	
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
	connectToDB()
}