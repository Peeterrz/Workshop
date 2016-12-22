package model

import (
	"app/shared/database"
)

type Bank struct {
	BKCD 		int16		`db:"bkcd"`
	NAME 		string 		`db:"name"`
	ENAME		string 		`db:"ename"`
}

func Banks() ([]Bank, error){
	
	var err error

	var result []Bank
	
	err = database.SQL.Select(&result, "select bkcd, name, ename from ZUTBLBK")

	return result, standardizeError(err)
}


func BanksByBankCode(bankCode int) (Bank, error){
	
	var err error

	result := Bank{}
	
	err = database.SQL.Get(&result, "select bkcd, name, ename from ZUTBLBK where bkcd = ?", bankCode)
	
	return result, standardizeError(err)
}

