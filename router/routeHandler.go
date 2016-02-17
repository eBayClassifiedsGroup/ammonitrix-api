package router

import (
	"fmt"
	"net/http"
)

func (r *Receiver) Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Index")
}

func (r *Receiver) DataIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "DataIndex")
}

func (r *Receiver) DataGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "DataGet")
}
