package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterStaticRouter(StaticRouter *mux.Router) {
	StaticRouter.HandleFunc("/", controllers.RenderHome).Methods("GET")
	StaticRouter.HandleFunc("/home", controllers.RenderHome).Methods("GET")
	StaticRouter.HandleFunc("/login", controllers.RenderLogin).Methods("GET")
	StaticRouter.HandleFunc("/login", controllers.HandleLoginUser).Methods("POST")
	StaticRouter.HandleFunc("/register", controllers.RenderRegister).Methods("GET")
	StaticRouter.HandleFunc("/register", controllers.HandleRegisterUser).Methods("POST")

}
