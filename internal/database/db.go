package conect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Conexion *sql.DB

func Conecta() {
	errorVariables := godotenv.Load()

	if errorVariables != nil {
		panic(errorVariables)
	}

	con, err := sql.Open(
		"mysql",
		os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+"/"+os.Getenv("DB_NAME"))

	if err != nil {
		panic(err)
	}

	Conexion = con
}

func CerrarConexion() {
	Conexion.Close()
}

//db, err := sql.Open("mysql", "user:password@/dbname")
