package controller

import (
	
	"net/http"
	_ "log"
	"strconv"
	"app/model"
	"app/shared/view"
	_ "github.com/gorilla/context"
	_ "github.com/julienschmidt/httprouter"


)



func Transfer_Verify(w http.ResponseWriter, r *http.Request) {

	
	
	
	


	
	// Get the note id
	//var params httprouter.Params
	//params = context.Get(r, "params").(httprouter.Params)
	//accfromInput := params.ByName("acc_from")
	//acctoInput := params.ByName("acc_to")
	//trnamtInput := params.ByName("trn_amt")
	//bankInput := params.ByName("bank_code")
	
	
	accfromInput := "1234567890"
	acctoInput := "1234567891"
	trnamtInput := "100"
	bankInput := "0"
	
	
	accfrom, err := strconv.Atoi(accfromInput)
	accto, err := strconv.Atoi(acctoInput)
	bankCode, err := strconv.Atoi(bankInput)


	accountFromObj, err := model.AccountByAccountNo(accfrom)
	if err != nil {
		panic(err)
	}
	
	
	accountToObj, err := model.AccountByAccountNo(accto)
	if err != nil {
		panic(err)
	}
	
	bankObj, err := model.GetBanksCode(bankCode)
	if err != nil {
		panic(err)
	}  
	
	

		
	// Display the view
	v := view.New(r)
	v.Name = "transfer/transfer_verify"
	v.Vars["from_acc_no"] = accountFromObj.CID
	v.Vars["from_acc_name"] = accountFromObj.ACCNAME
	v.Vars["to_acc_no"] = accountToObj.CID
	v.Vars["to_acc_name"] = accountToObj.ACCNAME
	v.Vars["bank_code"] = bankObj.BKCD
	v.Vars["bank_name"] = bankObj.NAME
	v.Vars["trn_amt"] = trnamtInput
	v.Vars["fee_amt"] = FeeCal(accountFromObj,accountToObj)



	v.Render(w)
	
}


func FeeCal(accountFromObj model.Account ,accountToObj model.Account)(output float64) {
	
	Interregion := Interregion(accountFromObj , accountFromObj)
	
	
	output=0.00
	if Interregion {
			output=10.00
	}

	return output
}

func Interregion(accountFromObj model.Account ,accountToObj model.Account)(output bool) {
	
	output=false
	
	if (accountFromObj.PROVINCE!=accountToObj.PROVINCE) {
		output=true
	}

	return output
	
}

