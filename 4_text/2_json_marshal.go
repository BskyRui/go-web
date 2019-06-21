package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	// ID 不会导出到 JSON 中
	ID int `json:"-"`

	// ServerName2 的值会进行二次 JSON 编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到 JSON 串中
	ServerIP string `json:"serverIP,omitempty"`
}

func main() {
	s := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	// os.Stdout.Write(b)
	fmt.Println(string(b))
}
