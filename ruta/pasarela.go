package ruta

import (
	"goweb1/pkg/utils"
	"net/http"
	"text/template"
)

func PasarelaHomePay(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoHome.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func PasarelaWebPay(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoWebPay.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func PasarelaPayPal(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoPayPal.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}
