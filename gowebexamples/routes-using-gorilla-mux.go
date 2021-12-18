package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // https://pkg.go.dev/github.com/gorilla/mux
)

// gorilla/muxパッケージを使用して、名前付きパラメーター、GET / POSTハンドラー、およびドメイン制限を使用してルートを作成
func main() {
    r := mux.NewRouter() //新しいリクエストルーターの作成

		// リクエストハンドラーの登録
    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r) 
        title := vars["title"] // the book title slug
        page := vars["page"] // the page

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r) // 独自のルータを使用するには、nil を r の変数に置き換える
}

// gorilla/mux Router機能
// Methods
// リクエストハンドラを特定のHTTPメソッドに制限する

// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")


// Hostnames & Subdomains
// リクエストハンドラを特定のホスト名またはサブドメインに制限する

// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")


// Schemes
// リクエストハンドラをhttp / httpsに制限する

// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")


// Path Prefixes & Subrouters
// リクエストハンドラを特定のパスプレフィックスに制限する

// bookrouter := r.PathPrefix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBooks)
// bookrouter.HandleFunc("/{title}", GetBook)