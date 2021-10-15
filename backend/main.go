package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)



// type Order struct {
// 	OrderName string `json:"Order name"`
// 	CustomerCompany string `json:"Customer Company"`
// 	CustomerName string `json:"Customer name"`
// 	OrderDate time.Time `json:"Order date"`
// 	DeliveredAmount float64 `json:"Delivered Amount"`
// 	TotalAmount float64 `json:"Total Amount"`
// }

type Order struct {
	id int32 `json:"id"`
	createdAt string `json:"created_at"`
	orderName string `json:"order_name"`
	customerID string `json:"customer_id"`
}

func allOrders(w http.ResponseWriter, r *http.Request) {
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

	// fmt.Println("Connected to MongoDB!")

	collection := client.Database("CustomerOrders").Collection("orders")

	

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
    	log.Fatal(err)
	}

	
	// Close the cursor once finished
	defer cur.Close(context.TODO())

	var orderRes []bson.M
	var orders []Order

	for cur.Next(context.TODO()) {
		var orderItem bson.M
		if err = cur.Decode(&orderItem); err != nil {
			log.Fatal(err)
		}
		orderRes = append(orderRes, orderItem)

		var order Order = Order{id: orderItem["id"].(int32), createdAt: orderItem["created_at"].(string), orderName: orderItem["order_name"].(string), customerID: orderItem["customer_id"].(string)}

		orders = append(orders, order)
		
	}

	fmt.Println(orders)
	json.NewEncoder(w).Encode(orders)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/orders", allOrders)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()	
}