package main

import (
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/dump", DumpHandler)
	http.ListenAndServe(":9999", nil)
}

func DumpHandler(res http.ResponseWriter, req *http.Request) {
	result, err := httputil.DumpRequest(req, true)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write(result)
}