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

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create processes the signup form when a user submits it. Ultimately creates a new user account
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
