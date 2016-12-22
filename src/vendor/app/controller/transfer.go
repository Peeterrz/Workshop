package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	
	"app/shared/view"
	"app/model"
	"app/utilities"
	
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func TransferGET(w http.ResponseWriter, r *http.Request) {
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	from_acc_no := params.ByName("acc_no")
	
	banks, err := model.Banks()
	if err != nil {
		log.Println(err)
		banks = []model.Bank{}
	}
	
	v := view.New(r)
	v.Name = "transfer/form"
	v.Vars["from_acc_no"] = from_acc_no
	v.Vars["banks"] = banks
	v.Render(w)
}

func Transfer_Verify(w http.ResponseWriter, r *http.Request) {
	accfromInput := r.FormValue("#from_acc_no")
	acctoInput := r.FormValue("to_acc_no")
	bankInput := r.FormValue("bank_code")
	trnamtInput := r.FormValue("tamt")
	
	trn_amt, err := strconv.ParseFloat(trnamtInput, 64)
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
	
	bankObj, err := model.BanksByBankCode(bankCode)
	if err != nil {
		panic(err)
	}  
	
	v := view.New(r)
	v.Name = "transfer/transfer_verify"
	v.Vars["from_acc_no"] = accountFromObj.CID
	v.Vars["from_acc_name"] = accountFromObj.ACCNAME
	v.Vars["to_acc_no"] = accountToObj.CID
	v.Vars["to_acc_name"] = accountToObj.ACCNAME
	v.Vars["bank_code"] = bankObj.BKCD
	v.Vars["bank_name"] = bankObj.NAME
	v.Vars["trn_amt"] = utilities.ThaiCurrencyFormat(trn_amt)
	v.Vars["fee_amt"] = utilities.ThaiCurrencyFormat(FeeCal(accountFromObj,accountToObj))
	v.Render(w)
}

func Transfer_Post(w http.ResponseWriter, r *http.Request) {
	accfromInput := r.FormValue("#from_acc_no")
	acctoInput := r.FormValue("#to_acc_no")
	bankInput := r.FormValue("bank_code")
	trnamtInput := r.FormValue("trn_amt")
	feeamtInput := r.FormValue("feeamt")
	
	trnamtNoFormat := strings.Replace(utilities.Substring(trnamtInput, 0, len([]rune(trnamtInput))-4), ",", "", -1)
	
	accfrom, err := strconv.Atoi(accfromInput)
	accto, err := strconv.Atoi(acctoInput)
	bankCode, err := strconv.Atoi(bankInput)
	
	trnamt, err := strconv.ParseFloat(trnamtNoFormat, 64)
	stramt := strconv.FormatFloat(trnamt, 'f', 2, 64)
	
	feeamt, err := strconv.ParseFloat(feeamtInput, 64)
	totalamt := trnamt+feeamt

	accountFromObj, err := model.AccountByAccountNo(accfrom)
	if err != nil {
		panic(err)
	}
	
	accountToObj, err := model.AccountByAccountNo(accto)
	if err != nil {
		panic(err)
	}
	
	bankObj, err := model.BanksByBankCode(bankCode)
	if err != nil {
		panic(err)
	}  
	
	err = Post(accountFromObj,accountToObj,totalamt)
	if err != nil {
		panic(err)
	}  
	
	v := view.New(r)
	v.Name = "transfer/transfer_post"
	v.Vars["from_acc_no"] = accountFromObj.CID
	v.Vars["to_acc_no"] = accountToObj.CID
	v.Vars["to_acc_name"] = accountToObj.ACCNAME
	v.Vars["bank_code"] = bankObj.BKCD
	v.Vars["bank_name"] = bankObj.NAME
	v.Vars["trn_amt"] = utilities.ThaiCurrencyFormat(trnamt)
	v.Vars["fee_amt"] = utilities.ThaiCurrencyFormat(feeamt)
	v.Render(w)
}

func FeeCal(accountFromObj model.Account ,accountToObj model.Account)(output float64) {
	Interregion := Interregion(accountFromObj , accountToObj)

	output=0.00
	if (Interregion) && (accountFromObj.TRNO>10) {
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

func Post(accountFromObj model.Account ,accountToObj model.Account,totalamt float64,)(err error) {
	FromAccountNewBal := accountFromObj.BAL-totalamt
	FromAccountNewBalAvl := accountFromObj.BALAVL-totalamt
	FromAccountTrno := accountFromObj.TRNO+1
	
	ToAccountNewBal := accountToObj.BAL+totalamt
	ToAccountNewBalAvl := accountToObj.BALAVL+totalamt
	ToAccountTrno := accountToObj.TRNO
	
	err = model.UpdateAccountAfterTransaction(accountFromObj.CID,FromAccountNewBal,FromAccountNewBalAvl,FromAccountTrno)
	if err != nil {
		panic(err)
	}  
	err = model.UpdateAccountAfterTransaction(accountToObj.CID,ToAccountNewBal,ToAccountNewBalAvl,ToAccountTrno)
	if err != nil {
		panic(err)
	} 
	
	return err
}