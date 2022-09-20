package handler

import "regexp"

func CheckPasswordCase(password string) bool {
	var CheckPassword bool

	var IsNumber = regexp.MustCompile(`[0-9]`).MatchString(password)
	var IsSymbol = regexp.MustCompile(`[^0-9A-Za-z_]`).MatchString(password)
	var IsLetter = regexp.MustCompile(`[a-z]`).MatchString(password)
	var IsUpper = regexp.MustCompile(`[A-Z]`).MatchString(password)

	if len(password) > 8 /*&& checkLetterCase == true*/ && IsNumber == true && IsSymbol == true && IsLetter == true && IsUpper == true {
		CheckPassword = true
	} else {
		CheckPassword = false
	}

	return CheckPassword
}

func CheckLoginCase(login string) bool {
	//var token string = "Admin12345"
	var checkLogin bool

	var IsLetter = regexp.MustCompile(`[a-z]`).MatchString(login)
	var IsUpper = regexp.MustCompile(`[A-Z]`).MatchString(login)
	var IsNumber = regexp.MustCompile(`[0-9]`).MatchString(login)
	/*if login == token {
		checkLogin = true
		return checkLogin
	}*/
	if len(login) > 8 /*&& login != token*/ && IsLetter == true && IsNumber == true && IsUpper == true {
		checkLogin = true
	} else {
		checkLogin = false
	}

	return checkLogin
}
