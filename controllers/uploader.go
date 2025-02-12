package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"main.go/services/msgq"
)

type payload struct {
	File []byte `json:"file"`
	Id   string `json:"id"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintf(w, "Error reading file: %v", err)
		return
	}

	fileData, _ := io.ReadAll(file)

	data := &payload{
		File: fileData,
		Id:   r.URL.Query().Get("key"),
	}
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	msgq.MsgQueue.Produce(jsonData)

	// defer file.Close()

	fmt.Println("File Recieved!")
}
