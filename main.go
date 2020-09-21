package main

import (
	"fmt"
	"log"
	"net/http"

	"./utils"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User User構造体
type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users", findAllUsers)
	router.HandleFunc("/users/{id}", findByID)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	var userList []User
	db.Find(&userList)

	// 共通化した処理を使う
	utils.RespondWithJSON(w, http.StatusOK, userList)
}

func findByID(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetID(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter")
		return
	}

	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	var user User
	db.Where("id = ?", id).Find(&user)

	// 共通化した処理を使う
	utils.RespondWithJSON(w, http.StatusOK, user)
}