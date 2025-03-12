package ruta

import (
	"encoding/json"
	"fmt"
	"goweb1/modelos"
	"goweb1/pkg/utils"
	"io"
	"net/http"
	"text/template"
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
	fmt.Println(datos)
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
