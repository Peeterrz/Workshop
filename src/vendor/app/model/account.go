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
}

func AccountByAccountNo(accountNo int) (Account, error){
	
	var err error

	result := Account{}

	switch database.ReadConfig().Type {
		case database.TypeMySQL:
			err = database.SQL.Get(&result, "select cid, accname, bal, balavl, province from ACN where cid = ?", accountNo)
		case database.TypeMongoDB:
		case database.TypeBolt:
		default:
			err = ErrCode
	}

	return result, standardizeError(err)
}




func SetAccount(content string, userID string, noteID string) error {
	var err error


	switch database.ReadConfig().Type {
	case database.TypeMySQL:
		_, err = database.SQL.Exec("UPDATE note SET content=? WHERE id = ? AND user_id = ? LIMIT 1", content, noteID, userID)

	}
	return standardizeError(err)
}