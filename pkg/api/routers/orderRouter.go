package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterOrderRouter(OrderRouter *mux.Router) {
	OrderRouter.HandleFunc("", controllers.HandleGetOrder).Methods("GET")
	OrderRouter.HandleFunc("", controllers.HandlePostOrder).Methods("POST")

}
