package handler

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateToken(password string, expireDuration time.Duration) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write(HashSalt(password))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	var signKey []byte
	jwtReg := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtReg)
	return token.SignedString(signKey)
}

func HashSalt(password string) []byte {
	pwd := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, pwd)
	//fmt.Println(err) // nil means it is a match
	return hashedPassword
}
