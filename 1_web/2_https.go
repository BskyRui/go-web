package main

// 1) 生成密钥文件: openssl genrsa -out server.key 2048
// 2) 生成证书文件: openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365

// 证书类型:
// 域名型SSL证书(DV SSL) 自动签发
// 企业型SSL证书(OV SSL) 提供合法手续就能签发
// 增强型SSL证书(EV SSL) 银行或金融机构

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":7777", "./server.crt", "./server.key", nil); err != nil {
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
