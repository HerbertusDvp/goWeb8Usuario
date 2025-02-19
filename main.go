package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(reponse, "Hola web con go")
	})
	log.Fatal(http.ListenAndServe("localhost:8081", nil))

}
