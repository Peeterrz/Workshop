package utilities

import (
		"github.com/leekchan/accounting"
)

func PadLeft(str, pad string, lenght int) string {
	for {
		if len(str) >= lenght {
			return str[0:lenght]
		}
		str = pad + str
	}
}

func ThaiCurrencyFormat(balance float64) string {	
	ac := accounting.Accounting{Precision: 2}
    return ac.FormatMoney(balance)+" บาท"
}
