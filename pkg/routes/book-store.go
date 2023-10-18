package routes

import (
	"github.com/gorilla/mux",
	"github.com/saeedmzr"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/",controllers)
}


