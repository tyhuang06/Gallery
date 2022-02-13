package controllers

import (
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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}
