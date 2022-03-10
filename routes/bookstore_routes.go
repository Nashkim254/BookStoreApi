package routes

import (
	"github.com/gorrilla/mux"
	"github.com/programming6131/BookStoreApi/controller"
)

var RegisterBookRoutes = func(router *mux.Router) {

	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book", controller.GetteBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
