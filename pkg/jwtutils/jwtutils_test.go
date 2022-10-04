package jwtutils

import (
	"log"
	"testing"

	"github.com/DreamBridgeNetwork/Go-Utils/pkg/rsakey"
)

func TestJWTGenerationValidation(t *testing.T) {
	log.Println("TestJWTGenerationValidation")

	key, err := rsakey.GenerateRSAKeyPair(rsakey.RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if key == nil {
		t.Error("Generated Key is null.")
		return
	}

	testeJson := `{
        "country" : "China",
        "city" : "Wuhan",
		"people" : [
			{
				"name" : "Name 1",
				"email" : "email1@email.com"
			},
			{
				"name" : "Name 2",
				"email" : "email2@email.com"
			}
		]
    }`

	claims, err := NewClaims(testeJson)
	if err != nil {
		t.Error("Error generating new claim: ", err)
		return
	}

	tokenString, err := GenerateSignedRSAJWT(claims, key)
	if err != nil {
		t.Error("Error generating signed jwt: ", err)
		return
	}

	log.Println("JWT: ", tokenString)

	key.Public()

	tokenClaims, err := ParseJWTWithClaims(tokenString, key.Public())
	if err != nil {
		t.Error("Error parsing jwt: ", err)
		return
	}

	claimsJson, err := tokenClaims.GetString()
	if err != nil {
		t.Error("Error converting clains to json string: ", err)
		return
	}

	log.Println("claims: ", claimsJson)
}
