package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@/sample?charset=utf8&parseTime=True&loc=Local")
	// 接続に失敗したらエラーログを出して終了する
	if err != nil {
		log.Fatalf("DB connection failed %v", err)
	}
	db.LogMode(true)

	return db
}

func GetID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}
