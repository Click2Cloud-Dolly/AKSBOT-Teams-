package main

import (
	x "aksbot/middleware"
	"fmt"
	"log"
	"net/http"
)

//var resume int

func main() {

	//http.HandleFunc("/ws", k8s.Request)
	//fmt.Println("Starting server on port:3979...")
	//log.Fatal(http.ListenAndServe(":3979", nil))

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ws", x.WsEndpoint)

	serverMux.HandleFunc("/http", x.HttpEndpoint)

	fmt.Println("Websocket server started")

	server := &http.Server{
		Addr:    "localhost:8082",
		Handler: serverMux,
	}
	fmt.Println("Http Server started")
	log.Print(server.ListenAndServe())
	fmt.Println("Starting port on:8082")
}
