package routers

import (
	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterAuthRouter(AuthRouter *mux.Router) {
	AuthRouter.HandleFunc("/login", controllers.RenderLogin).Methods("GET")
	AuthRouter.HandleFunc("/login", controllers.HandleLoginUser).Methods("POST")
	AuthRouter.HandleFunc("/register", controllers.RenderRegister).Methods("GET")
	AuthRouter.HandleFunc("/register", controllers.HandleRegisterUser).Methods("POST")

}
