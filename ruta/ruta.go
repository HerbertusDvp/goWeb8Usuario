package ruta

import (
	"goweb1/pkg/utils"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/home2.html", utils.Frontend))
	template.Execute(response, nil)

	/*
		template, err := template.ParseFiles("web/templates/home2.html", "web/layout/frontend.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}
	*/
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("web/templates/nosotros.html", utils.Frontend)

	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("web/templates/parametros.html", utils.Frontend)
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

	template, err := template.ParseFiles("web/templates/parametrosSQ.html", utils.Frontend)
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
	template, err := template.ParseFiles("web/templates/estructuras.html", utils.Frontend)
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

func Pagina404(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/404.html", utils.Frontend))
	template.Execute(response, nil)

	/*
		template, err := template.ParseFiles("web/templates/home2.html", "web/layout/frontend.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}
	*/
}
