package middleware

import (
	"goweb1/pkg/utils"
	"net/http"
)

func Proteger(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		session, _ := utils.Store.Get(request, "session-name")

		if session.Values["sesionId"] != nil {
			next.ServeHTTP(response, request)
		} else {
			utils.CrearMensaje(response, request, "warning", "Debes estar logueado")
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		}
	}
}

/*
func Proteccion(next http.HandlerFunc) http.HandlerFunc  {
	return func(response http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(response, request)
	}
}
*/
