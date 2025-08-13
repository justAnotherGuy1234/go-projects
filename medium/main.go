package main

import (
	"fmt"
	"medium/config"
	"net/http"
)

func main() {
	config.SetUpDb()

	server := http.Server{
		Addr: ":5000",
	}
	fmt.Println("server started at port ", server.Addr)
	server.ListenAndServe()
}
