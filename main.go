package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()

	fmt.Println("Rodando API")
	log.Fatal(http.ListenAndServe(":3000", r))
}
