package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	CustomerId string `bson:"customerid" json:"customerid"`
	ProductId  string `bson:"productid" json:"productid"`
	Status     string `bson:"status" json:"status"`
}

type Response struct {
	Success bool `json:"success"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db_instance").(*mongo.Client)
	ctx := r.Context().Value("db_ctx").(context.Context)

	q := r.URL.Query().Get("key")
	conn := Manager.Client[q]

	var data Order
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(data)
	newBody := bytes.NewReader(body)

	res, err := http.Post("http://localhost:1800/verify", "application/json", newBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var response Response
	if json.NewDecoder(res.Body).Decode(&response); !response.Success {
		conn.WriteMessage(websocket.TextMessage, []byte("Order Failed!"))
		fmt.Println("Order failed!")
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Order Initiated!"))
	orders := db.Database("test_db").Collection("orders")

	_, err = orders.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Order Created!"))
	fmt.Println("Order created")
}
