// testing: https://www.loginradius.com/blog/engineering/google-authentication-with-golang-and-goth/
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"log"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = false  // Set to true when serving over https

	gothic.Store = store

	goth.UseProviders(
		google.New("our-google-client-id", "our-google-client-secret", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			_, err := fmt.Fprintln(res, err)
			if err != nil {
				return
			}
			return
		}
		t, _ := template.ParseFiles("templates/success.html")
		err = t.Execute(res, user)
		if err != nil {
			return
		}
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		err := t.Execute(res, false)
		if err != nil {
			return
		}
	})
	log.Println("listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}
