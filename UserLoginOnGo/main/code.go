package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("trgtrgrtgsecretkey"))

var users = map[string]string{
	"admin": "admin",
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Только POST метод разрешен", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "session")
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Error(w, "Вы уже аутентифицированы", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	expectedPassword, ok := users[username]

	if !ok || expectedPassword != password {
		http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
		return
	}

	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(r, w)

	w.Write([]byte("Вы успешно аутентифицированы"))
}

// Обработчик для выхода из системы
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	session.Values["authenticated"] = false
	session.Values["username"] = nil
	session.Save(r, w)

	w.Write([]byte("Вы успешно вышли из системы"))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Только POST метод разрешен", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if _, exists := users[username]; exists {
		http.Error(w, "Пользователь уже существует", http.StatusBadRequest)
		return
	}

	users[username] = password
	fmt.Fprintf(w, "Пользователь %s успешно зарегистрирован", username)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
