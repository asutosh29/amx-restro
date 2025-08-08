package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterOrderRouter(OrderRouter *mux.Router) {
	OrderRouter.HandleFunc("", controllers.HandleGetOrder).Methods("GET")
	OrderRouter.HandleFunc("", controllers.HandlePostOrder).Methods("POST")

	OrderRouter.HandleFunc("/placed/{id}", controllers.HandleOrderPlaced).Methods("PATCH")
	OrderRouter.HandleFunc("/cooking/{id}", controllers.HandleOrderCooking).Methods("PATCH")
	OrderRouter.HandleFunc("/served/{id}", controllers.HandleOrderServed).Methods("PATCH")
	OrderRouter.HandleFunc("/bill/{id}", controllers.HandleOrderBilled).Methods("PATCH")
	OrderRouter.HandleFunc("/paid/{id}", controllers.HandleOrderPaid).Methods("PATCH")

}
