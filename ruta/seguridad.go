package ruta

import (
	"fmt"
	"goweb1/internal/database"
	"goweb1/modelos"
	"goweb1/pkg/utils"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func FormUsuario(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/UserForm.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func UsuarioListar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/UserList.html", utils.Frontend))

	//Conexion a la db

	database.Conecta()

	query := "select * from usuario"
	usuarios := modelos.Usuarios{}
	datos, err := database.Conexion.Query(query)

	for datos.Next() {
		dato := modelos.Usuario{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Password)
		usuarios = append(usuarios, dato)
	}

	//fmt.Println(clientes)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: ", err)
	}
	defer database.CerrarConexion()

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	data := modelos.HttpUsuario{
		Css:     cssSesion,
		Mensaje: cssMensaje,
		Datos:   usuarios,
	}

	template.Execute(response, data)

}

func UsuarioReceipt(response http.ResponseWriter, request *http.Request) {
	mensaje := ""

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El nombreo está vacío. "
	}
	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + "El correo está vacío. "
	}
	if len(request.FormValue("telefono")) == 0 {
		mensaje = mensaje + "El telefono está vacío. "
	}

	if len(request.FormValue("password")) == 0 {
		mensaje = mensaje + "La contraseña está vacía. "
	}

	if utils.RegexCorreo.FindStringSubmatch(request.FormValue("correo")) == nil {
		mensaje = mensaje + "El correo es inválido. "
	}

	if utils.ValidaPassword(request.FormValue("password")) {
		mensaje = mensaje + "La contraseña es inválida. "
	}

	if mensaje != "" {
		utils.CrearMensaje(response, request, "danger", mensaje)
		http.Redirect(response, request, "/usuario", http.StatusSeeOther)
		return
	}

	database.Conecta()
	query := "insert into usuario values (null, ?, ?, ?, ?);"

	//Generacion del hash
	costo := 8
	bytes, _ := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), costo)
	//fmt.Println(bytes)
	//fmt.Println(string(bytes))

	_, err := database.Conexion.Exec(query,
		request.FormValue("nombre"),
		request.FormValue("correo"),
		request.FormValue("telefono"),
		string(bytes))
	if err != nil {
		fmt.Println(response, err)
	}

	utils.CrearMensaje(response, request, "success", "Se creo el usuario")
	http.Redirect(response, request, "/usuario", http.StatusSeeOther)

}

func Login(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/UserLogin.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}
	template.Execute(response, data)
}

func LoginReceiver(response http.ResponseWriter, request *http.Request) {

	mensaje := ""

	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El nombreo está vacío. "
	}

	if len(request.FormValue("password")) == 0 {
		mensaje = mensaje + "La contraseña está vacía. "
	}

	if mensaje != "" {
		fmt.Println("Campo vacio")
		mensaje = "Usario no existe"

		utils.CrearMensaje(response, request, "danger", mensaje)
		http.Redirect(response, request, "/login", http.StatusSeeOther)

		return

	}
	// Hash de la contraseña

	database.Conecta()
	query := "select * from usuario where nombre = ?"
	datos, err := database.Conexion.Query(query, request.FormValue("nombre"))

	if err != nil {
		fmt.Println("Alguna clase de error: ", datos)
	}
	defer database.CerrarConexion()

	var dato modelos.Usuario
	for datos.Next() {
		errNext := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Password)

		if errNext != nil {
			fmt.Println("Alguna clase de error")
			utils.CrearMensaje(response, request, "danger", "Error al iniciar sesión")
			http.Redirect(response, request, "/login", http.StatusSeeOther)
			return
		}
		// Comprar hashes

		password := []byte(request.FormValue("password"))
		passwordBD := []byte(dato.Password)

		errPass := bcrypt.CompareHashAndPassword(passwordBD, password)

		if errPass == nil {
			session, _ := utils.Store.Get(request, "session-name")
			session.Values["sesionId"] = dato.Id // Puede ser cadena
			session.Values["sesionNombre"] = dato.Nombre

			errSession := session.Save(request, response)

			if errSession != nil {
				http.Error(response, errSession.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(response, request, "/usuario/index", http.StatusSeeOther)

			//fmt.Println("Acceso de: ", dato.Nombre)
			//utils.CrearMensaje(response, request, "success", "Acceso concedido")
			//http.Redirect(response, request, "/usuario/index", http.StatusSeeOther)
			return
		}
	}
	fmt.Println("Acceso denegado")
	utils.CrearMensaje(response, request, "danger", "Acceso denegado")
	http.Redirect(response, request, "/login", http.StatusSeeOther)

}

func UsuarioIndex(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/UserIndex.html", utils.Frontend))
	template.Execute(response, nil)
}

func SeguridadSesion(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/UserIndex.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	sesionId, sesionNombre := utils.RetornaLogin(request)
	data := map[string]string{
		"css":          cssSesion,
		"mensaje":      cssMensaje,
		"sesionId":     sesionId,
		"sesionNombre": sesionNombre,
	}

	template.Execute(response, data)
}

func LogOut(response http.ResponseWriter, request *http.Request) {

}
