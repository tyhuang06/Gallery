package controllers

import (
	"fmt"
	"net/http"

	"github.com/tyhuang06/Gallery/views"
)

// Create a new Users controller
// Panic if templates are not parsed corretly, should only be used for initial setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// Render the form for creating a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Process the signup form when submitted
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "temporary response")
}
