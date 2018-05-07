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
	Access1 []Post
	Posts2 []Post
	Access2 []Post
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
		Posted: jst(time.Now()),
	}
	params.Name = post.Author

	if post.Author != "たいち" && post.Author != "あゆみ" {
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
	if post.Author == "たいち" {
		q4 := datastore.NewQuery("Access2").Order("-Posted").Limit(1)
		q4.GetAll(ctx, &params.Access2)
	}
	if post.Author == "あゆみ" {
		q3 := datastore.NewQuery("Access1").Order("-Posted").Limit(1)
		q3.GetAll(ctx, &params.Access1)
	}

	for i := 0; i < len(params.Posts1); i++ {
		params.Posts1[i].Posted = jst(params.Posts1[i].Posted)
	}
	for i := 0; i < len(params.Posts2); i++ {
		params.Posts2[i].Posted = jst(params.Posts2[i].Posted)
	}
	for i := 0; i < len(params.Access1); i++ {
		params.Access1[i].Posted = jst(params.Access1[i].Posted)
	}
	for i := 0; i < len(params.Access2); i++ {
		params.Access2[i].Posted = jst(params.Access2[i].Posted)
	}

	if post.Message == "" {
		if post.Author == "たいち" {
			key := datastore.NewIncompleteKey(ctx, "Access1", nil)
			datastore.Put(ctx, key, &post)
		}
		if post.Author == "あゆみ" {
			key := datastore.NewIncompleteKey(ctx, "Access2", nil)
			datastore.Put(ctx, key, &post)
		}
		params.Name = ""
		indexTemplate.Execute(w, params)
		return
	}

	if post.Author == "たいち" {
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
		params.Name = ""
		indexTemplate.Execute(w, params)
		return
	}
	if post.Author == "あゆみ" {
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
		params.Name = ""
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
