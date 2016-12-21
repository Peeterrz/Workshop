package controller

import (
	//"fmt"
	"log"
	"net/http"
	
	"app/shared/view"
	"app/model"
	
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// TransferGET displays the login page
func TransferGET(w http.ResponseWriter, r *http.Request) {
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	from_acc_no := params.ByName("acc_no")
	
	banks, err := model.GetBanks()
	if err != nil {
		log.Println(err)
		banks = []model.Bank{}
	}
	
	// Display the view
	v := view.New(r)
	v.Name = "transfer/form"
	v.Vars["from_acc_no"] = from_acc_no
	v.Vars["banks"] = banks
	
	// Refill any form fields
	//view.Repopulate([]string{"email"}, r.Form, v.Vars)
	v.Render(w)
}

func TransferPOST(w http.ResponseWriter, r *http.Request) {
	log.Println("*** TransferPOST")
	// Validate with required fields
	/*if validate, missingField := view.Validate(r, []string{"note"}); !validate {
		sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
		sess.Save(r, w)
		NotepadCreateGET(w, r)
		return
	}*/
	
	// Get form values
	from_acc_no := r.FormValue("#from_acc_no")
	to_acc_no := r.FormValue("to_acc_no")
	bank_code := r.FormValue("bank_code")
	tamt := r.FormValue("tamt")
	
	log.Println("*** from_acc_no = " + from_acc_no)
	log.Println("*** to_acc_no = " + to_acc_no)
	log.Println("*** bank_code = " + bank_code)
	log.Println("*** tamt = " + tamt)
	
	v := view.New(r)
	v.Name = "transfer/verify"
	v.Render(w)
}