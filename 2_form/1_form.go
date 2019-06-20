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

// 1) r.Form 返回一个url.Values类型的值, 如果字段不存在, 返回一个长度为0的集合
// 2) r.Form.Get() 只返回第一个叫xxx的值, 如果字段不存在, 返回一个空字符串

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
