package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("0.0.0.0:7777", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 获取请求信息
	w.Write([]byte("path: " + r.URL.Path + " host: " + r.Host + " method: " + r.Method))
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	// 请求url
	println(strings.Join([]string{scheme, r.Host, r.RequestURI}, ""))
}
