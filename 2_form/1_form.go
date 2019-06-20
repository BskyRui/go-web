package main

import (
	"fmt"
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
	r.ParseForm()
	// fmt.Println(r.Form)
	if name, ok := r.Form["name"]; ok {
		fmt.Println(name[0])
	}

	for k, v := range r.Form {
		fmt.Println(k, "val: "+strings.Join(v, ""))
	}

	// 处理同名字段
	if username, ok := r.Form["username"]; ok {
		for k, v := range username {
			fmt.Println(k, v)
		}
	}

	uid := r.Form["uid"]
	fmt.Fprintf(w, "Name = %s\n", uid)
}
