package main

import (
	"net/http"
	// "net/http/cgi"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")

	router.Handle(rootPath+"/ping", &pingHandler{})

	server := &http.Server{
		Addr:    ":9999",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
	// cgi.Serve(router)
}
