package rsakey

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
	"log"
	"os"
)

const RecomendedSize = 2048

// GenerateRSAKeyPair - Generate and return a rsa key pair.
func GenerateRSAKeyPair(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		log.Println("rsakey.GenerateRSAKeyPair - Error generating rsa keys")
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		log.Println("rsakey.GenerateRSAKeyPair - Error validating generated keys")
		return nil, err
	}

	return privateKey, nil
}

// SaveKeyPairToGog - Save the private and public key to private.key and public.key files.
func SaveKeyPairToGog(key *rsa.PrivateKey) error {
	err := saveGobKey("private.key", key)
	if err != nil {
		log.Println("rsakey.SaveKeyPairToGog - Error creating private.key file.")
		return err
	}

	publicKey := key.PublicKey

	saveGobKey("public.key", publicKey)
	if err != nil {
		log.Println("rsakey.SaveKeyPairToGog - Error creating public.key file.")
		return err
	}

	return nil
}

// saveGobKey - Generate a gob file with private or public key.
func saveGobKey(fileName string, key interface{}) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Println("rsakey.SaveGobKey - Error creating the file ", fileName)
		return err
	}
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	if err != nil {
		log.Println("rsakey.SaveGobKey - Error ecoding and saving the key.")
		return err
	}

	return nil
}

// SaveKeyPairToPEM - Save the rsa key pair to private.pem and public.pem files.
func SaveKeyPairToPEM(key *rsa.PrivateKey) error {

	err := savePEMKey("private.pem", key)
	if err != nil {
		log.Println("rsakey.SaveKeyPairToGog - Error creating private.pem file.")
		return err
	}

	publicKey := key.PublicKey

	savePublicPEMKey("public.pem", publicKey)
	if err != nil {
		log.Println("rsakey.SaveKeyPairToGog - Error creating public.key file.")
		return err
	}

	return nil
}

// savePEMKey - Save a private key to a .pem file.
func savePEMKey(fileName string, privateKey *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Println("rsakey.SavePEMKey - Error creating the file ", fileName)
		return err
	}
	defer outFile.Close()

	var pemPrivateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(outFile, pemPrivateKey)
	if err != nil {
		log.Println("rsakey.SavePEMKey - Error ecoding and saving the key.")
		return err
	}

	return nil
}

// savePublicPEMKey - Save a public key to a .pem file
func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		log.Println("rsakey.SavePublicPEMKey - Error encoding public key. ")
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		log.Println("rsakey.SavePublicPEMKey - Error creating the file ", fileName)
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		log.Println("rsakey.SavePublicPEMKey - Error encoding the public key to file.")
		return err
	}

	return nil
}

// ImporPrivateKeyPEMFile - Import a private key pem file to rsa private key structure.
func ImporPrivateKeyPEMFile(fileName string) (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(fileName)
	if err != nil {
		log.Println("rsakey.ImporPrivateKeyPEMFile - Error opening privateKeyFile: ", fileName)
		return nil, err
	}
	defer privateKeyFile.Close()

	pemfileinfo, err := privateKeyFile.Stat()
	if err != nil {
		log.Println("rsakey.ImporPrivateKeyPEMFile - Error reading pem private key file information.")
		return nil, err
	}

	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		log.Println("rsakey.ImporPrivateKeyPEMFile - Error reading pem private key file data.")
		return nil, err
	}

	data, _ := pem.Decode([]byte(pembytes))
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		log.Println("rsakey.ImporPrivateKeyPEMFile - Error converting provate key data to rsa structure.")
		return nil, err
	}

	return privateKeyImported, nil
}

// EncryptOAEPSHA256 - Encrypt some data using OAEP SHA256 and rsa public key.
func EncryptOAEPSHA256(data string, publicKey *rsa.PublicKey) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(data), nil)
	if err != nil {
		log.Println("rsakey.EncryptOAEPSHA256 - Error encrypting data.")
		return nil, err
	}

	return encryptedBytes, nil
}

// DecryptOAEPSHA256 - Decrypt some data using OAEP SHA256 and rsa private key.
func DecryptOAEPSHA256(encryptedData []byte, privateKey *rsa.PrivateKey) ([]byte, error) {

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedData, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Println("rsakey.DecryptOAEPSHA256 - Error decrypting data.")
		return nil, err
	}

	return decryptedBytes, nil
}

// PSSSignData - Sign some data using PSS and a private key.
func PSSSignData(data string, privateKey *rsa.PrivateKey) ([]byte, error) {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(data))
	if err != nil {
		log.Println("rsakey.SignData - Error generating hash.")
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		log.Println("rsakey.SignData - Error signing the data.")
		return nil, err
	}

	return signature, nil
}

// VerifyPSSSignature - Verify the PSS signature of some data.
func VerifyPSSSignature(data string, signature []byte, publicKey *rsa.PublicKey) (bool, error) {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(data))
	if err != nil {
		log.Println("rsakey.VerifySignature - Error generating hash.")
		return false, err
	}
	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		log.Println("rsakey.VerifySignature - Error verifying signature.")
		return false, err
	}

	return true, nil
}

//Based on:
// https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251
// https://www.systutorials.com/how-to-generate-rsa-private-and-public-key-pair-in-go-lang/
// https://stackoverflow.com/questions/64104586/use-golang-to-get-rsa-key-the-same-way-openssl-genrsa
// https://medium.com/@Raulgzm/export-import-pem-files-in-go-67614624adc7
// https://www.sohamkamani.com/golang/rsa-encryption/
// https://gist.github.com/goliatone/e9c13e5f046e34cef6e150d06f20a34c
