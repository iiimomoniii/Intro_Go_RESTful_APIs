package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	//Case home
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, %q", html.EscapeString(r.URL.Path))
	})
	//Request => http://localhost:8080/home
	//Response => Home, "/home"

	//Case profile
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profile, %q", html.EscapeString(r.URL.Path))
	})
	//Request => http://localhost:8080/profile
	//Response => Profile, "/profile"

	//Case login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})
	//Request http://localhost:8080/login?username=admin&password=1234
	//Response Login, admin, 1234

	//Case empty route
	//Request => http://localhost:8080/about
	//404 page not found

	//logging
	log.Fatal(http.ListenAndServe(":8080", nil))
}
