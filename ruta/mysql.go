package ruta

import (
	"fmt"
	"goweb1/internal/database"
	"goweb1/modelos"
	"goweb1/pkg/utils"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func MysqlListar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/mysqlHome.html", utils.Frontend))

	//Conexion a la db

	database.Conecta()

	query := "select * from cliente"
	clientes := modelos.Clientes{}
	datos, err := database.Conexion.Query(query)

	for datos.Next() {
		dato := modelos.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}

	//fmt.Println(clientes)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: ", err)
	}
	defer database.CerrarConexion()

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	data := modelos.ClienteHttp{
		Css:     cssSesion,
		Mensaje: cssMensaje,
		Datos:   clientes,
	}

	template.Execute(response, data)

}

func MysqlCrear(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/mysqlCrear.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func MysqlCrearRecept(response http.ResponseWriter, request *http.Request) {
	mensaje := ""

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El nombre está vacío "
	}

	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + "El correo está vacío "
	}

	if utils.RegexCorreo.FindStringSubmatch(request.FormValue("correo")) == nil {
		mensaje = mensaje + "Correo inválido"
	}

	if len(request.FormValue("telefono")) == 0 {
		mensaje = mensaje + "El telefono está vacío "
	}

	if mensaje != "" {
		utils.CrearMensaje(response, request, "danger", mensaje)
		http.Redirect(response, request, "/mysql", http.StatusSeeOther)
		return
	}

	database.Conecta()
	query := "Insert into cliente values (null, ?,?,?);"
	_, err := database.Conexion.Exec(query, request.FormValue("nombre"), request.FormValue("correo"), request.FormValue("telefono"))

	if err != nil {
		fmt.Println(response, err)
	}

	utils.CrearMensaje(response, request, "success", "Se creó el registro")
	http.Redirect(response, request, "/mysql", http.StatusSeeOther)
}

func MysqlEditar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/mysqlEditar.html", utils.Frontend))

	database.Conecta()

	query := "Select*from cliente where id = ?"

	vars := mux.Vars(request)
	datos, err := database.Conexion.Query(query, vars["id"])

	if err != nil {
		fmt.Println(err)
	}

	defer database.CerrarConexion()

	var dato modelos.Cliente
	for datos.Next() {
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)

		if err != nil {
			log.Fatal(err)
		}
	}

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := modelos.ClienteHttp2{
		Css:     cssSesion,
		Mensaje: cssMensaje,
		Datos:   dato,
	}
	template.Execute(response, data)

}

func MysqlEditarRecept(response http.ResponseWriter, request *http.Request) {
	mensaje := ""

	vars := mux.Vars(request)

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El Nombre está vacío "
	}

	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + "El Correo está vacío "
	}
	if utils.RegexCorreo.FindStringSubmatch(request.FormValue("correo")) == nil {
		mensaje = mensaje + ". El correo no es válido"
	}
	if len(request.FormValue("telefono")) == 0 {
		mensaje = mensaje + ". El teléfono está vacío"
	}

	if mensaje != "" {
		utils.CrearMensaje(response, request, "danger", mensaje)
		http.Redirect(response, request, "/mysql/editar/"+vars["id"], http.StatusSeeOther)
	}
	database.Conecta()

	query := "update cliente set nombre = ?, correo = ?, telefono = ? where id = ?; "

	_, err := database.Conexion.Exec(query, request.FormValue("nombre"), request.FormValue("correo"), request.FormValue("telefono"), vars["id"])

	if err != nil {
		fmt.Println(err)
	}

	utils.CrearMensaje(response, request, "success", "Se modificó el cliente")
	http.Redirect(response, request, "/mysql", http.StatusSeeOther)
}

func MysqlEliminar(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	database.Conecta()

	query := "delete from cliente where id = ?"

	_, err := database.Conexion.Exec(query, vars["id"])

	if err != nil {
		fmt.Println(err)
	}
	utils.CrearMensaje(response, request, "success", "Se eliminó el cliente")
	http.Redirect(response, request, "/mysql", http.StatusSeeOther)
}
