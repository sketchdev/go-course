package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"hello/pkg/picker"
	"net/http"
	"strconv"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/random", func(rw http.ResponseWriter, req *http.Request) {

		rw.Header().Add("content-type", "application/json")

		if size := req.FormValue("size"); size == "" {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte("Bad request, size is required"))
			return
		} else {

			s, _ := strconv.Atoi(size)
			students := picker.Random(s)

			b, _ := json.Marshal(students)
			_, _ = rw.Write(b)
		}
	})

	server := &http.Server{
		Handler: router,
		Addr: ":8000",
	}

	fmt.Println("Server is up and listening...")
	_ = server.ListenAndServe()

}
