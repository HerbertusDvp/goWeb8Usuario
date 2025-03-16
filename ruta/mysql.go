package ruta

import (
	"goweb1/pkg/utils"
	"net/http"
	"text/template"
)

func MysqlListar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/mysqlHome.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)

}
