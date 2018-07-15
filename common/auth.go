package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "/keys/app.rsa"
	pubKeyPath  = "/keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)

	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}

// GenerateJWT token
func GenerateJWT(name, role string) (string, error) {
	// create a signer for rsa 256
	token := jwt.New(jwt.SigningMethodRS256)
	// set claims for JWT token
	claims := make(jwt.MapClaims)

	//Claims configuration
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["userInfo"] = struct {
		Name string
		Role string
	}{name, role}

	//Assigning claims to token claims
	token.Claims = claims

	//Generating token using private rsa key
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}
