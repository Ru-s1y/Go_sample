// テストファイルはテストされる関数と同じパッケージに置く
package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	// テストされる関数の呼び出し
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	// テストが予想通りかチェックし、違えばエラーメッセージを設定
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	// テストが予想通りかチェックし、違えばエラーメッセージを設定
	if post.Content != "Hello World" {
		t.Error("Wrong content, was expecting 'Hello World' but got", post.Content)
	}
}

// テストを全て省略
func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}