package server

import (
	"fmt"
	"http"
	"https://github.com/gorilla/mux"
)

func main() {
	

	r := mux.NewRouter()
    fileServer := http.FileServer(http.Dir("./static")) // New code
    r.Handle("/", fileServer) // New code
	r.Handle("/thanksyou", NotImplemented).Methods("POST")
	r.Handle("/login", creatUser).Methods("POST")
	r.HandleFunc("/quiz", NotImplemented).Methods("Get")
    r.HandleFunc("/quiz/answer", NotImplemented).Methods("POST")


    fmt.Printf("Starting server at port 3000\n")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}