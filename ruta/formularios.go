package ruta

import (
	"fmt"
	"goweb1/pkg/utils"
	"net/http"
	"text/template"
)

func Formulario(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/formulario.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}
	template.Execute(response, data)
}

func Formulariop(response http.ResponseWriter, request *http.Request) {
	mensaje := ""

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo nombre está vacio\n"
	}
	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + "El campo correo está vacío\n"
	}
	if utils.RegexCorreo.FindStringSubmatch(request.FormValue("correo")) != nil {
		mensaje = mensaje + "Correo inválido\n"
	}
	if !utils.ValidaPassword(request.FormValue("password")) {
		mensaje = mensaje + "Error de contraseña\n"
	}
	if mensaje != "" {
		//fmt.Fprintln(response, mensaje)
		//return
		//fmt.Fprintln(response, mensaje)
		utils.CrearMensaje(response, request, "danger", mensaje)
		http.Redirect(response, request, "/formulario", http.StatusSeeOther)
	}
	fmt.Fprintln(response, "Nombre: "+request.FormValue("nombre"))
	fmt.Fprintln(response, "Teléfono: "+request.FormValue("telefono"))
	fmt.Fprintln(response, "Correo: "+request.FormValue("correo"))
	fmt.Fprintln(response, "Password: "+request.FormValue("password"))
}

/*

func Formulariop(reponse http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(reponse, "Nombre: "+request.FormValue("nombre"))
	fmt.Fprintln(reponse, "Teléfono: "+request.FormValue("telefono"))
	fmt.Fprintln(reponse, "Correo: "+request.FormValue("correo"))
	fmt.Fprintln(reponse, "Password: "+request.FormValue("password"))
}
*/
