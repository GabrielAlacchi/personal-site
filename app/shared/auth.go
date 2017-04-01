package shared

import (
	"github.com/gabrielalacchi/personal-site/app/model"
	"golang.org/x/crypto/bcrypt"
	"encoding/hex"
	"io"
	"crypto/rand"
)

func Authenticate(email string, password string) (bool, error) {
	var auth model.Auth

	err := auth.QueryByEmail(email)
	if err != nil {
		return false, err
	}

	digestBytes, _ := hex.DecodeString(auth.BcryptDigest)
	passwordBytes := []byte(password)
	saltBytes, _ := hex.DecodeString(auth.BcryptSalt)

	passwordBytes = append(passwordBytes, saltBytes...)

	err = bcrypt.CompareHashAndPassword(digestBytes, passwordBytes)
	return err == nil, err
}

func GenerateHash(password string) (hash string, salt string) {
	saltBytes := make([]byte, 32)
	io.ReadFull(rand.Reader, saltBytes)
	salt = hex.EncodeToString(saltBytes)
	passwordBytes := []byte(password)

	passwordBytes = append(passwordBytes, saltBytes...)

	hashBytes, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost + 3)
	hash = hex.EncodeToString(hashBytes)
	return
}
