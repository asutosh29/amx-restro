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
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./pkg/static/images/"))))
	// r.StrictSlash(true)
	r.Use(middlewares.LogRequests)
	// r.Use(middlewares.LoggedIn)

	// Auth Routes
	AuthRouter := r.PathPrefix("/").Subrouter()
	AuthRouter.Use(middlewares.NewUser)
	routers.RegisterAuthRouter(AuthRouter)

	// Static Routes
	StaticRouter := r.PathPrefix("/").Subrouter()
	StaticRouter.Use(middlewares.RestrictToLoggedIn)
	routers.RegisterStaticRouter(StaticRouter)

	// Menu router
	MenuRouter := r.PathPrefix("/menu").Subrouter()
	MenuRouter.Use(middlewares.RestrictToLoggedIn)
	routers.RegisterMenuRouter(MenuRouter)

	// Order router
	OrderRouter := r.PathPrefix("/order").Subrouter()
	OrderRouter.Use(middlewares.RestrictToLoggedIn)
	routers.RegisterOrderRouter(OrderRouter)

	// Admin router
	AdminRouter := r.PathPrefix("/admin").Subrouter()
	AdminRouter.Use(middlewares.RestrictToLoggedIn)
	AdminRouter.Use(middlewares.AdminAccessOnly)
	routers.RegisterAdminRouter(AdminRouter)

	// User router
	UserRouter := r.PathPrefix("/user").Subrouter()
	UserRouter.Use(middlewares.RestrictToLoggedIn)
	routers.RegisterUserRouter(UserRouter)

	PORT := config.PORT
	fmt.Printf("Server starting running on port: %v\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), r)
}
