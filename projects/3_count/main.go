package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter = 0

func main() {
	http.HandleFunc("/count", countHandler)
	server := http.Server{Addr: ":3333"}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(strconv.Itoa(counter)))
	case http.MethodPost:
		r.ParseForm()
		numberString := r.Form.Get("count")
		number, err := strconv.Atoi(numberString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			w.Write([]byte("это не число"))
			return
		}
		counter += number
	}
}
