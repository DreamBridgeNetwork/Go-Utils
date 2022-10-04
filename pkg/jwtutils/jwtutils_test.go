package jwtutils

import (
	"log"
	"testing"

	"github.com/DreamBridgeNetwork/Go-Utils/pkg/rsakey"
)

/*func TestGenerateUnsignedJWT(t *testing.T) {
	log.Println("TestGenerateUnsignedJWT")

	tokenString, err := GenerateUnsignedJWT()
	if err != nil {
		t.Error("Error generating unsigned jwt: ", err)
		return
	}

	log.Println("JWT: ", tokenString)
}*/

func TestGenerateSignedRSAJWT(t *testing.T) {
	log.Println("TestGenerateSignedRSAJWT")

	key, err := rsakey.GenerateRSAKeyPair(rsakey.RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if key == nil {
		t.Error("Error generating RSA Key.")
		return
	}

	coronaVirusJSON := `{
        "name" : "covid-11",
        "country" : "China",
        "city" : "Wuhan",
        "reason" : "Non vedge Food",
		"place" : {
			"name" : "Nome 1",
			"pais" : "Brasil",
			"cidade" : "São Paulo"
		}
    }`

	claims, err := NewClaims(coronaVirusJSON)
	if err != nil {
		t.Error("Error generating new claim: ", err)
		return
	}

	tokenString, err := GenerateSignedRSAJWT(claims, key)
	if err != nil {
		t.Error("Error generating unsigned jwt: ", err)
		return
	}

	log.Println("JWT: ", tokenString)
}
