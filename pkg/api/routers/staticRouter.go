package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterStaticRouter(StaticRouter *mux.Router) {
	StaticRouter.HandleFunc("/", controllers.RenderHome).Methods("GET")
	StaticRouter.HandleFunc("/home", controllers.RenderHome).Methods("GET")
	StaticRouter.HandleFunc("/payment", controllers.RenderPayment).Methods("GET")
	StaticRouter.HandleFunc("/profile", controllers.RenderProfile).Methods("GET")

	StaticRouter.HandleFunc("/logout", controllers.HandleLogOut).Methods("GET")

}
