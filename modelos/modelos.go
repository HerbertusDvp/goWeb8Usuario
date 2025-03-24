package modelos

type Cliente struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
}

type Clientes []Cliente

type Usuario struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
	Password string
}

type Usuarios []Usuario

type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Categorias []Categoria

type ClienteHttp struct {
	Css     string
	Mensaje string
	Datos   Clientes
}

type ClienteHttp2 struct {
	Css     string
	Mensaje string
	Datos   Cliente
}

type HttpUsuario struct {
	Css     string
	Mensaje string
	Datos   Usuarios
}
