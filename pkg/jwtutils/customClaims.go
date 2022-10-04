package jwtutils

import (
	"encoding/json"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var DefaulMinutesToExpire = 15

type CustomClaims struct {
	jwt.StandardClaims
	Private *map[string]interface{} `json:"private,omitempty"`
	Public  *map[string]interface{} `json:"public,omitempty"`
}

func (claims *CustomClaims) SetIssuedAt() {
	claims.IssuedAt = time.Now().UTC().Unix()
}

func (claims *CustomClaims) SetExpiresAt(minutesToExpire int) {
	claims.ExpiresAt = time.Now().Local().Add(time.Minute * time.Duration(minutesToExpire)).UTC().Unix()
}

func (claims *CustomClaims) SetPrivate(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), &claims.Private)

	if err != nil {
		log.Println("jwtutils.CustomClaims.SetPrivate - Error converting json to map.")
		return err
	}

	return nil
}

func NewClaims(privateJsonString string) (*CustomClaims, error) {
	var newClaim CustomClaims

	newClaim.Id = uuid.New().String()
	newClaim.SetIssuedAt()
	newClaim.SetExpiresAt(DefaulMinutesToExpire)
	err := newClaim.SetPrivate(privateJsonString)

	if err != nil {
		log.Println("jwtutils.CreateNewClaim - Error converting json to map.")
		return nil, err
	}

	return &newClaim, nil
}

func (claims *CustomClaims) GetString() (string, error) {

	claimsJsonStr, err := json.Marshal(claims)

	if err != nil {
		log.Println("jwtutils.CustomClaims.GetString - Error converting map to json.")
		return "", err
	}

	return string(claimsJsonStr), nil
}
