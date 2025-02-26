package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Frontend string = "web/layout/frontend.html"
var Store = sessions.NewCookieStore([]byte("session-name"))

func RetornaMensaje(response http.ResponseWriter, request *http.Request) (string, string) {
	session, _ := Store.Get(request, "flash-session")

	fm := session.Flashes("css")
	session.Save(request, response)
	cssSesion := ""
	if len(fm) == 0 {
		cssSesion = ""
	} else {
		cssSesion = fm[0].(string)
	}

	fm2 := session.Flashes("mensaje")
	session.Save(request, response)
	cssMensaje := ""

	if len(fm2) == 0 {
		cssMensaje = ""
	} else {
		cssMensaje = fm2[0].(string)
	}

	return cssSesion, cssMensaje
}

func CrearMensaje(response http.ResponseWriter, request *http.Request, css string, mensaje string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(css, "css")
	session.AddFlash(mensaje, "mensaje")
	session.Save(request, response)
}
