package controllers

import (
	"bizzy_track/views"
	"fmt"
	"net/http"
)

/* NewUsers is used to create a new Users controller
This function will panic if the templates are not parsed correctly
Should only be used during intiial set up */

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

//New renders form where a user can create a new User account
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Create processes the signup form when a user submits it. Ultimately creates a new user account
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostFormValue("email"))
	fmt.Fprintln(w, r.PostForm["password"])
	fmt.Fprintln(w, r.PostFormValue("password"))

}
