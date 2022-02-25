package main

import (
	"fmt"
	"log"

	"net/http"

	"rest.com/pkg/server"
	//"github.com/friendsofgo/gopher-api/pkg/server"
	//"github.com/friendsofgo/gopherapi/pkg/server"
)

func main() {
	fmt.Println("Hello World!")

	s := server.New()
	//sv2 := server.New2()

	log.Fatal(http.ListenAndServe(":8080", s.Router()))
	//log.Fatal(http.ListenAndServe(":8090", sv2.Router()))

	//respuesta := http.ListenAndServe(":8080", s.Router())
	//fmt.Println("Error" + respuesta.Error())
}
