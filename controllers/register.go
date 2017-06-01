package controllers

import (
	"net/http"

	"github.com/gustavocd/confirmation-account/models"
	"github.com/gustavocd/confirmation-account/templates"
	"github.com/julienschmidt/httprouter"
)

// Index renders the sign up form view.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates.Render(w, "register.html", nil)
}

// Create saves a new user.
func Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := &models.User{
		Name:     r.FormValue("name"),
		LastName: r.FormValue("lastName"),
		Email:    r.FormValue("email"),
	}

	lastInsertID, err := u.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	at := &models.ActivationToken{}
	err = at.GenerateToken(lastInsertID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := &models.Mail{}
	if err = m.SendTo(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
