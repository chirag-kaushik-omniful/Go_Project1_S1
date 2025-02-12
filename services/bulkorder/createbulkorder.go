package bulkorder

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"main.go/controllers"
	"main.go/services/dbconn"
)

type payload struct {
	File []byte `json:"file"`
	Id   string `json:"id"`
}

type Response struct {
	Success bool `json:"success"`
}

type Order struct {
	CustomerId string `bson:"customerid" json:"customerid"`
	ProductId  string `bson:"productid" json:"productid"`
	Status     string `bson:"status" json:"status"`
}

func CreateBulkOrder(msg []byte) {
	db := dbconn.DB_Instance
	ctx := dbconn.DB_Ctx

	var data payload
	json.Unmarshal(msg, &data)

	conn := controllers.Manager.Client[data.Id]
	file := bytes.NewReader(data.File)

	// Read the CSV file
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Error processing CSV!"))
		fmt.Printf("Error processing CSV: %v", err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("CSV File Uploaded Successfully!"))
	var failed int = 0
	// Display the rows
	fmt.Println("Uploaded CSV Data: ")
	for i := 1; i < len(rows); i++ {
		data := &Order{rows[i][1], rows[i][2], "pending"}
		body, _ := json.Marshal(data)
		newBody := bytes.NewReader(body)
		res, err := http.Post("http://localhost:1800/verify", "application/json", newBody)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Failed to verify order info!"))
			return
		}

		var response Response
		if json.NewDecoder(res.Body).Decode(&response); !response.Success {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Order %s Failed!", rows[i][0])))
			fmt.Println("Order failed!")
			failed++
			continue
		}

		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Order %s Initiated!", rows[i][0])))
		orders := db.Database("test_db").Collection("orders")

		_, err = orders.InsertOne(ctx, data)
		if err != nil {
			log.Fatal(err)
		}

		conn.WriteMessage(websocket.TextMessage, []byte("Order Created!"))
		fmt.Println("Order created")
	}
	conn.WriteMessage(websocket.TextMessage, []byte("Order Successfully Placed!"))
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Successful Placements: %d", len(rows)-failed-1)))
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("UnSuccessful Placements: %d", failed)))
}
