package main

import (
	"code/web/3_cookie_session/session"
	_ "code/web/3_cookie_session/session/providers/memory"
	"log"
	"net/http"
	"text/template"
)

var globalSessions *session.Manager

func init() {
	// 获取key为memory的provider
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/count", handler)
	if err := http.ListenAndServe("0.0.0.0:7777", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 请求 / 的时候 chrome浏览器会请求favicon.ico导致被访问了2次
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sess := globalSessions.SessionStart(w, r)
		ct := sess.Get("countnum")

		if ct == nil {
			sess.Set("countnum", 1)
		} else {
			sess.Set("countnum", (ct.(int) + 1))
		}
		t, _ := template.ParseFiles("./count.html")
		w.Header().Set("Content-Type", "text/html")

		t.Execute(w, sess.Get("countnum"))

	}
}
