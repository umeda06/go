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

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)


type Post struct {
	Author  string
	Message string
	Posted  time.Time
}

type templateParams struct {
	Notice string
	Name string
	Message string
	Posts1 []Post
	Posts2 []Post
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

	if r.Method == "GET" {
		indexTemplate.Execute(w, params)
		return
	}

	post := Post{
		Author:  r.FormValue("name"),
		Message: r.FormValue("message"),
		Posted:  jst(time.Now()),
	}
	params.Name = post.Author

	if post.Author != "たいち225" && post.Author != "あゆみ117" {
		indexTemplate.Execute(w, params)
		return
	}

	ctx := appengine.NewContext(r)
	q1 := datastore.NewQuery("Post1").Order("-Posted").Limit(1)
	if _, err := q1.GetAll(ctx, &params.Posts1); err != nil {
		log.Errorf(ctx, "Getting posts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't get latest posts. Refresh?"
		indexTemplate.Execute(w, params)
		return
	}
	q2 := datastore.NewQuery("Post2").Order("-Posted").Limit(1)
	if _, err := q2.GetAll(ctx, &params.Posts2); err != nil {
		log.Errorf(ctx, "Getting posts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't get latest posts. Refresh?"
		indexTemplate.Execute(w, params)
		return
	}

	for i := 0; i < len(params.Posts1); i++ {
		params.Posts1[i].Posted = jst(params.Posts1[i].Posted)
	}
	for i := 0; i < len(params.Posts2); i++ {
		params.Posts2[i].Posted = jst(params.Posts2[i].Posted)
	}

	if post.Message == "" {
		indexTemplate.Execute(w, params)
		return
	}

	if post.Author == "たいち225" {
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
		indexTemplate.Execute(w, params)
		return
	}
	if post.Author == "あゆみ117" {
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
		indexTemplate.Execute(w, params)
		return
	}
}

func jst(now time.Time) time.Time {
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	return nowJST
}
