package token

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/nik/JWTDemo/model"
)

//secretKey var
var secretKey = "mysecret"

//generates token on the basis of claims passed
func GenerateToken(login *model.Login) (string,error) {
	//instantiate a new token with claims that include email and issuer
	token:= jwt.NewWithClaims(jwt.SigningMethodHS384,jwt.MapClaims{
		"email":login.Email,
		"username": login.UserID,
		"iss":"Server",
	})

	//generate a token string by signing the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

//creates the jwt middleware that carries out the validation of jwt token in the incoming request
//in case of successful validation the request is enriched with the user details
//in other case panic is raised which is delegated to the jwtmiddleware
func GetJWTMiddleware() (*jwtmiddleware.JWTMiddleware){
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecret"), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS384,
	})

	return jwtMiddleware
}