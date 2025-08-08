package api

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/api/routers"
	"github.com/asutosh29/amx-restro/pkg/middlewares"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	//Adding static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// r.StrictSlash(true)
	r.Use(middlewares.LogRequests)
	r.Use(middlewares.LoggedIn)

	// Static Routes
	StaticRouter := r.PathPrefix("/").Subrouter()
	routers.RegisterStaticRouter(StaticRouter)

	// Menu route
	MenuRouter := r.PathPrefix("/menu").Subrouter()
	routers.RegisterMenuRouter(MenuRouter)

	// Menu route
	OrderRouter := r.PathPrefix("/order").Subrouter()
	routers.RegisterOrderRouter(OrderRouter)

	// Auth Routes
	AuthRouter := r.PathPrefix("/").Subrouter()
	AuthRouter.Use(middlewares.NewUser)
	routers.RegisterAuthRouter(AuthRouter)

	fmt.Println("Server starting running on port: 8000")
	http.ListenAndServe(":8000", r)
}
