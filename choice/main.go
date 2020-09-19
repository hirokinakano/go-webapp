package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/onlyPost", handleOnlyPost)
	http.ListenAndServe(":8080", nil)
}

// func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World from Go.")
// }

func handleOnlyPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			w.Write([]byte("POSTだけだよー"))
			return
	}

	w.Write([]byte("OK"))
}