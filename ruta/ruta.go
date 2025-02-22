package ruta

import (
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

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/nosotros.html")

	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/parametros.html")
	vars := mux.Vars(request) // Obtiene los paramatros de la url
	data := map[string]string{
		"id":     vars["id"],
		"nombre": vars["nombre"],
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}

func ParametrosQS(response http.ResponseWriter, request *http.Request) {

	template, err := template.ParseFiles("views/parametrosSQ.html")
	id := request.URL.Query().Get("id")
	nombre := request.URL.Query().Get("nombre")
	data := map[string]string{
		"id":     id,
		"nombre": nombre,
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}

	/*
		fmt.Fprintln(response, "-- Parametros ocn String Query -- ")
		fmt.Fprintln(response, request.URL)
		fmt.Fprintln(response, request.URL.RawQuery)
		fmt.Fprintln(response, request.URL.Query())
		fmt.Fprintln(response, request.URL.Query().Get("id"))
		fmt.Fprintln(response, request.URL.Query().Get("nombre"))
	*/
}

type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("views/estructuras.html")
	habilidades := []Habilidad{
		{Nombre: "Ineligencia"},
		{Nombre: "Videojuegos"},
		{Nombre: "Programación"},
		{Nombre: "Canto"},
	}

	if err != nil {
		panic(err)
	} else {
		template.Execute(response, Datos{"Juan perez", 16, 1, habilidades})
	}
}
