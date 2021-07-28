package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:     "127.0.0.1:8080",
    Handler:  mux,
	}

  TILESDB = tilesDB()
  fmt.Println("Mosaix server started.")
  server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("upload.html")
  t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
  t0 := time.Now()

  r.ParseMultipartForm(104857760)
  // アップロードされたファイルとタイルサイズを取得
  file, _, _ := r.FormFile("image")
  defer file.Close()
  tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

  // アップロードされたターゲット画像をデコード
  original, _, _ := image.Decode(file)
  bounds := original.Bounds()

  newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))

  // タイル画像データベースを複製
  db := cloneTilesDB()

  // 各タイルの基準ピクセルを設定
  sp := image.Point{0, 0}

  // ターゲット画像を巡回
  for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSie {
    for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
      r, g, b, _ := original.At(x, y).RGBA()
      color := [3]float64{float64(r), float64(g), float64(b)}

      nearest := nearest(color, &db)
      file, err := os.Open(nearest)
      if err == nil {
        img, _ := image.Decode(file)
        if err == nil {
          t := resize(img, tileSize)
          tile := t.SubImage(t.Bounds())
          tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)

          draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
        } else {
          fmt.Println("error: ", nearest)
        }
      } else {
        fmt.Println("error: ", nearest)
      }
      file.Close()
    }
  }

  // JPEGにエンコードし、base64文字列に変換してブラウザに送信
  buf1 := new(bytes.Buffer)
  jpeg.Encode(buf1, original, nil)
  originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

  buf2 := new(bytes.Buffer)
  jpeg.Encode(buf2, original, nil)
  mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
  t1 := time.Now()
  images := map[string]string{
    "original": originalStr,
    "mosaic": mosaic,
    "duration": fmt.Sprintf("%v ", t1.Sub(t0)),
  }
  t, _ := template.ParseFiles("results.html")
  t.Execute(w, images)
}