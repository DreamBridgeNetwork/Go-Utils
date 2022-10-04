package rsakey

import (
	"log"
	"testing"
)

var data = "Rafael de Aquino Cunha"

func TestGenerateRSAKeyPair(t *testing.T) {

	log.Println("TestGenerateRSAKeyPair")

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

	log.Println("TestSaveKeyPairToGog")

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

	log.Println("TestSaveKeyPairToPEM")

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

func TestImporPrivateKeyPEMFile(t *testing.T) {

	log.Println("TestImporPrivateKeyPEMFile")

	genPrivatekey, err := GenerateRSAKeyPair(RecomendedSize)
	if err != nil {
		t.Error("Error generating RSA Key: ", err)
		return
	}

	if genPrivatekey == nil {
		t.Error("Error generating RSA Key.")
		return
	}

	err = SaveKeyPairToPEM(genPrivatekey)
	if err != nil {
		t.Error("Error saving Key pair to pem file: ", err)
		return
	}

	importedPrivateKey, err := ImporPrivateKeyPEMFile("private.pem")
	if err != nil {
		t.Error("Error importing private key from file: ", err)
		return
	}
	if importedPrivateKey == nil {
		t.Error("Error importing private key.")
		return
	}

	signatureGen, err := PSSSignData(data, genPrivatekey)
	if err != nil {
		t.Error("Error signing data with generated private key: ", err)
		return
	}
	if signatureGen == nil {
		t.Error("Error signing data with generated private key: nil signature")
		return
	}

	resultGen, err := VerifyPSSSignature(data, signatureGen, &importedPrivateKey.PublicKey)
	if err != nil {
		t.Error("Error veryfing signature: ", err)
		return
	}

	if !resultGen {
		t.Error("Signature not confirmed!")
		return
	}

	signatureImp, err := PSSSignData(data, importedPrivateKey)
	if err != nil {
		t.Error("Error signing data with imported private key: ", err)
		return
	}
	if signatureImp == nil {
		t.Error("Error signing data with imported private key: nil signature")
		return
	}

	resultImp, err := VerifyPSSSignature(data, signatureGen, &genPrivatekey.PublicKey)
	if err != nil {
		t.Error("Error veryfing signature: ", err)
		return
	}

	if !resultImp {
		t.Error("Signature not confirmed!")
		return
	}

	log.Println("Signatures confirmed!")

}

func TestEncryptDecryptOAEPSHA256(t *testing.T) {

	log.Println("TestEncryptDecryptOAEPSHA256")

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

	log.Println("TestDataSignature")

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
