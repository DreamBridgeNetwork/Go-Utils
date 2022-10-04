package jwtutils

import (
	"crypto/rsa"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateSignedRSAJWT - Generate a signed with RS512 jwt with claims
func GenerateSignedRSAJWT(claims *CustomClaims, privateKey *rsa.PrivateKey) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("jwtutils.GenerateSignedRSAJWT - Error signing jwt.")
		return "", err
	}

	return tokenString, nil
}

// ParseJWTWithClaims - Parse one jwt string, verify signature and return claims.
func ParseJWTWithClaims(jwtString string, publicKey interface{}) (*CustomClaims, error) {

	var claims CustomClaims

	_, err := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Println("jwtutils.ParseJWTWithClaims - Error parsing jwt.")
		return nil, err
	}

	return &claims, nil
}

// Bibliography
// https://jwt.io/introduction
// https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/?utm_source=google&utm_medium=cpc&utm_term=-g-security%20jwt&pm=true&utm_campaign=latam-pt-brazil-generic-authentication&utm_source=google&utm_campaign=latam_mult_bra_all_ciam-dev_dg-ao_auth0_search_google_text_kw_utm2&utm_medium=cpc&utm_term=security%20jwt-c&utm_id=aNK4z0000004GagGAE&gclid=CjwKCAjws--ZBhAXEiwAv-RNL1vSO56fdsTqzIF9Y3A8eGBgr8IdpXBqMW0pixhCUpP5watA3WppuxoC7AYQAvD_BwE
// https://blog.logrocket.com/jwt-authentication-go/
// https://www.bacancytechnology.com/blog/golang-jwt
// https://github.com/golang-jwt/jwt
