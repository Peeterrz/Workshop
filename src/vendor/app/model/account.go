package model

import (
	"app/shared/database"
)

type Account struct {
	CID			 		int64		`db:"cid"`
	ACCNAME 			string 		`db:"accname"`
	BAL					float64 	`db:"bal"`
	BALAVL 				float64 	`db:"balavl"`
	PROVINCE 			int64 		`db:"province"`
	TRNO 				int64 		`db:"trno"`
}

func AccountByAccountNo(accountNo int) (Account, error){
	
	var err error

	result := Account{}
	
	err = database.SQL.Get(&result, "select cid, accname, bal, balavl, province,trno from ACN where cid = ?", accountNo)

	return result, standardizeError(err)
}

func UpdateAccountAfterTransaction(accountNo int64, bal float64, balavl float64, trno int64) error {
	var err error

	_,err = database.SQL.Exec("update ACN set bal=?,balavl=?,trno=? WHERE cid=?", bal, balavl, trno,accountNo)

	return standardizeError(err)
}