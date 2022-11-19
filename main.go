package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pitabwire/frame"
	"io"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method : %s /\n", r.Method)
	fmt.Printf("query : %s\n", r.URL.Path)
	fmt.Printf("content-type : %s\n", r.Header.Get("content-type"))
	fmt.Printf("headers :\n")
	for headerName, headerValue := range r.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error : could not read request body: %s\n", err)
	} else {
		fmt.Printf("request body: %s\n", reqBody)
	}
	fmt.Print("\n\n\n----------------------------  emoji_sparkles -----------------------------------\n\n\n\n")
	_, _ = fmt.Fprintf(w, `{"message": "well received emoji_sparkles"}`)
}

func main() {

	serviceName := "log_http_service"
	ctx := context.Background()

	router := mux.NewRouter().PathPrefix("/").HandlerFunc(handler)

	server := frame.HttpHandler(router.GetHandler())
	service := frame.NewService(serviceName, server)
	err := service.Run(ctx, ":8080")
	if err != nil {
		log.Fatal("main -- Could not run Server : %v", err)
	}

}
