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

func Substring(s string, start int, end int) string {
    start_str_idx := 0
    i := 0
    for j := range s {
        if i == start {
            start_str_idx = j
        }
        if i == end {
            return s[start_str_idx:j]
        }
        i++
    }
    return s[start_str_idx:]
}