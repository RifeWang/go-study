package main

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	hmacSampleSecret = []byte("secret") // 密钥
)

func main() {
	token, err := jwtSign("admin")
	fmt.Println(token, err)

	c, err := jwtCheck(token)
	fmt.Println(" =========== ", c, err, c["username"])
}

func jwtSign(username string) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().AddDate(0, 0, 7),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString(hmacSampleSecret)
	return
}

func jwtCheck(tokenString string) (claims map[string]interface{}, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claim to mapclaim")
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return
}

/*
type Token struct {
    Raw       string                 // The raw token.  Populated when you Parse a token
    Method    SigningMethod          // The signing method used or to be used
    Header    map[string]interface{} // The first segment of the token
    Claims    Claims                 // The second segment of the token
    Signature string                 // The third segment of the token.  Populated when you Parse a token
    Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}
*/
