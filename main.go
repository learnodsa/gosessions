package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/learnodsa/gosessions/pages"
	"github.com/learnodsa/gosessions/sessions"
	_ "github.com/learnodsa/gosessions/sessions/provider"
)

var globalSessions *sessions.Manager

func init() {
	globalSessions, _ = sessions.NewManager("memory", "gosessionid", 60)
	go globalSessions.GC()
}

func handlePages(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	if sess.Get("username") != nil {
		fmt.Fprintf(w, "%s", pages.Home)
	} else {
		fmt.Fprintf(w, "%s", pages.Login)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if len(r.Form["username"]) > 0 {
		sess.Set("username", r.Form["username"])
		fmt.Fprintf(w, "%s", pages.Home)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionDestroy(w, r)
	fmt.Fprintf(w, "%s", pages.Login)
}

func main() {
	//Handling all the url requests.
	http.HandleFunc("/", handlePages)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
