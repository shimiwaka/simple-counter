package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gorilla/mux"
)

type hitHandler struct{}

func (p *hitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db := connectDB()

	access := Access{}
	access.Date = time.Now()
	access.Target = vars["target"]

	db.Create(&access)
	fmt.Fprintln(w, "ok")
}
