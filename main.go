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

	//Manejador para formularios
	mux.HandleFunc("/formulario", ruta.Formulario)
	mux.HandleFunc("/formulariop", ruta.Formulariop).Methods("POST")

	//Formulario para subir archivos
	mux.HandleFunc("/formulario/file", ruta.FormularioFile)
	mux.HandleFunc("/formulario/fileup", ruta.FormularioFileUp)

	//Recursos PDF
	mux.HandleFunc("/recursos", ruta.Recursos)
	mux.HandleFunc("/recursos/pdf", ruta.RecursosPdf)
	mux.HandleFunc("/recursos/generaPDF", ruta.RecursosGeneraPDF2)

	// recursos Excel
	mux.HandleFunc("/recursos/excel", ruta.RecursosExcel)
	mux.HandleFunc("/recursos/generaExcel", ruta.RecursosGeneraExcel)

	//Para recursos estaicos
	s := http.StripPrefix("/web/static/", http.FileServer(http.Dir("./web/static/")))
	mux.PathPrefix("/web/static/").Handler(s)

	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(ruta.Pagina404).GetHandler()

	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
