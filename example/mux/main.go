package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Post struct {
	Title string `json:"title"`
	Tag string `json:"tag"`
	URL string `json:"url"`
}

func main(){
	r := mux.NewRouter()
	// /publicでpublic handler
	r.Handle("/public", public)

	// run server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe", nil)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
		Tag:   "Vue.js",
		URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)
})