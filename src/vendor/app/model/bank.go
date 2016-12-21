package model

import (
	"app/shared/database"
)

type Bank struct {
	BKCD 		int16		`db:"bkcd"`
	NAME 		string 		`db:"name"`
	ENAME		string 		`db:"ename"`
}

func GetBanks() ([]Bank, error){
	
	var err error

	var result []Bank

	switch database.ReadConfig().Type {
		case database.TypeMySQL:
			err = database.SQL.Select(&result, "select bkcd, name, ename from ZUTBLBK")
		case database.TypeMongoDB:
		case database.TypeBolt:
		default:
			err = ErrCode
	}

	return result, standardizeError(err)
}


func GetBanksCode(bankcd int) (Bank, error){
	
	var err error

	result := Bank{}

	switch database.ReadConfig().Type {
		case database.TypeMySQL:
			err = database.SQL.Get(&result, "select bkcd, name, ename from ZUTBLBK where bkcd = ?", bankcd)
		case database.TypeMongoDB:
		case database.TypeBolt:
		default:
			err = ErrCode
	}

	return result, standardizeError(err)
}

