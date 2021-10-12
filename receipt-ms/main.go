package main

import "net/http"

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Test"})
	http.ListenAndServe(":5000", nil)
}
