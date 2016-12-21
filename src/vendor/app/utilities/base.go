package utilities


func PadLeft(str, pad string, lenght int) string {
	for {
		if len(str) >= lenght {
			return str[0:lenght]
		}
		str = pad + str
	}
}
