package handler

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/*type Token struct {
	token string
	expiresAt time.Time
	issuedAt time.Time
	exist bool
}*/

type Claims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected auth method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Login, nil
	}
	return "", Claims{}.Valid()
}

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
