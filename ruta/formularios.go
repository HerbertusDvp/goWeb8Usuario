package ruta

import (
	"fmt"
	"goweb1/pkg/utils"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
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

func FormularioFile(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/formulariofile.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}
	template.Execute(response, data)
}

func FormularioFileUp(response http.ResponseWriter, request *http.Request) {

	file, handler, err := request.FormFile("archivo")
	if err != nil || file == nil {
		utils.CrearMensaje(response, request, "danger", "No se detectó ningun archivo")
		http.Redirect(response, request, "/formulario/file", http.StatusSeeOther)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]

	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "web/static/uploads/" + foto
	f, errCopy := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)

	if errCopy != nil {
		utils.CrearMensaje(response, request, "danger", "Error al guardar el archivo - if 2")
		http.Redirect(response, request, "/formulario/file", http.StatusSeeOther)
		return
	}
	_, errCopiar := io.Copy(f, file)
	if errCopiar != nil {
		utils.CrearMensaje(response, request, "danger", "Error al guardar el archivo - if 3")
		http.Redirect(response, request, "/formulario/file", http.StatusSeeOther)
		return
	}
	utils.CrearMensaje(response, request, "success", "Se subió el archivo")
	http.Redirect(response, request, "/formulario/file", http.StatusSeeOther)

}
