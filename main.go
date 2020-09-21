package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests()  {
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", returnAllArticles)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}

type Article struct {
	Id int `json:"id"`
	Title string `json:Title`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{}
	for i := 0; i < 10; i++ {
		title := "Hello_%d"
		articles = append(
			articles,
			Article{Title: fmt.Sprintf(title, i), Desc: "Article Description", Content: "Article Content"})
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, vars)
	fmt.Fprintf(w, "Key: " + key + "\n")
	//article := Article{Id: key, Title: "タイトル" }
	//json.NewEncoder(w).Encode(article)

}
