package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterAdminRouter(AdminRouter *mux.Router) {
	AdminRouter.HandleFunc("", controllers.RenderAdminHome).Methods("GET")
	AdminRouter.HandleFunc("/users", controllers.RenderAdminUsers).Methods("GET")
	AdminRouter.HandleFunc("/orders", controllers.RenderAdminOrders).Methods("GET")
	// AdminRouter.HandleFunc("/inventory", controllers.RenderAdminHome).Methods("GET")
	AdminRouter.HandleFunc("/chef", controllers.RenderChef).Methods("GET")

}
