package api

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/api/routers"
	"github.com/asutosh29/amx-restro/pkg/middlewares"
	"github.com/asutosh29/amx-restro/pkg/utils/config"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	//Adding static files
	r.PathPrefix("/pkg/static/").Handler(http.StripPrefix("/pkg/static/", http.FileServer(http.Dir("./pkg/static/"))))
	// r.StrictSlash(true)
	r.Use(middlewares.LogRequests)
	r.Use(middlewares.LoggedIn)

	// Static Routes
	StaticRouter := r.PathPrefix("/").Subrouter()
	routers.RegisterStaticRouter(StaticRouter)

	// Menu router
	MenuRouter := r.PathPrefix("/menu").Subrouter()
	routers.RegisterMenuRouter(MenuRouter)

	// Order router
	OrderRouter := r.PathPrefix("/order").Subrouter()
	routers.RegisterOrderRouter(OrderRouter)

	// Admin router
	AdminRouter := r.PathPrefix("/admin").Subrouter()
	AdminRouter.Use(middlewares.AdminAccessOnly)
	routers.RegisterAdminRouter(AdminRouter)

	// Admin router
	UserRouter := r.PathPrefix("/user").Subrouter()
	routers.RegisterUserRouter(UserRouter)

	// Auth Routes
	AuthRouter := r.PathPrefix("/").Subrouter()
	AuthRouter.Use(middlewares.NewUser)
	routers.RegisterAuthRouter(AuthRouter)

	PORT := config.PORT
	fmt.Printf("Server starting running on port: %v\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), r)
}
