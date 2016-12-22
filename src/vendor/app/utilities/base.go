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
	formatter := accounting.Accounting{Precision: 2}
    return formatter.FormatMoney(balance)+" บาท"
}
