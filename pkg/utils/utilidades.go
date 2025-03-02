package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
	gomail "gopkg.in/gomail.v2"
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

func EnviarCorreo() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "trabajosteschi@gmail.com")
	msg.SetHeader("To", "HerbertusDvp@gmail.com")
	msg.SetHeader("Subject", "Curso de golang")
	msg.SetBody("text/html", "<h1>Curso de Golang</h1><b>Texto en negritas</b><p>Este es un parrafo</p>")
	//msg.Attach()
	n := gomail.NewDialer("smtp.gmail.com", 587, "trabajosteschi@gmail.com", "pzphnxowayjoekrw")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
