package base

import "net/http"

func httpserver() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8989", nil)
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	head0 := req.Header.Get("0")
	body := make([]byte, 100)
	req.Body.Read(body)
	println("head: ", head0)
	println("body: ", string(body))
}
