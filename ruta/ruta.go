package ruta

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Servicio1() {
	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(reponse, "Hola web con go con la ruta del servicio 1")
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/nosotros.html")

	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "Estas son las peticiones: \nid: "+vars["id"]+"\nnombre: "+vars["nombre"])
}

func ParametrosQS(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "-- Parametros ocn String Query -- ")
	fmt.Fprintln(response, request.URL)
	fmt.Fprintln(response, request.URL.RawQuery)
	fmt.Fprintln(response, request.URL.Query())
	fmt.Fprintln(response, request.URL.Query().Get("id"))
	fmt.Fprintln(response, request.URL.Query().Get("nombre"))
}
