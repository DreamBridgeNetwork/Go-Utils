package stringutils

import (
	"encoding/base64"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// GeneratePasswordHash - Returns the hash of a pessaword.
func GeneratePasswordHash(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)

	if err != nil {
		log.Println("stringutils.GeneratePasswordHash - Error generating password hash.")
		return "", nil
	}
	return string(bytes), nil
}

// CheckStringHash - Confirms if some string matchs the hash
func CheckStringHash(str, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))

	if err != nil {
		log.Println("stringutils.checkStringHash - Error validating string hash.")
		return false, err
	}

	return true, nil
}

// SeparaNomeSobrenome - Separa uma string de nome completo em nome e sobrenome. OBS: Não funciona para nome composto.
func SeparaNomeSobrenome(nomeCompleto string) (string, string, error) {
	var nome string
	var sobrenome string

	i := strings.Index(nomeCompleto, " ")

	if i < 0 {
		log.Println("stringutils.SeparaNomeSobrenome - O nome não contém sobrenome.")
		nome = nomeCompleto
		return nome, sobrenome, errors.New("stringutils.SeparaNomeSobrenome - O nome não contém sobrenome")
	} else if i == 0 {
		log.Println("stringutils.SeparaNomeSobrenome - Existe um espeço na primeira letra? A string está vazia?")
		nome = nomeCompleto
		return nome, sobrenome, errors.New("stringutils.SeparaNomeSobrenome - Existe um espeço na primeira letra? A string está vazia?")
	} else {
		nome = nomeCompleto[0:i]
		sobrenome = nomeCompleto[(i + 1):]
	}

	return nome, sobrenome, nil
}

func GeraStringAleatoria(tamanho int) string {
	rand.Seed(time.Now().UnixNano())

	//const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, tamanho)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// MapToString - Generate one string with key and value from a map[string][]string
func MapToString(mapVariable *map[string][]string) string {

	var mapString string

	if mapVariable != nil {
		// Loop over trailer names
		for name, values := range *mapVariable {
			mapString += name + ": " + VectorStringToStringLine(values)
			mapString += "\n"
		}
	}
	return mapString
}

// VectorStringToStringLine - Create one string with one line of all strings from a string vector separated with ","
func VectorStringToStringLine(vectorString []string) string {
	var stringLine string

	// Loop over all values for the name.
	for index, value := range vectorString {

		stringLine += value

		if index < (len(vectorString) - 1) {
			stringLine += ", "
		}
	}

	return stringLine
}

// Base64Decode - Decode from base64
func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("stringutils.Base64Decode - Error decoding from base64.")
		return "", err
	}
	return string(data), nil
}

// Base64Encode - Encode string to base64
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64URLDecode - Decode from base64 url
func Base64URLDecode(str string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		log.Println("stringutils.Base64URLDecode - Error decoding from base64 url.")
		return "", err
	}
	return string(data), nil
}

// Base64URLEncode - Encode string to base64URL
func Base64URLEncode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}
