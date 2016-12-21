package controller

import (
	//"fmt"
	//"log"
	"net/http"
	
	"app/shared/view"
)

// LoginGET displays the login page
func TransferGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	//sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "transfer/form"
	//v.Vars["token"] = csrfbanana.Token(w, r, sess)
	// Refill any form fields
	//view.Repopulate([]string{"email"}, r.Form, v.Vars)
	v.Render(w)
}