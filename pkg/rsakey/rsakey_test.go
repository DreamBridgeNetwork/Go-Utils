package rsakey

import (
	"log"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {

	key, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if key == nil {
		t.Error("Error generating RSA Key.")
	}
}

func TestSaveKeyPairToGog(t *testing.T) {

	key, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if key == nil {
		t.Error("Error generating RSA Key.")
		return
	}

	err = SaveKeyPairToGog(key)
	if err != nil {
		t.Error("Error saving Gob Key pair: ", err)
	}
}

func TestSaveKeyPairToPEM(t *testing.T) {

	key, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if key == nil {
		t.Error("Error generating RSA Key.")
		return
	}

	err = SaveKeyPairToPEM(key)
	if err != nil {
		t.Error("Error saving Key pair to pem file: ", err)
	}
}

/*func TestImporPrivateKeyPEMFile(t *testing.T) {
	key, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
	}

	if key == nil {
		t.Error("Error generating RSA Key.")
	}

	err = SaveKeyPairToPEM(key)
	if err != nil {
		t.Error("Error saving Key pair to pem file: ", err)
	}

	privateKeyImported, err := ImporPrivateKeyPEMFile("private.pem")
	if err != nil {
		t.Error("Error importing private key from file: ", err)
	}

	publicKey := &privateKeyImported.PublicKey

	log.Println("Private Key: ", privateKeyImported)
	log.Println("Public key: ", publicKey)
}*/

var data = "Rafael de Aquino Cunha"

func TestEncryptDecryptOAEPSHA256(t *testing.T) {
	privateKey, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if privateKey == nil {
		t.Error("Error generating RSA Key.")
		return
	}

	log.Println("Data to encrypt: ", data)

	encryptedData, err := EncryptOAEPSHA256(data, &privateKey.PublicKey)
	if err != nil {
		t.Error("Error encrypting data: ", err)
		return
	}

	//log.Println("Encrypted data: ", string(encryptedData))

	decryptedData, err := DecryptOAEPSHA256(encryptedData, privateKey)
	if err != nil {
		t.Error("Error decrypting data: ", err)
		return
	}

	log.Println("Decrypted data: ", string(decryptedData))

	if data != string(decryptedData) {
		t.Error("Error encrypting and decrypting data. Different results.")
	}

}

func TestDataSignature(t *testing.T) {
	privateKey, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if privateKey == nil {
		t.Error("Error generating RSA Key.")
		return
	}

	log.Println("Data to sign: ", data)

	signature, err := PSSSignData(data, privateKey)
	if err != nil {
		t.Error("Error signing data: ", err)
		return
	}

	//log.Println("Signature: ", string(signature))

	result, err := VerifyPSSSignature(data, signature, &privateKey.PublicKey)
	if err != nil {
		t.Error("Error veryfing signature: ", err)
		return
	}

	if !result {
		t.Error("Signature not confirmed!")
		return
	}

	log.Println("Signature confirmed!")
}
