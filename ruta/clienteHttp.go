package ruta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goweb1/modelos"
	"goweb1/pkg/utils"
	"io"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTc0MTgxNDMxNiwiZXhwIjoxNzQ0NDA2MzE2fQ.TAFYPufRL2gJPL117USamfhYCOun2Syz3n4O74vQiPA"

func ClienteHttp(response http.ResponseWriter, request *http.Request) {
	//cargar la plantilla
	template := template.Must(template.ParseFiles("web/templates/clienteHttp.html", utils.Frontend))

	// Crear la solicitud HTTP
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias", nil)

	if err != nil {
		http.Error(response, "Error al crear la solicitud HTTP", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Solicitud http: Ok")
	}
	//Agregar el token de autorizacion
	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)

	if err != nil {
		http.Error(response, "Error al conectarse a la API", http.StatusInternalServerError)
		return
	}
	defer reg.Body.Close()

	//verificar el codigo de estado
	if reg.StatusCode != http.StatusOK {
		http.Error(response, fmt.Sprintf("Error en la API: %s", reg.Status), reg.StatusCode)
		return
	} else {
		fmt.Println("Verificación del codigo: Ok")
	}

	//Leer el cuerpo de la respuesta
	body, err := io.ReadAll(reg.Body)

	if err != nil {
		http.Error(response, "Error al leer la respuesta de la API", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Leer el cuerpo de la respuesta: Ok")
	}

	//Decodificar JSON
	datos := modelos.Categorias{} // Slice de Categoria: Id, Nombre Slug
	fmt.Println("Impresion de datos")
	errJson := json.Unmarshal(body, &datos)
	//fmt.Println(datos)
	if errJson != nil {
		http.Error(response, "Error al decodificar la respuesta JSON", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Decodificar jason: OK")
	}

	//Pasar datos a la plantilla
	data := map[string]modelos.Categorias{
		"datos": datos,
	}

	//renderizar la plantilla
	err = template.Execute(response, data)

	if err != nil {
		http.Error(response, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("renderizar plantilla: Ok")
	}
}

func ClienteHttpCrear(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/clienteHttpCrear.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)

}

func ClienteHttpCrearPost(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo nombre está vacío"
	}

	if mensaje != "" {
		utils.CrearMensaje(response, request, "danger", mensaje)
		http.Redirect(response, request, "/clientehttp/crear", http.StatusSeeOther)
	}

	datos := map[string]string{"nombre": request.FormValue("nombre")}

	//Conversion a jason
	jsonValue, _ := json.Marshal(datos)

	// Crear la solicitud HTTP
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.api.tamila.cl/api/categorias", bytes.NewBuffer(jsonValue))

	if err != nil {
		http.Error(response, "Error al crear la solicitud HTTP", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Solicitud http: Ok")
	}
	//Agregar el token de autorizacion
	req.Header.Set("Authorization", Token)

	reg, _ := client.Do(req)
	defer reg.Body.Close()

	utils.CrearMensaje(response, request, "success", "Registro exitoso")
	http.Redirect(response, request, "/clientehttp/crear", http.StatusSeeOther)
}

func ClienteHttpEditar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/clienteHttpEditar.html", utils.Frontend))

	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)
	defer reg.Body.Close()
	body, err := io.ReadAll(reg.Body)
	datos := modelos.Categoria{}
	errJson := json.Unmarshal(body, &datos)

	if errJson != nil {

	}

	data := map[string]string{
		"id":     vars["id"],
		"nombre": datos.Nombre,
		"slug":   datos.Slug,
	}

	template.Execute(response, data)

}

func ClienteHttpEditarPost(response http.ResponseWriter, request *http.Request) {
	mensaje := ""

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo nombre está vacío"
	}

	if mensaje != "" {
		fmt.Println("Nombre correcto")
	}

	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)
	defer reg.Body.Close()
	body, err := io.ReadAll(reg.Body)

	datos := modelos.Categoria{}
	errJson := json.Unmarshal(body, &datos)

	if errJson != nil {

	}

	datosJson := map[string]string{
		"nombre": request.FormValue("nombre"),
	}
	//Edicion del registro
	jsonValue, _ := json.Marshal(datosJson)
	req2, err2 := http.NewRequest("PUT", "https://www.api.tamila.cl/api/categorias/"+vars["id"], bytes.NewBuffer(jsonValue))
	req2.Header.Set("Authorization", Token)
	if err2 != nil {
		fmt.Println(err2)
	}

	reg2, err3 := client.Do(req2)
	defer reg.Body.Close()

	if err3 != nil {

	}

	defer reg2.Body.Close()
	http.Redirect(response, request, "/clientehttp/editar/"+vars["id"], http.StatusSeeOther)
}

func ClienteHttpEliminar(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)
	defer reg.Body.Close()

	body, err := io.ReadAll(reg.Body)

	datos := modelos.Categoria{}
	errJson := json.Unmarshal(body, &datos)
	if errJson != nil {
		fmt.Println("Despues del modelos.categoria: ", errJson)
	}

	req2, err2 := http.NewRequest("DELETE", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	req2.Header.Set("Authorization", Token)
	if err2 != nil {
		fmt.Println("Error 2: ", err2)
	}

	reg2, err3 := client.Do(req2)
	defer reg.Body.Close()

	if err != nil {
		fmt.Println("Error 3: ", err3)
	}
	defer reg2.Body.Close()
	http.Redirect(response, request, "/clientehttp", http.StatusSeeOther)
}
