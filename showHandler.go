package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const layout = "2006-01-02 15:04:05"

type showHandler struct{}

func timeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func (p *showHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db := connectDB()

	accesses := []Access{}
	err := db.Where("target LIKE ?", "%" + vars["target"] + "%").Order("id desc").Find(&accesses).Error

	if err != nil {
		fmt.Fprintln(w,"error")
	}

	fmt.Fprintln(w, len(accesses))
	fmt.Fprintln(w, "--------")

	i := 0
	for _, access := range accesses {
		fmt.Fprintln(w, timeToString(access.Date))
		if i > 100 {
			break
		}
	}
}
