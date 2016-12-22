package controller

import (
	"net/http"
	"log"
	"strconv"
	
	"app/model"
	"app/utilities"
	"app/shared/view"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func AccountOverviewGET(w http.ResponseWriter, r *http.Request) {
	
	log.Println("Account Overview GET : Start")	
	
	v := view.New(r)
	v.Name = "account/overview"
	
	defer func() {
			if err := recover(); err != nil {
				errorObj := err.(error)
				v.Vars["errorMessage"] = errorObj.Error()
			}
			v.Render(w)			
			log.Println("Account Overview GET : End")
		}()
	
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	accInput := params.ByName("account_no_input")
	accNo, err := strconv.Atoi(accInput)
	accountObj, err := model.AccountByAccountNo(accNo)
	if err != nil {
		panic(err)
	} 	
	
	v.Vars["acc_no"] = utilities.PadLeft(accInput, "0", 10)
	v.Vars["acc_name"] = accountObj.ACCNAME
	v.Vars["acc_bal"] = utilities.ThaiCurrencyFormat(accountObj.BALAVL)
}

