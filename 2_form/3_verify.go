package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("0.0.0.0:1111", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var (
		name string
	)

	// fmt.Print(r.Form["nickname"])

	if r.Method == "POST" {

		if nickname, ok := r.Form["nickname"]; ok {
			name = nickname[0]
		}

		if m, _ := regexp.MatchString("^\\p{Han}+$", name); !m {
			fmt.Println("不是中文")
		} else {
			fmt.Println("是中文")
		}

		if e, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, r.Form.Get("email")); !e {
			fmt.Println("不是邮箱")
		} else {
			fmt.Println("是邮箱")
		}

		// mobile
		if mobile, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !mobile {
			//
		}

		// md5
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)
	}

	fmt.Fprintf(w, "Name = %s\n", name)
}
