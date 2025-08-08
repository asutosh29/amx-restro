package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRouter(UserRouter *mux.Router) {
	UserRouter.HandleFunc("/orders", controllers.RenderUserOrder).Methods("GET")

}
