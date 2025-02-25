package ruta

import (
	"fmt"
	"goweb1/pkg/utils"
	"net/http"
	"text/template"
)

func Formulario(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/formulario.html", utils.Frontend))
	template.Execute(response, nil)
}

func Formulariop(reponse http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(reponse, "Nombre: "+request.FormValue("nombre"))
	fmt.Fprintln(reponse, "Teléfono: "+request.FormValue("telefono"))
	fmt.Fprintln(reponse, "Correo: "+request.FormValue("correo"))
	fmt.Fprintln(reponse, "Password: "+request.FormValue("password"))
}
