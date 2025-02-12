package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/controllers"
	"main.go/routes"

	"main.go/services/bulkorder"
	"main.go/services/dbconn"
	"main.go/services/msgq"
)

func main() {

	queue, err := msgq.NewRabbitMQ("amqp://guest:guest@localhost:5672/", "test_queue")
	if err != nil {
		log.Fatal(err)
	}
	msgq.MsgQueue = queue
	defer msgq.MsgQueue.Close()

	message, _ := msgq.MsgQueue.Consume()

	go func() {
		fmt.Println("Output:")
		for d := range message {
			bulkorder.CreateBulkOrder(d.Body)

		}
	}()

	// msgq.MsgQueue.Produce([]byte("Hello World"))

	r := routes.GetRouter()
	controllers.SocketInit()

	dbconn.Connect("mongodb://localhost:27017")

	err1 := http.ListenAndServe(":1900", r)
	if err1 != nil {
		log.Fatal("ListenAndServe:", err1)
	}
}
