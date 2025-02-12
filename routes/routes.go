package routes

import (
	"github.com/gorilla/mux"
	"main.go/controllers"
	"main.go/middlewares"
)

func GetRouter() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/ws", middlewares.DBcontext(controllers.HandleWS))
	Router.HandleFunc("/", middlewares.DBcontext(controllers.Home)).Methods("GET")
	Router.HandleFunc("/upload", middlewares.DBcontext(controllers.UploadHandler)).Methods("POST")
	Router.HandleFunc("/createOrder", middlewares.DBcontext(controllers.CreateOrder)).Methods("POST")
	// Router.HandleFunc("/fetchOrders", middlewares.DBcontext(controllers.FetchOrders)).Methods("POST")

	return Router
}
