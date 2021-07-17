package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id		int		`json:"id"`
	Content	string	`json:"content"`
	Author	string	`json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

// リクエストを正しい関数に振り分けるためのハンドラ関数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 投稿の取り出し
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// リクエストを正しい関数に振り分けるためのハンドラ関数
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// 構造体PostをJSON文字列に組み替え
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	// JSONをResponseWriterに書き出し
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// 投稿の作成
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	// バイト列を作成
	body := make([]byte, len)
	// バイト列にリクエストの本体を読み込み
	r.Body.Read(body)
	var post Post
	// バイト列を構造体Postに組み替え
	json.Unmarshal(body, &post)
	// データベースのレコードを作成
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// 投稿の更新
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// データベースから構造体Postにデータを取得
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	// リクエスト本体からJSONデータを読み出し
	r.Body.Read(body)
	// JSONデータを構造体Postに組み替え
	json.Unmarshal(body, &post)
	// データベースを更新
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// 投稿の削除
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// データベースから構造体Postにデータを取得
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// データベースから投稿データを削除
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
