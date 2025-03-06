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

var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTY4MzE1NTM1NiwiZXhwIjoxNjg1NzQ3MzU2fQ.eEZZHIqiM5FpR8ZwK3jPd-qT367epSK5qjoHU9f7r1I"

func ClienteHttp(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/clienteHttp.html", utils.Frontend))
	//cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	// conexion a la api
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias", nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", Token)
	reg, _ := client.Do(req)
	defer reg.Body.Close()
	fmt.Println(reg.Status)
	//COnvertir la informacion a slice
	body, err := io.ReadAll(reg.Body)

	if err != nil {
		fmt.Println("Error con el ReadAll")
	}

	datos := modelos.Categorias{}
	errJson := json.Unmarshal(body, &datos)

	if errJson != nil {
	}

	data := map[string]modelos.Categorias{
		"body": datos,
	}

	template.Execute(response, data)
}
