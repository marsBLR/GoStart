package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//var posts map[string]*Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	t.ExecuteTemplate(w, "write", nil)
}

//

func main() {
	fmt.Println("Listening on port :3001")
	//	posts = make(map[string]*Post, 0)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	//http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":3001", nil)
}
