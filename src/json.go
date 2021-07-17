package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id			int			`json:"id"`
	Content		string		`json:"content"`
	Author		Author		`json:"author"`
	Comments	[]Comment	`json:"comments"`
}

type Author struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
}

type Comment struct {
	Id			int		`json:"id"`
	Content		string	`json:"content"`
	Author		string	`json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON File:", err)
		return
	}
	defer jsonFile.Close()

	// jsonData, err := ioutil.ReadAll(jsonFile)
	// if err != nil {
	// 	fmt.Println("Error reading JSON File", err)
	// 	return
	// }

	// var post Post
	// json.Unmarshal(jsonData, &post)
	// fmt.Println(post)

	// jsonデータからデコーダを生成する
	decoder := json.NewDecoder(jsonFile)
	// EOFが検出されるまで繰り返す
	for {
		var post Post
		err := decoder.Decode(&post) // jsonデータをデコードし構造体に収納する
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}

}