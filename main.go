package main

import "net/http"

// func main() {
// 	println("Hello, world.")
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1>Full Cycle Esquenta</h1>"))
	})
	http.ListenAndServe(":8080", nil)
}
