package auth

import "github.com/golang-jwt/jwt"

var secretKey = []byte("secret")

func GenerateJwt(username string, id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["id"] = id

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyJwtToken(token string) interface{} {
	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil
	}
	claims := data.Claims.(jwt.MapClaims)
	return claims["id"]

}
