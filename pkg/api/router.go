package api

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/controllers"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	//Adding static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", controllers.RenderHome).Methods("GET")
	fmt.Println("Server starting running on port: 8000")
	http.ListenAndServe(":8000", r)
}
