package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterMenuRouter(MenuRouter *mux.Router) {
	MenuRouter.HandleFunc("/", controllers.RenderMenu).Methods("GET")

}
