package generate

import "strings"

func UpperFirstSymbol(login string) string {
	firstsymbol := login[:1]
	strsToUp := strings.ToUpper(firstsymbol)
	login = strsToUp + login[1:]
	return login
}
