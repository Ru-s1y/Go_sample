package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id		int
	Content	string
	Author	string
}

// DBとの接続
var Db *sql.DB
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// 投稿複数件の取得
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// 投稿一件の取得
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// 新規投稿の作成
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id) // Scan()の引数にはSQL実行後の返り値IDが入る
	return
}

// 投稿の更新
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// 投稿の削除
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	fmt.Println(post)	// {0 Hello World! Sau Sheong} と表示されるはず
	post.Create()		// 新規作成
	fmt.Println(post)	// {1 Hello World! Sau Sheong} と表示されるはず

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
}