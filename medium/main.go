package main

import (
	"fmt"
	"medium/config"
	"medium/controller"
	"medium/router"
	"net/http"
)

func main() {
	db, err := config.SetUpDb()

	if err != nil {
		fmt.Println("error setting up db", err)
	}

	uc := controller.NewController(db)
	ur := router.NewRouter(uc)

	bc := controller.NewBlogController(db)
	br := router.NewBlogRouter(bc)

	mainRouter := router.SetupRouter(ur, br)

	server := http.Server{
		Addr:    ":5000",
		Handler: mainRouter,
	}
	fmt.Println("server started at port ", server.Addr)
	server.ListenAndServe()
}
