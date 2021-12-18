package main

import (
	"fmt"
	"net/http"
)

func main() {
	  //リクエストハンドラーの登録 
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

		// HTTP接続をリッスン
    http.ListenAndServe(":8080", nil)
}