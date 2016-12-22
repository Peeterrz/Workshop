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

func UpdateAccount(content string, userID string, noteID string) error {
	var err error

	_, err = database.SQL.Exec("UPDATE note SET content=? WHERE id = ? AND user_id = ? LIMIT 1", content, noteID, userID)

	return standardizeError(err)
}