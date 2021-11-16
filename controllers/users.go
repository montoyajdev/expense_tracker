package controllers

import (
	"bizzy_track/views"
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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}
