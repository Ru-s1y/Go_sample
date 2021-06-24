package main

import (
	"fmt"
	"encoding/json"
	"encoding/base64"
	"net/http"
	"time"
	"html/template"
	// "io/ioutil"
	"math/rand"
)

type Post struct {
	User	string
	Threads	[]string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func headers(w http.ResponseWriter, r  *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap { "fdate": formatDate }
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("web/tmpl.html")
	t.Execute(w, time.Now())
}

func processContext(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/cxt_tmpl.html")
	content := `I asked: <i>"What's up?</i>"`
	t.Execute(w, content)
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "そのようなサービスはありません。他をあたってください。")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:		"Sau Sheong",
		Threads:	[]string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:		"first_cookie",
		Value:		"Go Web Programing",
		HttpOnly:	true,
	}
	c2 := http.Cookie{
		Name:		"second_cookie",
		Value:		"Manning Publications Co",
		HttpOnly:	true,
	}
	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cl, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, cl)
	fmt.Fprintln(w, cs)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:	"flash",
		Value:	base64.URLEncoding.EncodeToString(msg),	// ヘッダ内ではクッキの値をURLエンコードする必要がある
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintf(w, "メッセージがありません")
		}
	} else {
		rc := http.Cookie{
			Name:		"flash",
			MaxAge:		-1,
			Expires:	time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func processForm(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("X-XSS-Protection", "0") // JavaScript 埋め込み用
	t, _ := template.ParseFiles("web/form_tmpl.html")
	// t.Execute(w, template.HTML(r.FormValue("comment"))) // JavaScript 埋め込み用
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/form.html")
	t.Execute(w, nil)
}

func processLayout(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("web/layout.html", "web/red_hello.html")
	} else {
		t, _ = template.ParseFiles("web/layout.html", "web/blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr:		"127.0.0.1:8000",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	http.HandleFunc("/process_context", processContext)
	http.HandleFunc("/process_form", processForm)
	http.HandleFunc("/form", form)
	http.HandleFunc("/process_layout", processLayout)
	server.ListenAndServe()
}