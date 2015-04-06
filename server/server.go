package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const directoryRoot = "/tmp/ansible_share"

func main() {
	createDirectory(directoryRoot)
	r := mux.NewRouter()
	r.HandleFunc("/roles/{role}/{tag}", UploadRoleHandler).Methods("POST")
	r.HandleFunc("/roles/{role}/{tag}", DownloadRoleHandler).Methods("GET")
	r.HandleFunc("/_ping", Ping).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func Ping(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "ping")
}
