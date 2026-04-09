package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WebApi() {
	http.HandleFunc("/", respondWithHello)
	http.HandleFunc("/post", respondWithJson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func respondWithHello(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Hi there, I love the route %s!", request.URL.Path[1:])
	if err != nil {
		return
	}
}

type MyResponse struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
}

func respondWithJson(writer http.ResponseWriter, request *http.Request) {
	log.Printf("got %s request\n", request.Method)

	response, _ := json.Marshal(&MyResponse{Message: "foo", Number: 42})
	_, err := fmt.Fprintf(writer, string(response))
	if err == nil {
		fmt.Print(fmt.Errorf("error while writing response: %s\n", err))
	}
}
