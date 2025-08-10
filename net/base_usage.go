package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册一个 getuser 路由, 写入 Tom 文本
	http.HandleFunc("/getuser", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Tom")
	})

	// 监听 8080 端口并启动服务
	http.ListenAndServe(":8080", nil)
}
