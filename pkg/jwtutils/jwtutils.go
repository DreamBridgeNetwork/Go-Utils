package jwtutils

import (
	"crypto/rsa"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

/*func GenerateUnsignedJWT(payload string) (string, error) {

	token := jwt.New(jwt.SigningMethodNone)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"

	token := jwt.New(jwt.SigningMethodNone)

	tokenString := token.Raw

	return tokenString, nil
}*/

func GenerateSignedRSAJWT(claims *CustomClaims, privateKey *rsa.PrivateKey) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("jwtutils.GenerateSignedRSAJWT - Error signing jwt.")
		return "", err
	}

	return tokenString, nil
}

/*
func VerifyJWTSignature(tokenString string, publicKey *rsa.PublicKey) (bool, error) {
	claims := &kitekey.KiteClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().UTC().Unix(),
		},
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return publicKey, nil
	})

	signingString, err := jwt.SigningString()

	if err != nil {
		log.Println("jwtutils.VerifyJWTSignature - Error getting signing string.")
		return false, err
	}

	err = jwt.Method.Verify(signingString, jwt.Signature, publicKey)
	if err != nil {
		log.Println("jwtutils.VerifyJWTSignature - Error verifying signature.")
		return false, err
	}

	return true, nil
}*/

// Bibliography
// https://jwt.io/introduction
// https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/?utm_source=google&utm_medium=cpc&utm_term=-g-security%20jwt&pm=true&utm_campaign=latam-pt-brazil-generic-authentication&utm_source=google&utm_campaign=latam_mult_bra_all_ciam-dev_dg-ao_auth0_search_google_text_kw_utm2&utm_medium=cpc&utm_term=security%20jwt-c&utm_id=aNK4z0000004GagGAE&gclid=CjwKCAjws--ZBhAXEiwAv-RNL1vSO56fdsTqzIF9Y3A8eGBgr8IdpXBqMW0pixhCUpP5watA3WppuxoC7AYQAvD_BwE
// https://blog.logrocket.com/jwt-authentication-go/
// https://www.bacancytechnology.com/blog/golang-jwt
// https://github.com/golang-jwt/jwt
