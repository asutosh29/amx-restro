package api

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/api/routers"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	//Adding static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	StaticRouter := r.PathPrefix("/").Subrouter()
	routers.RegisterStaticRouter(StaticRouter)

	fmt.Println("Server starting running on port: 8000")
	http.ListenAndServe(":8000", r)
}
