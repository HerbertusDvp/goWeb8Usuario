package main

import (
	"goweb1/ruta"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	//ruta.Servicio1()

	mux := mux.NewRouter()
	mux.HandleFunc("/", ruta.Home)
	mux.HandleFunc("/nosotros", ruta.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{nombre:.*}", ruta.Parametros)
	mux.HandleFunc("/parametrosQS", ruta.ParametrosQS)
	mux.HandleFunc("/estructuras", ruta.Estructuras)

	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
