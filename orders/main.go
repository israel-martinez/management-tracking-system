package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type item struct {
	Description	string 	`json:"Description"`
	Price		float64 `json:"Price"`
	Quantity	int32 	`json:"Quantity"`
}

type order struct {
	ID          string 	`json:"ID"`
	Status      string 	`json:"Status"`
	Description string 	`json:"Description"`
	Address 	string 	`json:"Address"`
	Items		[]item 	`json:"Items"`
}

type allOrders []order

//recibido, preparando, en camino y entregado
var orders = allOrders{
	{
		ID:          "1",
		Status:      "Recibido",
		Description: "Compra promocional, Pizza Individual + Lata Bebida",
		Address: 	 "Avda. Principal Golang 8080",
		Items:	[]item {
					{
						Description: "Pizza Individual Con Peperonni",
						Price: 4900.0,
						Quantity: 1,
					},
			},
	},
	{
		ID:          "2",
		Status:      "Preparando",
		Description: "Fiesta Familiar Pizza + Botella bebida 1.5L",
		Address: 	 "Avda. Mongo Cassandra 8000",
		Items:	[]item {
					{
						Description: "Pizza Familiar Cuatro Estaciones",
						Price: 10900.0,
						Quantity: 2,
					},
					{
						Description: "Porción Aros de cebolla",
						Price: 1990.0,
						Quantity: 4,
					},
			},
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Microservice MS-A-01")
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder order
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please insert required register input for Order")
	}
	
	json.Unmarshal(reqBody, &newOrder)
	orders = append(orders, newOrder)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newOrder)
}

func getOneOrder(w http.ResponseWriter, r *http.Request) {
	OrderID := mux.Vars(r)["id"]

	for _, singleOrder := range orders {
		if singleOrder.ID == OrderID {
			json.NewEncoder(w).Encode(singleOrder)
		}
	}
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]
	var updatedOrder order

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please insert required register input for Order")
	}
	json.Unmarshal(reqBody, &updatedOrder)

	for i, singleOrder := range orders {
		if singleOrder.ID == orderID {
			singleOrder.Status = updatedOrder.Status
			singleOrder.Description = updatedOrder.Description
			singleOrder.Address = updatedOrder.Address
			singleOrder.Items = updatedOrder.Items
			orders = append(orders[:i], singleOrder)
			json.NewEncoder(w).Encode(singleOrder)
		}
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]

	for i, singleOrder := range orders {
		if singleOrder.ID == orderID {
			orders = append(orders[:i], orders[i+1:]...)
			fmt.Fprintf(w, "The order with ID %v has been deleted successfully", orderID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getAllOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", getOneOrder).Methods("GET")
	router.HandleFunc("/orders/{id}", updateOrder).Methods("PATCH")
	router.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}