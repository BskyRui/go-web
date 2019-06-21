package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("0.0.0.0:1111", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 设置cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	// 过期时间, 不设置的话就是临时cookie, 在关闭浏览器后会被清除
	cookie := http.Cookie{Name: "username", Value: "xxxx", Expires: expiration}
	http.SetCookie(w, &cookie)

	c, _ := r.Cookie("username")
	fmt.Fprint(w, c)

	// for _, cookie := range r.Cookies()
}
