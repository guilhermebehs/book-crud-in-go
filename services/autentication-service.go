package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type AuthenticationService struct{}

func (as AuthenticationService) Authenticate(username string, password string) (string, error) {

	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		panic("Error loading config file: " + dotenvErr.Error())
	}
	authenticatedUser := os.Getenv("AUTHENTICATED_USER")
	authenticatedPass := os.Getenv("AUTHENTICATED_PASS")

	if username == authenticatedUser && password == authenticatedPass {
		accessToken := jwt.New(jwt.SigningMethodHS256)
		claims := accessToken.Claims.(jwt.MapClaims)
		claims["user_id"] = time.Now().Unix()
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		token, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			return "", errors.New("internal error")
		}
		return token, nil
	}
	return "", errors.New("Unauthorized")
}

func (as AuthenticationService) Validate(token string) bool {

	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		panic("Error loading config file: " + dotenvErr.Error())
	}
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	return err == nil

}

func CreateAuthenticationService() AuthenticationService {
	authenticationService := AuthenticationService{}

	return authenticationService
}
