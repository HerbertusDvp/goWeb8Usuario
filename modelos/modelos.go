package modelos

type Cliente struct {
	Id     int
	Nombre string
	Correo string
}

type Clientes []Cliente

type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Categorias []Categoria
