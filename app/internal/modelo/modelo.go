package modelo

type Cliente struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
	Mensaje  string `json:"mensaje"`
}

type Clientes []Cliente
