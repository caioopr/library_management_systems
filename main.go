package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	rtr := router.Generate()

	fmt.Printf("Server runing on port: %d", config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), rtr)

	if err != nil {
		log.Fatal(err)
	}
}
