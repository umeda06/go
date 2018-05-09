package main

import (
	_ "fmt"
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	T = "たいち"
	A = "あゆみ"
)

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)

// フォーム入力データ
type Post struct {
	Author  string
	Message string
	Posted  time.Time
}

// テンプレート埋め込みデータ
type templateParams struct {
	Notice string
	Name string
	Message string
	Posts1 []Post // メッセージ検索結果
	Access1 []Post // アクセス日時検索結果
	Check1 string
	Posts2 []Post
	Access2 []Post
	Check2 string
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	params := templateParams{}

	// GET要求時
	if r.Method == "GET" {
		indexTemplate.Execute(w, params)
		return
	}

	// POST要求時
	post := Post{
		Author:  r.FormValue("name"),
		Message: r.FormValue("message"),
		Posted: jst(time.Now()),
	}
	params.Name = post.Author

	// 名前が正しくない場合
	if post.Author != T && post.Author != A {
		indexTemplate.Execute(w, params)
		return
	}

	ctx := appengine.NewContext(r)

	// アクセス日時更新
	if post.Author == T {
		key := datastore.NewIncompleteKey(ctx, "Access1", nil)
		datastore.Put(ctx, key, &post)
		// メッセージ更新
		if post.Message != "" {
			key := datastore.NewIncompleteKey(ctx, "Post1", nil)
			if _, err := datastore.Put(ctx, key, &post); err != nil {
				log.Errorf(ctx, "datastore.Put: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				params.Notice = "Couldn't add new post. Try again?"
				params.Message = post.Message
				indexTemplate.Execute(w, params)
				return
			}
			params.Posts1 = []Post{post}
		}
	}
	if post.Author == A {
		key := datastore.NewIncompleteKey(ctx, "Access2", nil)
		datastore.Put(ctx, key, &post)
		if post.Message != "" {
			key := datastore.NewIncompleteKey(ctx, "Post2", nil)
			if _, err := datastore.Put(ctx, key, &post); err != nil {
				log.Errorf(ctx, "datastore.Put: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				params.Notice = "Couldn't add new post. Try again?"
				params.Message = post.Message
				indexTemplate.Execute(w, params)
				return
			}
			params.Posts2 = []Post{post}
		}
	}

	// メッセージ検索
	if post.Message == "" || post.Author == T {
		q2 := datastore.NewQuery("Post2").Order("-Posted").Limit(1)
		if _, err := q2.GetAll(ctx, &params.Posts2); err != nil {
			log.Errorf(ctx, "Getting posts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			params.Notice = "Couldn't get latest posts. Refresh?"
			indexTemplate.Execute(w, params)
			return
		}
	}
	if post.Message == "" || post.Author == A {
		q1 := datastore.NewQuery("Post1").Order("-Posted").Limit(1)
		if _, err := q1.GetAll(ctx, &params.Posts1); err != nil {
			log.Errorf(ctx, "Getting posts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			params.Notice = "Couldn't get latest posts. Refresh?"
			indexTemplate.Execute(w, params)
			return
		}
	}

	// アクセス日時検索
	if post.Author == T {
		q4 := datastore.NewQuery("Access2").Order("-Posted").Limit(1)
		q4.GetAll(ctx, &params.Access2)
	}
	if post.Author == A {
		q3 := datastore.NewQuery("Access1").Order("-Posted").Limit(1)
		q3.GetAll(ctx, &params.Access1)
	}

	// テンプレート埋め込みデータ補正
	params.Name = ""
	if len(params.Posts1) == 1 {
		params.Posts1[0].Posted = jst(params.Posts1[0].Posted)
	}
	if len(params.Access1) == 1 {
		params.Access1[0].Posted = jst(params.Access1[0].Posted)
	}
	if len(params.Posts2) == 1 {
		params.Posts2[0].Posted = jst(params.Posts2[0].Posted)
	}
	if len(params.Access2) == 1 {
		params.Access2[0].Posted = jst(params.Access2[0].Posted)
	}
	if post.Author == T && len(params.Access2) == 1 && len(params.Posts1) == 1 && params.Access2[0].Posted.After(params.Posts1[0].Posted) {
		params.Check1 = "Checked"
	}
	if post.Author == A && len(params.Access1) == 1 && len(params.Posts2) == 1 && params.Access1[0].Posted.After(params.Posts2[0].Posted) {
		params.Check2 = "Checked"
	}

	indexTemplate.Execute(w, params)
	return
}

func jst(now time.Time) time.Time {
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	return nowJST
}
