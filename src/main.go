package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type entry struct {
	ID    int
	Title string
	Text  string
}

func getEntries() ([]entry, error) {
	rows, err := database.Query("SELECT id, Title, Text FROM tasks")
	if err != nil {
		return nil, err
	}
	var title string
	var text string
	var id int
	var entries = make([]entry, 0)
	for rows.Next() {
		rows.Scan(&id, &title, &text)
		entries = append(entries, entry{id, title, text})
	}
	return entries, nil
}

var database *sql.DB

func viewHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := getEntries()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "FAIL")
	}
	t, _ := template.ParseFiles("view.html")
	fmt.Print(entries)
	t.Execute(w, entries)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/edit/"):]

	// entry, ok := entries[title]

	// if !ok {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprint(w, title+" not found")
	// 	return
	// }

	// switch r.Method {
	// case "GET":
	// 	t, _ := template.ParseFiles("edit.html")
	// 	t.Execute(w, entry)
	// case "POST":
	// 	entries[title].Text = r.FormValue("text")
	// 	http.Redirect(w, r, "/edit/"+title, http.StatusSeeOther)
	// }
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprint(w, " Not found")
	// 	return
	// }

	// title := r.FormValue("title")

	// if title == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, "Empty title not allowed")
	// 	return
	// }
	// if _, ok := entries[title]; ok {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, title+" already exists")
	// 	return
	// }

	// entries[title] = &entry{Title: title, Text: r.Form.Get("title")}
	// http.Redirect(w, r, "/view", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/delete/"):]

	// if _, ok := entries[title]; !ok {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, title+" does not exists")
	// }

	// delete(entries, title)
	// http.Redirect(w, r, "/view", http.StatusSeeOther)
}

func main() {
	database, _ = sql.Open("sqlite3", "db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, Title TEXT, Text TEXT)")
	statement.Exec()

	// insertStatement, _ := database.Prepare("INSERT INTO tasks (Title, Text) VALUES (?, ?)")
	// insertStatement.Exec("WIFE", "get a wife")

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/delete/", deleteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
