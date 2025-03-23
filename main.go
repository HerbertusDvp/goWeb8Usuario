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

	// Recursos Codigo QR
	mux.HandleFunc("/recursos/qr", ruta.RecursosQR)

	// Recursos Email
	mux.HandleFunc("/recursos/email", ruta.RecursosEmail)

	// cliente http
	mux.HandleFunc("/clientehttp", ruta.ClienteHttp)
	mux.HandleFunc("/clientehttp/crear", ruta.ClienteHttpCrear)
	mux.HandleFunc("/clientehttp/crear-post", ruta.ClienteHttpCrearPost).Methods("POST")

	mux.HandleFunc("/clientehttp/editar/{id:.*}", ruta.ClienteHttpEditar)
	mux.HandleFunc("/clientehttp/editar-post/{id:.*}", ruta.ClienteHttpEditarPost)
	mux.HandleFunc("/clientehttp/eliminar/{id:.*}", ruta.ClienteHttpEliminar)

	// Mysql
	mux.HandleFunc("/mysql", ruta.MysqlListar)
	mux.HandleFunc("/mysql/crear", ruta.MysqlCrear).Methods("POST")
	mux.HandleFunc("/mysql/crearPost", ruta.MysqlCrearRecept)

	mux.HandleFunc("/mysql/editar/{id:.*}", ruta.MysqlEditar)
	mux.HandleFunc("/mysql/editarPost/{id:.*}", ruta.MysqlEditarRecept).Methods("POST")

	mux.HandleFunc("/mysql/eliminar/{id:.*}", ruta.MysqlEditar)

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
