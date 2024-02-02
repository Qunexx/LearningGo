package main

import (
	// "fmt"
	"fmt"
	"net/http"

	// "time"
	"html/template"
)

func main() {

	type Item struct {
		ID   string
		Name string
	}

	var items = make(map[string]Item)

	type TemplateData struct {
		Message string
	}

	// Загрузка шаблонов
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Unsucces", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			name := r.FormValue("name")
			items[id] = Item{ID: id, Name: name}

			templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Success!"})
			return
		}

		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if item, ok := items[id]; ok {
			templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Item: " + item.ID + ", Name: " + item.Name})
		} else {
			templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Unsuccess"})
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Failed to parse form", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			newName := r.FormValue("name")
			if _, ok := items[id]; ok {
				items[id] = Item{ID: id, Name: newName}
				templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Success"})
			} else {
				templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Unsuccess"})
			}
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Failed to parse form", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			if _, ok := items[id]; ok {
				delete(items, id)
				templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Success"})
			} else {
				templates.ExecuteTemplate(w, "index.html", TemplateData{Message: "Unsuccess"})
			}
		}
	})
	fmt.Println("Сервер запущен по адресу: localhost:8080")
	http.ListenAndServe(":8080", nil)

}
